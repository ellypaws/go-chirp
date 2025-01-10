package main

import (
	"log"
	"net/http"

	"github.com/ellypaws/go-chirp/internal/handlers"
	"github.com/ellypaws/go-chirp/pkg/db"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/tweet", handlers.CreateTweetHandler).Methods("POST")
	router.HandleFunc("/follow", handlers.FollowHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
