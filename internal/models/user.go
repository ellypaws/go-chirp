package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Picture  string `json:"picture,omitempty"`

	Email    string   `json:"email"`
	Password Password `json:"password,omitempty"`

	Following      []*User `json:"following,omitempty" gorm:"many2many:follows;"`
	FollowingCount uint    `json:"following_count"`

	Followers     []*User `json:"followers,omitempty" gorm:"many2many:follows;"`
	FollowerCount uint    `json:"follower_count"`

	Tweets     []*Tweet `json:"tweets,omitempty"`
	TweetCount uint     `json:"tweet_count"`

	Likes      []*Like `json:"liked_tweets,omitempty"`
	LikesCount uint    `json:"likes_count"`
}

type Password string

func (Password) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

func (p *Password) Scan(value interface{}) error {
	if value == nil {
		*p = ""
		return nil
	}
	*p = Password(value.(string))
	return nil
}

func (p Password) Value() (driver.Value, error) {
	return string(p), nil
}
