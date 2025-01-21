package services

import (
	"github.com/ellypaws/go-chirp/internal/models"
	"gorm.io/gorm"
)

func FollowUser(db *gorm.DB, follow models.Follow) error {
	result := db.Create(&follow)
	return result.Error
}

func UnfollowUser(db *gorm.DB, follow models.Follow) error {
	result := db.Delete(&follow)
	return result.Error
}

func GetFollowers(db *gorm.DB, userID string) ([]models.User, error) {
	var followers []models.User
	err := db.Where("id = ?", userID).Find(&models.User{}).Association("Followers").Find(&followers)
	return followers, err
}

func GetFollowing(db *gorm.DB, userID string) ([]models.User, error) {
	var following []models.User
	err := db.Where("id = ?", userID).Find(&models.User{}).Association("Following").Find(&following)
	return following, err
}
