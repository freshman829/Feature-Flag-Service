package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	_ "feature-flag-service/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"feature-flag-service/internal/config"
	"feature-flag-service/internal/handlers"
	"feature-flag-service/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Feature Flag Service API
// @version 1.0
// @description API for managing feature flags
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your token with "Bearer " prefix: Bearer <your_token>
// @host localhost:8080
// @BasePath /
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}

	// Initialize database and Redis
	config.Init()

	// Create a new Gin router
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins (Change this to specific domains in production)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
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
