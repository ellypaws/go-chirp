package services

import (
    "twitter-backend/internal/models"
    "twitter-backend/pkg/db"
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
