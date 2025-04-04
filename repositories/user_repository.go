package repositories

import (
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
	return validate.Struct(user)
}

func (r *UserRepository) FindAll(page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	r.DB.Model(&models.User{}).Count(&total)

	offset := (page - 1) * pageSize
	result := r.DB.Offset(offset).Limit(pageSize).Find(&users)

	return users, total, result.Error
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
