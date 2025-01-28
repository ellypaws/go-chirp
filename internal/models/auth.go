package models

import "github.com/golang-jwt/jwt"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthError struct {
	message string
	code    int
}

func NewAuthError(message string, code int) AuthError {
	return AuthError{message: message, code: code}
}

func (e AuthError) Error() string {
	return e.message
}

func (e AuthError) StatusCode() int {
	return e.code
}
