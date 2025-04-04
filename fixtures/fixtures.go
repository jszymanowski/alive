package fixtures

import (
	"github.com/jszymanowski/alive/models"
)

func BuildMonitor(options ...func(*models.Monitor)) *models.Monitor {
	monitor := &models.Monitor{
		Name:   "Some Monitor",
		Slug:   "",
		Status: "active",
		UserID: 1,
	}

	for _, option := range options {
		option(monitor)
	}

	return monitor
}

// functional options pattern
func WithDescription(description string) func(*models.Monitor) {
	return func(m *models.Monitor) {
		m.Description = description
	}
}

func WithName(name string) func(*models.Monitor) {
	return func(m *models.Monitor) {
		m.Name = name
	}
}

func WithSlug(slug string) func(*models.Monitor) {
	return func(m *models.Monitor) {
		m.Slug = slug
	}
}

func WithStatus(status string) func(*models.Monitor) {
	return func(m *models.Monitor) {
		m.Status = status
	}
}

func WithUserID(userID uint) func(*models.Monitor) {
	return func(m *models.Monitor) {
		m.UserID = userID
	}
}

func BuildUser(options ...func(*models.User)) *models.User {
	user := &models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	for _, option := range options {
		option(user)
	}
	return user
}

func WithUserName(name string) func(*models.User) {
	return func(u *models.User) {
		u.Name = name
	}
}

func WithUserEmail(email string) func(*models.User) {
	return func(u *models.User) {
		u.Email = email
	}
}
