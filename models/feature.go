package models

import (
	"time"

	"gorm.io/gorm"
)

type Feature struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Enabled     bool
	Slug        string `gorm:"uniqueIndex"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
