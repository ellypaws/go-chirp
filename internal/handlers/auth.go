package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ellypaws/go-chirp/internal/middleware"
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/internal/services"

	"github.com/golang-jwt/jwt"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = services.Signup(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login models.Credentials
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}
	user, err := services.Login(login.Username, login.Password)
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
	_ = json.NewEncoder(w).Encode(models.LoginResponse{
		User:  user,
		Token: token,
	})
}

func generateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(middleware.JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
