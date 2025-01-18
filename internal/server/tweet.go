package server

import (
	"fmt"
	"net/http"

	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/internal/services"
	"github.com/ellypaws/go-chirp/internal/utils"
)

func (s *Server) CreateTweetHandler(w http.ResponseWriter, r *http.Request) {
	tweet, err := utils.Decode[models.Tweet](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value("jwt").(*models.Claims)
	if !ok {
		http.Error(w, "Failed to get user from token", http.StatusUnauthorized)
		return
	}

	tweet.UserID = claims.UserID

	tweet, err = services.CreateTweet(s.db, tweet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = utils.Encode(w, tweet)
}

func (s *Server) DeleteTweetHandler(w http.ResponseWriter, r *http.Request) {
	tweet, err := utils.Decode[models.Tweet](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value("jwt").(*models.Claims)
	if !ok {
		http.Error(w, "Failed to get user from token", http.StatusUnauthorized)
		return
	}

	err = services.DeleteTweet(s.db, tweet.ID, claims.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) FetchTweetsHandler(w http.ResponseWriter, r *http.Request) {
	tweets, err := services.FetchTweets(s.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = utils.Encode(w, tweets)
}

func (s *Server) FetchUserTweetsHandler(w http.ResponseWriter, r *http.Request) {
	var tweets []models.Tweet
	var err error
	if username := r.PathValue("username"); username != "" {
		_, err = s.db.GetUserByUsername(username)
		if err != nil {
			http.Error(w, fmt.Sprintf("error fetching user by username: %v", err), http.StatusBadRequest)
			return
		}
		tweets, err = services.FetchUserTweetsByUsername(s.db, username)
	} else if userID := r.PathValue("userID"); userID != "" {
		_, err = s.db.GetUserByID(userID)
		if err != nil {
			http.Error(w, fmt.Sprintf("error fetching user by userID: %v", err), http.StatusBadRequest)
			return
		}
		tweets, err = services.FetchUserTweets(s.db, userID)
	} else {
		http.Error(w, "missing username or userID query parameter", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = utils.Encode(w, tweets)
}
