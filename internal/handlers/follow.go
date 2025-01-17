package handlers

import (
	"net/http"

	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/ellypaws/go-chirp/internal/services"
	"github.com/ellypaws/go-chirp/internal/utils"
)

func FollowHandler(w http.ResponseWriter, r *http.Request) {
	follow, err := utils.Decode[models.Follow](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value("jwt").(*models.Claims)
	if !ok {
		http.Error(w, "Failed to get user from token", http.StatusUnauthorized)
		return
	}

	follow.FollowerID = claims.UserID

	err = services.FollowUser(follow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func UnfollowHandler(w http.ResponseWriter, r *http.Request) {
	follow, err := utils.Decode[models.Follow](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = services.UnfollowUser(follow)
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
	_ = utils.Encode(w, followers)
}

func GetFollowingHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	following, err := services.GetFollowing(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = utils.Encode(w, following)
}
