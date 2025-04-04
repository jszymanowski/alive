package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	MonitorStatusActive   = "active"
	MonitorStatusInactive = "inactive"
	MonitorStatusPending  = "pending"
)

type Monitor struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null" validate:"required,min=3"`
	Description string
	Slug        string `gorm:"uniqueIndex" validate:"required,min=3"`
	Status      string `gorm:"not null"  validate:"required,oneof=active inactive pending"`
	UserID      uint   `gorm:"not null;foreignKey:ID;references:users"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
