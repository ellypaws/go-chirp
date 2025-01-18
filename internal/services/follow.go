package services

import (
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/pkg/db"
)

func FollowUser(db *database.Service, follow models.Follow) error {
	return db.CreateFollow(follow)
}

func UnfollowUser(db *database.Service, follow models.Follow) error {
	return db.DeleteFollow(follow)
}

func GetFollowers(db *database.Service, userID string) ([]models.User, error) {
	return db.GetFollowers(userID)
}

func GetFollowing(db *database.Service, userID string) ([]models.User, error) {
	return db.GetFollowing(userID)
}
