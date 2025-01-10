package middleware

import (
	"context"
	"crypto/rand"
	"github.com/ellypaws/go-chirp/internal/models"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

var JWTKey = func() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret != "" {
		return []byte(secret)
	}
	randomKey := make([]byte, 32)
	_, err := rand.Read(randomKey)
	if err != nil {
		panic("Failed to generate random key")
	}
	return randomKey
}()

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
