package fixtures

import (
	"github.com/jszymanowski/alive/models"
)

func BuildMonitor(options ...func(*models.Monitor)) *models.Monitor {
	monitor := &models.Monitor{
		Name:   "Some Monitor",
		Slug:   "some-monitor-slug",
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

func BuildUser(overrides ...models.User) *models.User {
	user := &models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}

	for _, override := range overrides {
		if override.Name != "" {
			user.Name = override.Name
		}
		if override.Email != "" {
			user.Email = override.Email
		}
	}

	return user
}
