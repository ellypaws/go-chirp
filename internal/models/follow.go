package models

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowerID uint `json:"follower_id"`
	FollowedID uint `json:"followed_id"`
}
