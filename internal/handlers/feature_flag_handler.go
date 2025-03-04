package handlers

import (
	"net/http"
	"feature-flag-service/internal/config"
	"feature-flag-service/internal/models"

	"github.com/gin-gonic/gin"
)

// CreateFeatureFlag handles creating a new feature flag
func CreateFeatureFlag(c *gin.Context) {
	var featureFlag models.FeatureFlag

	if err := c.ShouldBindJSON(&featureFlag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&featureFlag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create feature flag"})
		return
	}

	c.JSON(http.StatusCreated, featureFlag)
}

// GetFeatureFlags retrieves all feature flags
func GetFeatureFlags(c *gin.Context) {
	var featureFlags []models.FeatureFlag

	if err := config.DB.Find(&featureFlags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve feature flags"})
		return
	}

	c.JSON(http.StatusOK, featureFlags)
}

// GetFeatureFlag retrieves a specific feature flag by ID
func GetFeatureFlag(c *gin.Context) {
	var featureFlag models.FeatureFlag

	id := c.Param("id")
	if err := config.DB.First(&featureFlag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feature flag not found"})
		return
	}

	c.JSON(http.StatusOK, featureFlag)
}

// UpdateFeatureFlag updates an existing feature flag
func UpdateFeatureFlag(c *gin.Context) {
	var featureFlag models.FeatureFlag
	id := c.Param("id")

	if err := config.DB.First(&featureFlag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feature flag not found"})
		return
	}

	if err := c.ShouldBindJSON(&featureFlag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&featureFlag)
	c.JSON(http.StatusOK, featureFlag)
}

// DeleteFeatureFlag deletes a feature flag
func DeleteFeatureFlag(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.FeatureFlag{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete feature flag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feature flag deleted successfully"})
}
