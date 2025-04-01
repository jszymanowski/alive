package repositories

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/jszymanowski/alive/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func ValidateUser(user *models.User) error {
	validate := validator.New()
	return validate.Struct(user)
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.DB.Find(&users)
	return users, result.Error
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := r.DB.Take(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := ValidateUser(user); err != nil {
		return nil, err
	}

	result := r.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
