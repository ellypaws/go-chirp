package services

import (
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/pkg/db"
)

func FollowUser(follow models.Follow) error {
	return db.CreateFollow(follow)
}

func UnfollowUser(follow models.Follow) error {
	return db.DeleteFollow(follow)
}

func GetFollowers(userID string) ([]models.User, error) {
	return db.GetFollowers(userID)
}

func GetFollowing(userID string) ([]models.User, error) {
	return db.GetFollowing(userID)
}
