package services

import (
	"errors"
	"net/http"
	"net/mail"

	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/pkg/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(db *database.Service, req models.SignupRequest) (*models.User, error) {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("username, email and password are required")
	}

	// Check if username exists
	var existingUser models.User
	result := db.Gorm().Where("username = ?", req.Username).First(&existingUser)
	if result.Error == nil {
		return nil, models.NewAuthError("username already exists", http.StatusConflict)
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, models.NewAuthError("database error", http.StatusInternalServerError)
	}

	// Check if email exists
	result = db.Gorm().Where("email = ?", req.Email).First(&existingUser)
	if result.Error == nil {
		return nil, models.NewAuthError("email already exists", http.StatusConflict)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, models.NewAuthError("failed to hash password", http.StatusInternalServerError)
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: models.Password(hashedPassword),
	}

	if err := db.Gorm().Create(user).Error; err != nil {
		return nil, models.NewAuthError("failed to create user", http.StatusInternalServerError)
	}

	return user, nil
}

func Login(db *database.Service, username, password string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, models.NewAuthError("username and password are required", http.StatusBadRequest)
	}

	var user models.User
	var result *gorm.DB

	// Try to parse as email first
	_, err := mail.ParseAddress(username)
	if err == nil {
		result = db.Gorm().Where("email = ?", username).First(&user)
	} else {
		result = db.Gorm().Where("username = ?", username).First(&user)
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, models.NewAuthError("user not found", http.StatusNotFound)
		}
		return nil, models.NewAuthError("database error", http.StatusInternalServerError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, models.NewAuthError("invalid credentials", http.StatusUnauthorized)
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
