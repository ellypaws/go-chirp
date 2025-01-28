package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/ellypaws/go-chirp/internal/middleware"
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/internal/services"
	"github.com/ellypaws/go-chirp/internal/utils"

	"github.com/golang-jwt/jwt"
)

func (s *Server) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req models.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := services.Signup(s.db, models.SignupRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		var authError models.AuthError
		if errors.As(err, &authError) {
			http.Error(w, authError.Error(), authError.StatusCode())
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	token, err := generateJWT(user)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = utils.Encode(w, models.LoginResponse{
		User:  user,
		Token: token,
	})
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	login, err := utils.Decode[models.Credentials](r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := services.Login(s.db, login.Username, login.Password)
	if err != nil {
		var authError models.AuthError
		if errors.As(err, &authError) {
			http.Error(w, authError.Error(), authError.StatusCode())
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
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

func (s *Server) VerifyHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("jwt").(*models.Claims)
	if !ok {
		http.Error(w, "Failed to get user from token", http.StatusUnauthorized)
		return
	}

	user, err := services.GetUserByID(s.db, claims.UserID)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	_ = utils.Encode(w, models.LoginResponse{
		User:  user,
		Token: strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "),
	})
}
