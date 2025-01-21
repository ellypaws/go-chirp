package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Picture  string `json:"picture,omitempty"`

	Email    string `json:"email"`
	Password string `json:"password,omitempty"`

	Following      []*User `json:"following,omitempty" gorm:"many2many:follows;"`
	FollowingCount uint    `json:"following_count"`

	Followers     []*User `json:"followers,omitempty" gorm:"many2many:follows;"`
	FollowerCount uint    `json:"follower_count"`

	Tweets     []*Tweet `json:"tweets,omitempty"`
	TweetCount uint     `json:"tweet_count"`

	Likes      []*Like `json:"liked_tweets,omitempty"`
	LikesCount uint    `json:"likes_count"`
}
