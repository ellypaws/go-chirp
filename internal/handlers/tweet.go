package handlers

import (
	"encoding/json"
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/internal/services"
	"net/http"
)

func CreateTweetHandler(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	json.NewDecoder(r.Body).Decode(&tweet)

	claims, ok := r.Context().Value("jwt").(*models.Claims)
	if !ok {
		http.Error(w, "Failed to get user from token", http.StatusUnauthorized)
		return
	}

	tweet.UserID = claims.UserID

	err := services.CreateTweet(tweet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteTweetHandler(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	json.NewDecoder(r.Body).Decode(&tweet)
	err := services.DeleteTweet(tweet.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func FetchTweetsHandler(w http.ResponseWriter, r *http.Request) {
	tweets, err := services.FetchTweets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tweets)
}
