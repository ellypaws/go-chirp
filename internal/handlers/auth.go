package handlers

import (
    "encoding/json"
    "net/http"
    "twitter-backend/internal/models"
    "twitter-backend/internal/services"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    err := services.Signup(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Handle login and return JWT token
}
