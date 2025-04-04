package utilities

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/jszymanowski/alive/models"
)

// SetupTestDB initializes an in-memory SQLite database for testing purposes.
// It automatically migrates the User and Monitor models and registers a cleanup
// function to close the database connection after the test completes.
//
// Example usage:
//
//	func TestSomething(t *testing.T) {
//	  db := utilities.SetupTestDB(t)
//	  // Use db for your tests
//	}
func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&models.User{}, &models.Monitor{})
	require.NoError(t, err)

	t.Cleanup(func() {
		sqlDB, err := db.DB()
		require.NoError(t, err)

		closeErr := sqlDB.Close()
		require.NoError(t, closeErr)
	})

	return db
}
