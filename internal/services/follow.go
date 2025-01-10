package services

import (
    "errors"
    "twitter-backend/internal/models"
    "twitter-backend/pkg/db"
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
