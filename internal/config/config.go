package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global variables for DB and Redis
var (
	DB  *gorm.DB
	RDB *redis.Client
	Ctx = context.Background()
)

// ConnectDB initializes PostgreSQL connection
func ConnectDB() {
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
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // No password for now
		DB:       0,
	})

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
}
