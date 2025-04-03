package models

import (
	"time"

	"gorm.io/gorm"
)

type Monitor struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null" validate:"required,min=3"`
	Description string
	Slug        string `gorm:"uniqueIndex validate:"required,min=3"`
	Status      string `gorm:"not null"` // TODO: add enum for status
	UserID      uint   `gorm:"not null;foreignKey:ID;references:users"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
