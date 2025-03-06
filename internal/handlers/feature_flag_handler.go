package handlers

import (
	"net/http"
	"feature-flag-service/internal/config"
	"feature-flag-service/internal/models"

	"github.com/gin-gonic/gin"
)

// FeatureFlagRequest represents the expected body for creating a feature flag
type FeatureFlagRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	IsEnabled   bool   `json:"is_enabled"`
}

// CreateFeatureFlag handles creating a new feature flag
// @Summary Create a new feature flag
// @Description Adds a new feature flag to the system
// @Tags Feature Flags
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param featureFlag body FeatureFlagRequest true "Feature flag details"
// @Success 201 {object} models.FeatureFlag
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/flags [post]
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
// @Summary Get all feature flags
// @Description Retrieves all feature flags from the system
// @Tags Feature Flags
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.FeatureFlag
// @Failure 500 {object} map[string]string
// @Router /api/flags [get]
func GetFeatureFlags(c *gin.Context) {
	var featureFlags []models.FeatureFlag

	if err := config.DB.Find(&featureFlags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve feature flags"})
		return
	}

	c.JSON(http.StatusOK, featureFlags)
}

// GetFeatureFlag retrieves a specific feature flag by ID
// @Summary Get a feature flag by ID
// @Description Retrieves details of a specific feature flag
// @Tags Feature Flags
// @Produce json
// @Security BearerAuth
// @Param id path int true "Feature flag ID"
// @Success 200 {object} models.FeatureFlag
// @Failure 404 {object} map[string]string
// @Router /api/flags/{id} [get]
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
// @Summary Update a feature flag
// @Description Updates the details of an existing feature flag
// @Tags Feature Flags
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Feature flag ID"
// @Param featureFlag body models.FeatureFlag true "Updated feature flag details"
// @Success 200 {object} models.FeatureFlag
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/flags/{id} [put]
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
// @Summary Delete a feature flag
// @Description Deletes a feature flag from the system
// @Tags Feature Flags
// @Security BearerAuth
// @Param id path int true "Feature flag ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/flags/{id} [delete]
func DeleteFeatureFlag(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.FeatureFlag{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete feature flag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feature flag deleted successfully"})
}
