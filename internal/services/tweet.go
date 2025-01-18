package services

import (
	"fmt"
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/pkg/db"
	"time"
)

func CreateTweet(db *database.Service, tweet models.Tweet) (models.Tweet, error) {
	if err := db.CreateTweet(tweet); err != nil {
		return models.Tweet{}, err
	}
	tweet.CreatedAt = time.Now().UTC().Format("2006-01-02 15:04:05.999999")
	return tweet, nil
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
