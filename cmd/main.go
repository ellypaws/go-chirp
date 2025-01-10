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
	v1 := router.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	v1.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	v1.HandleFunc("/tweets", handlers.FetchTweetsHandler).Methods("GET")
	v1.HandleFunc("/tweet", middleware.JWTMiddleware(http.HandlerFunc(handlers.CreateTweetHandler)).ServeHTTP).Methods("POST")
	v1.HandleFunc("/follow", middleware.JWTMiddleware(http.HandlerFunc(handlers.FollowHandler)).ServeHTTP).Methods("POST")
	v1.HandleFunc("/user/{userID:[0-9]+}/tweets", handlers.FetchUserTweetsHandler).Methods("GET")
	v1.HandleFunc("/username/{username}/tweets", handlers.FetchUserTweetsHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
