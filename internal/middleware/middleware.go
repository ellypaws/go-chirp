package middleware

import (
	"context"
	"github.com/ellypaws/go-chirp/internal/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var JWTKey = []byte("my_secret_key")

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		claims := new(models.Claims)
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			return JWTKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if jwtExpired(claims) {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "jwt", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func jwtExpired(claims *models.Claims) bool {
	return claims.ExpiresAt < jwt.TimeFunc().Unix()
}
