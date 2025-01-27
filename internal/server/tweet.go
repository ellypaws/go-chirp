package server

import (
	"gorm.io/gorm"
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

	tweet, err = services.CreateTweet(s.db.Gorm(), tweet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
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

	err = services.DeleteTweet(s.db.Gorm(), tweet.ID, claims.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) FetchTweetsHandler(w http.ResponseWriter, r *http.Request) {
	tweets, err := services.FetchTweets(s.db.Gorm())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = utils.Encode(w, tweets)
}

func (s *Server) FetchUserTweetsHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var result *gorm.DB
	if username := r.PathValue("username"); username != "" {
		result = s.db.Gorm().Where("username = ?", username).Preload("Tweets").First(&user)
	} else if userID := r.PathValue("userID"); userID != "" {
		result = s.db.Gorm().Where("id = ?", userID).Preload("Tweets").First(&user)
	} else {
		http.Error(w, "missing username or userID query parameter", http.StatusBadRequest)
		return
	}
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	shallow := user
	shallow.Tweets = nil
	for i := range user.Tweets {
		user.Tweets[i].User = &shallow
	}
	_ = utils.Encode(w, user.Tweets)
}
