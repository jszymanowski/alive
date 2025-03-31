package models

import (
	"time"

	"gorm.io/gorm"
)

type Feature struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Enabled     bool
	Slug        string `gorm:"uniqueIndex;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
