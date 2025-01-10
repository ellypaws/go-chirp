package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "twitter-backend/internal/models"
)

var db *sql.DB

func InitDB() {
    var err error
    db, err = sql.Open("postgres", "user=youruser dbname=twitter_clone sslmode=disable")
    if err != nil {
        panic(err)
    }
}

func CreateUser(user models.User) error {
    _, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
    return err
}
