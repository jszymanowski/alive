package repositories

import (
	"strings"

	"gorm.io/gorm"

	"github.com/jszymanowski/alive/models"
)

type MonitorRepository struct {
	DB *gorm.DB
}

func NewMonitorRepository(db *gorm.DB) *MonitorRepository {
	return &MonitorRepository{DB: db}
}

func ValidateMonitor(monitor *models.Monitor) error {
	return validate.Struct(monitor)
}

func (r *MonitorRepository) FindAll(page, pageSize int) ([]models.Monitor, int64, error) {
	var monitors []models.Monitor
	var total int64

	r.DB.Model(&models.Monitor{}).Count(&total)

	offset := (page - 1) * pageSize
	result := r.DB.Offset(offset).Limit(pageSize).Find(&monitors)

	return monitors, total, result.Error
}

func (r *MonitorRepository) FindByID(id uint) (*models.Monitor, error) {
	var monitor models.Monitor
	result := r.DB.Take(&monitor, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &monitor, nil
}

func (r *MonitorRepository) Create(monitor *models.Monitor) (*models.Monitor, error) {
	if err := ValidateMonitor(monitor); err != nil {
		return nil, err
	}

	// Generate slug if not provided
	// TODO: TestMonitorRepository_Create trips validation error, unless slug is provided -- it should be generated
	if monitor.Slug == "" {
		monitor.Slug = generateSlug(monitor.Name)
	}

	result := r.DB.Create(monitor)
	if result.Error != nil {
		return nil, result.Error
	}
	return monitor, nil
}

func generateSlug(name string) string {
	// Replace spaces with hyphens, remove special characters,
	// convert to lowercase, etc.
	// This is a placeholder - implement a proper slug generator
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
