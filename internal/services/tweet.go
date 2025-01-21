package services

import (
	"github.com/ellypaws/go-chirp/internal/models"
	"gorm.io/gorm"
)

func CreateTweet(db *gorm.DB, tweet models.Tweet) (models.Tweet, error) {
	result := db.Create(&tweet)
	if result.Error != nil {
		return tweet, result.Error
	}

	var user models.User
	result = db.Where("id = ?", tweet.UserID).First(&user)
	if result.Error != nil {
		return tweet, result.Error
	}
	user.TweetCount++
	result = db.Save(&user)
	if result.Error != nil {
		return tweet, result.Error
	}

	return tweet, result.Error
}

func DeleteTweet(db *gorm.DB, tweetID, userID uint) error {
	var tweet models.Tweet
	result := db.Where("id = ?", tweetID).Preload("Likes").Preload("Replies").First(&tweet)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&tweet)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	var user models.User
	result = db.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return result.Error
	}
	user.TweetCount--
	result = db.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	var parentTweet models.Tweet
	if tweet.ParentID != nil {
		result = db.Where("id = ?", *tweet.ParentID).First(&parentTweet)
		if result.Error != nil {
			return result.Error
		}
		parentTweet.RepliesCount--
		result = db.Save(&parentTweet)
		if result.Error != nil {
			return result.Error
		}
	}

	return result.Error
}

func FetchTweets(db *gorm.DB) ([]models.Tweet, error) {
	var tweets []models.Tweet
	result := db.Preload("Likes").Preload("Replies").Find(&tweets)
	return tweets, result.Error
}

func FetchUserTweets(db *gorm.DB, userID string) ([]models.Tweet, error) {
	var tweets []models.Tweet
	result := db.Where("user_id = ?", userID).Preload("Likes").Preload("Replies").Find(&tweets)
	return tweets, result.Error
}

func FetchUserTweetsByUsername(db *gorm.DB, username string) ([]models.Tweet, error) {
	var tweets []models.Tweet
	result := db.Where("username = ?", username).Preload("Likes").Preload("Replies").Find(&tweets)
	return tweets, result.Error
}
