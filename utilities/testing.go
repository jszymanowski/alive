package utilities

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/jszymanowski/alive/models"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&models.User{}, &models.Monitor{})
	require.NoError(t, err)

	t.Cleanup(func() {
		sqlDB, _ := db.DB()
		err := sqlDB.Close()
		require.NoError(t, err)
	})

	return db
}
