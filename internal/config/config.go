package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"net/url"
	"strings"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"feature-flag-service/internal/models"
)

// Global variables for DB and Redis
var (
	DB  *gorm.DB
	RDB *redis.Client
	Ctx = context.Background()
)
// RunMigrations applies database migration
func RunMigrations() {
	err := DB.AutoMigrate(&models.FeatureFlag{}, &models.User{})
	if err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}
	fmt.Println("✅ Database migrations applied successfully")
}

// ConnectDB initializes PostgreSQL connection
func ConnectDB() {
	var dsn string

	// Check if running tests
	if os.Getenv("TEST_MODE") == "true" {
		fmt.Println("🛠️ Running in TEST mode: Connecting to test database")
		dsn = "postgres://postgres:password@test-postgres:5432/test_feature_flags?sslmode=disable"
	} else {
		dsn = os.Getenv("DATABASE_URL")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Could not connect to PostgreSQL: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
}


// ConnectRedis initializes Redis connection
func ConnectRedis() {
	redisURL := os.Getenv("REDIS_URL")

	// Parse REDIS_URL if it contains authentication
	if strings.HasPrefix(redisURL, "redis://") {
		parsedURL, err := url.Parse(redisURL)
		if err != nil {
			log.Fatalf("❌ Invalid REDIS_URL format: %v", err)
		}

		// Extract Redis credentials
		redisAddr := parsedURL.Host
		redisPassword, _ := parsedURL.User.Password()

		RDB = redis.NewClient(&redis.Options{
			Addr:     redisAddr,     // Host:Port
			Password: redisPassword, // Extracted password
			DB:       0,
		})
	} else {
		// Fallback if REDIS_URL is incorrectly formatted
		RDB = redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: os.Getenv("REDIS_PASSWORD"), // Fallback password
			DB:       0,
		})
	}

	// Test Redis connection
	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("❌ Could not connect to Redis: %v", err)
	}

	fmt.Println("✅ Connected to Redis")
}

// Init initializes all connections
func Init() {
	ConnectDB()
	ConnectRedis()
	RunMigrations()
}
