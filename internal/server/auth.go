package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ellypaws/go-chirp/internal/middleware"
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/internal/services"
	"github.com/ellypaws/go-chirp/internal/utils"

	"github.com/golang-jwt/jwt"
)

func (s *Server) SignupHandler(w http.ResponseWriter, r *http.Request) {
	user, err := utils.Decode[models.User](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = services.Signup(s.db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	login, err := utils.Decode[models.Credentials](r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}
	if login.Username == "" || login.Password == "" {
		http.Error(w, "username and password are required", http.StatusBadRequest)
		return
	}
	user, err := services.Login(s.db, login.Username, login.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("error logging in: %v", err), http.StatusBadRequest)
		return
	}

	token, err := generateJWT(user)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	user.Password = ""
	_ = utils.Encode(w, models.LoginResponse{
		User:  user,
		Token: token,
	})
}

func generateJWT(user *models.User) (string, error) {
	claims := models.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(middleware.JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
