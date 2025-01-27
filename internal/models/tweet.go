package models

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	User     *User  `json:"user,omitempty"`
	Content  string `json:"content"`
	ParentID *uint  `json:"parent_id,omitempty"`

	Likes     []*Like `json:"liked_by,omitempty"`
	LikeCount uint    `json:"like_count"`

	Replies      []*Tweet `json:"replies,omitempty" gorm:"foreignKey:parent_id"`
	RepliesCount uint     `json:"replies_count"`
}

type Like struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	TweetID uint `json:"tweet_id"`

	Tweet *Tweet `json:"tweet"`
	User  *User  `json:"user"`
}
