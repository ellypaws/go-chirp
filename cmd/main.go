package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/ellypaws/go-chirp/internal/handlers"
	"github.com/ellypaws/go-chirp/internal/middleware"
	"github.com/ellypaws/go-chirp/pkg/db"

	"github.com/gorilla/mux"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/tweet", middleware.JWTMiddleware(http.HandlerFunc(handlers.CreateTweetHandler)).ServeHTTP).Methods("POST")
	router.HandleFunc("/follow", middleware.JWTMiddleware(http.HandlerFunc(handlers.FollowHandler)).ServeHTTP).Methods("POST")
	router.HandleFunc("/user/{userID}/tweets", middleware.JWTMiddleware(http.HandlerFunc(handlers.FetchUserTweetsHandler)).ServeHTTP).Methods("GET")
	router.HandleFunc("/username/{username}/tweets", handlers.FetchUserTweetsHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
