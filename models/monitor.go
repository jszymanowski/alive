package models

import (
	"time"

	"gorm.io/gorm"
)

type Monitor struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Slug        string `gorm:"uniqueIndex"`
	Status      string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
