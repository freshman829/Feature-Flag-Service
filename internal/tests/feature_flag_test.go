package tests

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"feature-flag-service/internal/config"
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

	// Expecting a query (not Exec) because GORM auto-appends RETURNING "id" in PostgreSQL
	config.Mock.ExpectQuery(`INSERT INTO "feature_flags" \("name","description","is_enabled","created_at","updated_at","deleted_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6\) RETURNING "id"`).
		WithArgs(flag.Name, flag.Description, flag.IsEnabled, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	err := config.DB.Create(&flag).Error
	assert.NoError(t, err)

	// Ensure all expectations were met
	err = config.Mock.ExpectationsWereMet()
	assert.NoError(t, err)
}


func TestGetFeatureFlags(t *testing.T) {
	// Mock query result
	rows := sqlmock.NewRows([]string{"id", "name", "description", "is_enabled"}).
		AddRow(1, "test_feature", "A test feature", true)

	config.Mock.ExpectQuery(`SELECT \* FROM "feature_flags" WHERE "feature_flags"."deleted_at" IS NULL`).
		WillReturnRows(rows)

	var flags []models.FeatureFlag
	err := config.DB.Find(&flags).Error
	assert.NoError(t, err)
	assert.Len(t, flags, 1)
	assert.Equal(t, "test_feature", flags[0].Name)

	// Ensure all expectations were met
	err = config.Mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
