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
func TestMain(m *testing.M) {
	// Set test mode environment variable
	os.Setenv("TEST_MODE", "true")

	// Initialize database with mock
	config.ConnectDB()

	// Run tests
	code := m.Run()

	// Exit with the correct status code
	os.Exit(code)
}

func TestCreateFeatureFlag(t *testing.T) {
	flag := models.FeatureFlag{
		Name:        "test_feature",
		Description: "A test feature",
		IsEnabled:   true,
	}

	// Expect a successful insert
	config.Mock.ExpectExec("INSERT INTO feature_flags").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := config.DB.Create(&flag).Error
	assert.NoError(t, err)
}

func TestGetFeatureFlags(t *testing.T) {
	// Mock query result
	rows := sqlmock.NewRows([]string{"id", "name", "description", "is_enabled"}).
		AddRow(1, "test_feature", "A test feature", true)

	config.Mock.ExpectQuery("SELECT (.+) FROM feature_flags").WillReturnRows(rows)

	var flags []models.FeatureFlag
	err := config.DB.Find(&flags).Error
	assert.NoError(t, err)
	assert.Len(t, flags, 1)
	assert.Equal(t, "test_feature", flags[0].Name)
}