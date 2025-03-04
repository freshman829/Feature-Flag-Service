package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"net/url"
	"strings"
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"feature-flag-service/internal/models"
)

// Global variables for DB and Redis
var (
	DB   *gorm.DB
	RDB  *redis.Client
	Ctx  = context.Background()
	Mock sqlmock.Sqlmock // Mock variable for testing
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
	if os.Getenv("TEST_MODE") == "true" {
		fmt.Println("🛠️ Running in TEST mode with a mock database")

		var err error
		var mockDB *sql.DB
		mockDB, Mock, err = sqlmock.New()
		if err != nil {
			log.Fatalf("❌ Failed to create SQL mock: %v", err)
		}

		DB, err = gorm.Open(postgres.New(postgres.Config{
			Conn:                 mockDB,
			PreferSimpleProtocol: true, // Disable prepared statements for simplicity
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true, // 🚀 Fix: Disable transactions in test mode
		})
		if err != nil {
			log.Fatalf("❌ Failed to connect to mock database: %v", err)
		}

		fmt.Println("✅ Mock database connected")
		return
	}

	// Production/PostgreSQL setup
	dsn := os.Getenv("DATABASE_URL")
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
	if os.Getenv("TEST_MODE") != "true" {
		RunMigrations()
	}
}