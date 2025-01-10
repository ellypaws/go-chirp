package db

import (
	"github.com/ellypaws/go-chirp/internal/models"
)

func CreateUser(user models.User) error {
	_, err := db.Exec(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		user.Username, user.Email, user.Password,
	)
	return err
}

func GetUserByID(userID string) (models.User, error) {
	var user models.User
	err := db.QueryRow(
		"SELECT id, username, email, password FROM users WHERE id = $1", userID,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := db.QueryRow(
		"SELECT id, username, email, password FROM users WHERE username = $1", username,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := db.QueryRow(
		"SELECT id, username, email, password FROM users WHERE email = $1", email,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

func CreateTweet(tweet models.Tweet) error {
	_, err := db.Exec(
		"INSERT INTO tweets (user_id, content) VALUES ($1, $2)",
		tweet.UserID,
		tweet.Content,
	)
	return err
}

func DeleteTweet(tweetID int) error {
	_, err := db.Exec("DELETE FROM tweets WHERE id = $1", tweetID)
	return err
}

func FetchTweet(tweetID int) (models.Tweet, error) {
	var tweet models.Tweet
	err := db.QueryRow(
		"SELECT id, user_id, content, created_at FROM tweets WHERE id = $1", tweetID,
	).Scan(&tweet.ID, &tweet.UserID, &tweet.Content, &tweet.CreatedAt)
	return tweet, err
}

func FetchUserTweets(userID string) ([]models.Tweet, error) {
	rows, err := db.Query("SELECT id, user_id, content, created_at FROM tweets WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets []models.Tweet
	for rows.Next() {
		var tweet models.Tweet
		err := rows.Scan(&tweet.ID, &tweet.UserID, &tweet.Content, &tweet.CreatedAt)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}
	return tweets, nil
}

func FetchUserTweetsByUsername(username string) ([]models.Tweet, error) {
	rows, err := db.Query(`SELECT tweets.id, tweets.user_id, tweets.content, tweets.created_at
								 FROM tweets
								 JOIN users
								 ON tweets.user_id = users.id
								 WHERE users.username = $1`, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets []models.Tweet
	for rows.Next() {
		var tweet models.Tweet
		err := rows.Scan(&tweet.ID, &tweet.UserID, &tweet.Content, &tweet.CreatedAt)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}
	return tweets, nil
}

func FetchTweets() ([]models.Tweet, error) {
	rows, err := db.Query("SELECT id, user_id, content, created_at FROM tweets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets []models.Tweet
	for rows.Next() {
		var tweet models.Tweet
		err := rows.Scan(&tweet.ID, &tweet.UserID, &tweet.Content, &tweet.CreatedAt)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}
	return tweets, nil
}

func CreateFollow(follow models.Follow) error {
	_, err := db.Exec("INSERT INTO follows (follower_id, following_id) VALUES ($1, $2)",
		follow.FollowerID, follow.FollowedID,
	)
	return err
}

func DeleteFollow(follow models.Follow) error {
	_, err := db.Exec("DELETE FROM follows WHERE follower_id = $1 AND following_id = $2",
		follow.FollowerID, follow.FollowedID,
	)
	return err
}

func GetFollowers(userID string) ([]models.User, error) {
	rows, err := db.Query(`SELECT users.id, users.username, users.email
                                 FROM users
                                 JOIN follows
                                 ON users.id = follows.follower_id
                                 WHERE follows.following_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		followers = append(followers, user)
	}
	return followers, nil
}

func GetFollowing(userID string) ([]models.User, error) {
	rows, err := db.Query(`SELECT users.id, users.username, users.email
                                 FROM users
                                 JOIN follows 
                                 ON users.id = follows.following_id
                                 WHERE follows.follower_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var following []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		following = append(following, user)
	}
	return following, nil
}
