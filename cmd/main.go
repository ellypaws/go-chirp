package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/ellypaws/go-chirp/internal/handlers"
	"github.com/ellypaws/go-chirp/internal/middleware"
	"github.com/ellypaws/go-chirp/pkg/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

	router := http.NewServeMux()
	router.HandleFunc("POST /signup", handlers.SignupHandler)
	router.HandleFunc("POST /login", handlers.LoginHandler)
	router.HandleFunc("GET /tweets", handlers.FetchTweetsHandler)
	router.HandleFunc("POST /tweet", middleware.JWTMiddleware(http.HandlerFunc(handlers.CreateTweetHandler)).ServeHTTP)
	router.HandleFunc("POST /follow", middleware.JWTMiddleware(http.HandlerFunc(handlers.FollowHandler)).ServeHTTP)
	router.HandleFunc("GET /user/{userID}/tweets", handlers.FetchUserTweetsHandler)
	router.HandleFunc("GET /username/{username}/tweets", handlers.FetchUserTweetsHandler)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	api := http.NewServeMux()
	api.Handle("/api/", http.StripPrefix("/api", v1))

	log.Fatal(http.ListenAndServe(":8080", api))
}
