package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"feature-flag-service/internal/config"
	"feature-flag-service/internal/handlers"
	"feature-flag-service/internal/models"
)

// TestMain sets up the test environment before running tests
func TestMain(m *testing.M) {
	// Set test mode environment variable
	os.Setenv("TEST_MODE", "true")

	// Initialize database (test database)
	config.Init()

	// Run migrations for test database
	config.RunMigrations()

	// Run tests
	exitCode := m.Run()

	// Cleanup: truncate the tables after tests
	config.DB.Exec("TRUNCATE TABLE feature_flags RESTART IDENTITY CASCADE;")

	// Exit tests
	os.Exit(exitCode)
}

// setupTestRouter initializes a test Gin router with API routes
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/flags", handlers.CreateFeatureFlag)
		api.GET("/flags", handlers.GetFeatureFlags)
		api.GET("/flags/:id", handlers.GetFeatureFlag)
		api.PUT("/flags/:id", handlers.UpdateFeatureFlag)
		api.DELETE("/flags/:id", handlers.DeleteFeatureFlag)
	}

	return router
}

// TestCreateFeatureFlag checks if we can create a feature flag successfully
func TestCreateFeatureFlag(t *testing.T) {
	router := setupTestRouter()
	flag := models.FeatureFlag{
		Name:        "test_feature",
		Description: "A test feature",
		IsEnabled:   true,
	}
	body, _ := json.Marshal(flag)

	req, _ := http.NewRequest("POST", "/api/flags", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Check response body
	var response models.FeatureFlag
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "test_feature", response.Name)
}

// TestGetFeatureFlags verifies retrieving all feature flags
func TestGetFeatureFlags(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/flags", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetFeatureFlag ensures we can retrieve a single feature flag
func TestGetFeatureFlag(t *testing.T) {
	router := setupTestRouter()

	// Create a test feature flag
	testFlag := models.FeatureFlag{Name: "feature_1", Description: "Test flag", IsEnabled: true}
	config.DB.Create(&testFlag)

	req, _ := http.NewRequest("GET", "/api/flags/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestUpdateFeatureFlag checks if an existing flag can be updated
func TestUpdateFeatureFlag(t *testing.T) {
	router := setupTestRouter()

	// Create a test feature flag
	testFlag := models.FeatureFlag{Name: "feature_2", Description: "To be updated", IsEnabled: false}
	config.DB.Create(&testFlag)

	updateData := models.FeatureFlag{
		Name:        "updated_feature",
		Description: "Updated description",
		IsEnabled:   true,
	}
	body, _ := json.Marshal(updateData)

	req, _ := http.NewRequest("PUT", "/api/flags/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestDeleteFeatureFlag confirms deletion works correctly
func TestDeleteFeatureFlag(t *testing.T) {
	router := setupTestRouter()

	// Create a test feature flag
	testFlag := models.FeatureFlag{Name: "feature_3", Description: "To be deleted", IsEnabled: true}
	config.DB.Create(&testFlag)

	req, _ := http.NewRequest("DELETE", "/api/flags/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
