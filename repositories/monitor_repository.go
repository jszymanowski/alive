package repositories

import (
	"github.com/go-playground/validator/v10"
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
	validate := validator.New()
	return validate.Struct(monitor)
}

func (r *MonitorRepository) FindAll() ([]models.Monitor, error) {
	var monitors []models.Monitor
	result := r.DB.Find(&monitors)
	return monitors, result.Error
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

	result := r.DB.Create(monitor)
	if result.Error != nil {
		return nil, result.Error
	}
	return monitor, nil
}
