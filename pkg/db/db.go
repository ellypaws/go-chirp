package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"regexp"
	"strings"
)

var DB *sql.DB

func InitDB() {
	err := assertDatabase()
	if err != nil {
		panic(err)
	}
	dataSource, err := loadDataSource()
	if err != nil {
		panic(err)
	}
	DB, err = sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	err = migrations(DB)
	if err != nil {
		panic(err)
	}
}

func loadDataSource() (string, error) {
	var dataSourceArgs []string
	envVars := map[string]bool{
		"host":     true,
		"port":     true,
		"dbname":   false,
		"user":     true,
		"password": true,
		"sslmode":  false,
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

func migrations(database *sql.DB) error {
	_, err := database.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS tweets (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS follows (
			id SERIAL PRIMARY KEY,
			follower_id INTEGER NOT NULL REFERENCES users(id),
			following_id INTEGER NOT NULL REFERENCES users(id)
		);
	`)
	return err
}

func assertDatabase() error {
	dataSource, err := loadDataSource()
	if err != nil {
		return fmt.Errorf("failed to load data source: %w", err)
	}

	dataSource = regexp.MustCompile(`dbname=\w+ ?`).ReplaceAllString(dataSource, "")
	adminDB, err := sql.Open("postgres", dataSource)
	if err != nil {
		return fmt.Errorf("failed to connect to admin database: %w", err)
	}
	defer adminDB.Close()

	var exists bool
	err = adminDB.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname='chirp')").Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	if !exists {
		_, err = adminDB.Exec("CREATE DATABASE chirp")
		if err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
		log.Printf("Database 'chirp' created")
	} else {
		log.Printf("Loading existing database 'chirp'")
	}

	return nil
}
