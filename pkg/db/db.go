package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/ellypaws/go-chirp/internal/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	db *gorm.DB
}

func (s *Service) Gorm() *gorm.DB {
	return s.db
}

func InitDB() *Service {
	dataSource, err := loadDataSource()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(dataSource))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Tweet{}, &models.Follow{}, &models.Like{})
	if err != nil {
		panic(err)
	}
	return &Service{db: db}
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

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *Service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	sql, err := s.db.DB()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Ping the database
	err = sql.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := sql.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}
