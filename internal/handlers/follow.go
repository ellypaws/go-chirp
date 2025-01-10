package handlers

import (
    "encoding/json"
    "net/http"
    "twitter-backend/internal/models"
    "twitter-backend/internal/services"
)

func FollowHandler(w http.ResponseWriter, r *http.Request) {
    var follow models.Follow
    json.NewDecoder(r.Body).Decode(&follow)
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
