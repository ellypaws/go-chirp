package services

import (
	"errors"
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/pkg/db"

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
