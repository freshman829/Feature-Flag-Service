package models

import (
	"time"

	"gorm.io/gorm"
)

// FeatureFlag represents a feature flag in the system
type FeatureFlag struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"unique;not null" json:"name"`
	Description string         `json:"description"`
	IsEnabled   bool           `json:"is_enabled"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
