package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ellypaws/go-chirp/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

var db *sql.DB

func InitDB() {
	dataSource, err := loadDataSource()
	if err != nil {
		panic(err)
	}
	db, err = sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

}

func loadDataSource() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	var dataSourceArgs []string
	envVars := map[string]bool{
		"host":     true,
		"port":     true,
		"user":     true,
		"password": true,
	}

	for key, required := range envVars {
		envVar, err := getEnv(key, required)
		if err != nil {
			return "", err
		}
		if envVar != "" {
			dataSourceArgs = append(dataSourceArgs, envVar)
		}
	}

	return strings.Join(dataSourceArgs, " "), nil
}

func getEnv(key string, required ...bool) (string, error) {
	envKey := fmt.Sprintf("DB_%s", strings.ToUpper(key))
	value := os.Getenv(envKey)
	if len(required) > 0 && required[0] {
		if value == "" {
			return "", errors.New(envKey + " is required")
		}
	} else if value == "" {
		return "", nil
	}
	return fmt.Sprintf("%s=%s", strings.ToLower(key), value), nil
}

// TODO: refactor this to set pragma version
func migrations(database *sql.DB) error {
	err := assertDatabase(database)
	if err != nil {
		return err
	}
	_, err = database.Exec(`
		\c chirp;
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS tweets (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			body TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS follows (
			id SERIAL PRIMARY KEY,
			follower_id INTEGER NOT NULL,
			following_id INTEGER NOT NULL
		);
	`)
	return err
}

func assertDatabase(database *sql.DB) error {
	_, err := database.Exec("SELECT 1 FROM pg_database WHERE datname='chirp'")
	if err != nil {
		_, err = db.Exec("CREATE DATABASE chirp")
	}
	return err
}

func CreateUser(user models.User) error {
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	return err
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, username, email, password FROM users WHERE username = $1", username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	)
	return user, err
}
