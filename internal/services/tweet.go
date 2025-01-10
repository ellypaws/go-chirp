package services

import (
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/pkg/db"
)

func CreateTweet(tweet models.Tweet) error {
	return db.CreateTweet(tweet)
}

func DeleteTweet(tweetID int) error {
	return db.DeleteTweet(tweetID)
}

func FetchTweets() ([]models.Tweet, error) {
	return db.FetchTweets()
}
