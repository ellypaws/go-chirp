package main

import (
	"log"
	"net/http"

	"github.com/ellypaws/go-chirp/internal/handlers"
	"github.com/ellypaws/go-chirp/internal/middleware"
	"github.com/ellypaws/go-chirp/pkg/db"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/tweet", middleware.JWTMiddleware(handlers.CreateTweetHandler{}).ServeHTTP).Methods("POST")
	router.HandleFunc("/follow", middleware.JWTMiddleware(handlers.FollowHandler{}).ServeHTTP).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
