package services

import (
	"errors"
	"gorm.io/gorm"
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
	return db.Gorm().Create(&user).Error
}

func Login(db *database.Service, username, password string) (*models.User, error) {
	var user models.User
	var result *gorm.DB
	_, err := mail.ParseAddress(username)
	if err == nil {
		result = db.Gorm().Where("email = ?", username).First(&user)
	} else {
		result = db.Gorm().Where("username = ?", username).First(&user)
	}
	if result.Error != nil {
		return nil, result.Error
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func GetUserByID(db *database.Service, userID uint) (*models.User, error) {
	var user models.User
	result := db.Gorm().Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
