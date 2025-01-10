package services

import (
	"errors"
	"time"

	"github.com/ellypaws/go-chirp/internal/middleware"
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/pkg/db"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return db.CreateUser(user)
}

func Login(username, password string) (models.User, error) {
	user, err := db.GetUserByUsername(username)
	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("invalid credentials")
	}

	return user, nil
}

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(middleware.JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
