package server

import (
	"github.com/ellypaws/go-chirp/internal/middleware"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("POST /signup", s.SignupHandler)
	router.HandleFunc("POST /login", s.LoginHandler)
	router.HandleFunc("GET /tweets", s.FetchTweetsHandler)
	router.HandleFunc("GET /user/{userID}/tweets", s.FetchUserTweetsHandler)
	router.HandleFunc("GET /username/{username}/tweets", s.FetchUserTweetsHandler)
	router.Handle("POST /tweet", middleware.JWTMiddleware(http.HandlerFunc(s.CreateTweetHandler)))
	router.Handle("POST /follow", middleware.JWTMiddleware(http.HandlerFunc(s.FollowHandler)))

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	api := http.NewServeMux()
	api.Handle("/api/", http.StripPrefix("/api", v1))

	// Wrap the mux with CORS middleware
	return s.corsMiddleware(api)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}
