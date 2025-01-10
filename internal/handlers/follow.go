package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/internal/services"
)

func FollowHandler(w http.ResponseWriter, r *http.Request) {
	var follow models.Follow
	json.NewDecoder(r.Body).Decode(&follow)

	claims, ok := r.Context().Value("jwt").(*models.Claims)
	if !ok {
		http.Error(w, "Failed to get user from token", http.StatusUnauthorized)
		return
	}

	follow.FollowerID = claims.UserID

	err := services.FollowUser(follow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func UnfollowHandler(w http.ResponseWriter, r *http.Request) {
	var follow models.Follow
	json.NewDecoder(r.Body).Decode(&follow)
	err := services.UnfollowUser(follow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetFollowersHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	followers, err := services.GetFollowers(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(followers)
}

func GetFollowingHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	following, err := services.GetFollowing(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(following)
}
