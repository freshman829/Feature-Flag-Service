package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"feature-flag-service/internal/config"
	"feature-flag-service/internal/handlers"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}

	// Initialize database and Redis
	config.Init()

	// Create a new Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		api.POST("/flags", handlers.CreateFeatureFlag)
		api.GET("/flags", handlers.GetFeatureFlags)
		api.GET("/flags/:id", handlers.GetFeatureFlag)
		api.PUT("/flags/:id", handlers.UpdateFeatureFlag)
		api.DELETE("/flags/:id", handlers.DeleteFeatureFlag)
	}

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Server running on port %s\n", port)
	r.Run(":" + port)
}
