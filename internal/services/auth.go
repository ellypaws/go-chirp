package services

import (
	"errors"
	"net/mail"

	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/pkg/db"

	"golang.org/x/crypto/bcrypt"
)

func Signup(db *database.Service, user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return db.CreateUser(user)
}

func Login(db *database.Service, username, password string) (*models.User, error) {
	var user models.User
	_, err := mail.ParseAddress(username)
	if err == nil {
		user, err = db.GetUserByEmail(username)
	} else {
		user, err = db.GetUserByUsername(username)
	}
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
