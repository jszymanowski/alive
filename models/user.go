package models

type User struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
