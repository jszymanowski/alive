package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/jszymanowski/alive/models"
	"github.com/jszymanowski/alive/repositories"
)

func TestUserRepository_WithSQLite(t *testing.T) {
	// Setup in-memory SQLite database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)

	// Auto migrate schema
	err = db.AutoMigrate(&models.User{})
	require.NoError(t, err)

	// Setup test data
	testUser := models.User{Name: "Test User", Email: "test@example.com"}
	result := db.Create(&testUser)
	require.NoError(t, result.Error)

	// Initialize repository with SQLite DB
	repo := repositories.NewUserRepository(db)

	// Test FindByID
	user, err := repo.FindByID(testUser.ID)
	require.NoError(t, err)
	assert.Equal(t, testUser.ID, user.ID)
	assert.Equal(t, "Test User", user.Name)
	assert.Equal(t, "test@example.com", user.Email)
}
