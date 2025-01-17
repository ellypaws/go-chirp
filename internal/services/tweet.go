package services

import (
	"fmt"
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/pkg/db"
)

func CreateTweet(db *database.Service, tweet models.Tweet) error {
	return db.CreateTweet(tweet)
}

func DeleteTweet(db *database.Service, tweetID, userID int) error {
	tweet, err := db.FetchTweet(tweetID)
	if err != nil {
		return err
	}
	if tweet.UserID != userID {
		return fmt.Errorf("user %d is not the owner of tweet %d", userID, tweetID)
	}

	return db.DeleteTweet(tweetID)
}

func FetchTweets(db *database.Service) ([]models.Tweet, error) {
	return db.FetchTweets()
}

func FetchUserTweets(db *database.Service, userID string) ([]models.Tweet, error) {
	return db.FetchUserTweets(userID)
}

func FetchUserTweetsByUsername(db *database.Service, username string) ([]models.Tweet, error) {
	return db.FetchUserTweetsByUsername(username)
}
