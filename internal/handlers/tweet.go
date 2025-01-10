package handlers

import (
	"encoding/json"
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/internal/services"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateTweetHandler(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	claims, ok := r.Context().Value("jwt").(*models.Claims)
	if !ok {
		http.Error(w, "Failed to get user from token", http.StatusUnauthorized)
		return
	}

	tweet.UserID = claims.UserID

	err = services.CreateTweet(tweet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteTweetHandler(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	claims, ok := r.Context().Value("jwt").(*models.Claims)
	if !ok {
		http.Error(w, "Failed to get user from token", http.StatusUnauthorized)
		return
	}

	err = services.DeleteTweet(tweet.ID, claims.UserID)
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tweets)
}

func FetchUserTweetsHandler(w http.ResponseWriter, r *http.Request) {
	var tweets []models.Tweet
	var err error
	vars := mux.Vars(r)
	if username := vars["username"]; username != "" {
		tweets, err = services.FetchUserTweetsByUsername(username)
	} else if userID := vars["userID"]; userID != "" {
		tweets, err = services.FetchUserTweets(userID)
	} else {
		http.Error(w, "missing username or userID query parameter", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tweets)
}
