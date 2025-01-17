package server

import (
	"fmt"
	database "github.com/ellypaws/go-chirp/pkg/db"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port int
	db   *database.Service
}

func NewServer() *http.Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	NewServer := &Server{
		port: port,
		db:   database.InitDB(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
