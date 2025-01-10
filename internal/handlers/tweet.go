package handlers

import (
    "encoding/json"
    "net/http"
    "twitter-backend/internal/models"
    "twitter-backend/internal/services"
)

func CreateTweetHandler(w http.ResponseWriter, r *http.Request) {
    var tweet models.Tweet
    json.NewDecoder(r.Body).Decode(&tweet)
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
