package models

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
