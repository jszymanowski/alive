package models

type User struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `json:"name" gorm:"not null" validate:"required,min=3"`
	Email string `json:"email" gorm:"uniqueIndex" validate:"required,email"`
}
