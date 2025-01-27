package services

import (
	"errors"

	"github.com/ellypaws/go-chirp/internal/models"
	"gorm.io/gorm"
)

func CreateTweet(db *gorm.DB, tweet models.Tweet) (models.Tweet, error) {
	result := db.Create(&tweet)
	if result.Error != nil {
		return tweet, result.Error
	}

	result = db.Where("id = ?", tweet.UserID).First(&tweet.User)
	if result.Error != nil {
		return tweet, result.Error
	}
	tweet.User.TweetCount++
	result = db.Save(&tweet.User)
	if result.Error != nil {
		return tweet, result.Error
	}

	if tweet.ParentID != nil {
		result = db.Model(&models.Tweet{}).Where("id = ?", *tweet.ParentID).Update("replies_count", gorm.Expr("replies_count + ?", 1))
		if result.Error != nil {
			return tweet, result.Error
		}
	}

	return tweet, result.Error
}

func DeleteTweet(db *gorm.DB, tweetID, userID uint) error {
	var tweet models.Tweet
	result := db.Where("id = ?", tweetID).First(&tweet)
	if result.Error != nil {
		return result.Error
	}

	if tweet.UserID != userID {
		return errors.New("not the owner of the tweet")
	}

	result = db.Delete(&tweet)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	result = db.Model(&models.User{}).Where("id = ?", userID).Update("tweet_count", gorm.Expr("tweet_count - 1"))
	if result.Error != nil {
		return result.Error
	}

	if tweet.ParentID != nil {
		db.Model(&models.Tweet{}).Where("id = ?", *tweet.ParentID).Update("replies_count", gorm.Expr("replies_count - 1"))
	}

	var tweets []models.Like
	db.Where("tweet_id = ?", tweetID).Find(&tweets)
	for _, tweet := range tweets {
		db.Delete(tweet)
	}

	return result.Error
}

func FetchTweets(db *gorm.DB) ([]models.Tweet, error) {
	var tweets []models.Tweet
	result := db.Preload("User").Preload("Likes").Preload("Replies").Find(&tweets)
	return tweets, result.Error
}
