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

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&models.User{})
	require.NoError(t, err)

	t.Cleanup(func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})

	return db
}

func TestUserRepository_FindByID_Found(t *testing.T) {
	db := SetupTestDB(t)

	testUser := models.User{Name: "Test User", Email: "test@example.com"}
	result := db.Create(&testUser)
	require.NoError(t, result.Error)

	repo := repositories.NewUserRepository(db)

	user, err := repo.FindByID(testUser.ID)
	require.NoError(t, err)
	assert.Equal(t, testUser.ID, user.ID)
	assert.Equal(t, "Test User", user.Name)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestUserRepository_FindByID_NotFound(t *testing.T) {
	db := SetupTestDB(t)

	repo := repositories.NewUserRepository(db)

	user, err := repo.FindByID(9999999999)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUserRepository_FindAll(t *testing.T) {
	db := SetupTestDB(t)

	testUsers := []*models.User{
		{Name: "Test User 1", Email: "test1@example.com"},
		{Name: "Test User 2", Email: "test2@example.com"},
	}
	result := db.Create(&testUsers)
	require.NoError(t, result.Error)

	repo := repositories.NewUserRepository(db)

	users, err := repo.FindAll()
	require.NoError(t, err)

	assert.Len(t, users, 2)
	assert.Equal(t, "Test User 1", users[0].Name)
	assert.Equal(t, "test1@example.com", users[0].Email)
	assert.Equal(t, "Test User 2", users[1].Name)
	assert.Equal(t, "test2@example.com", users[1].Email)
}

func TestUserRepository_FindAll_WithNone(t *testing.T) {
	db := SetupTestDB(t)
	repo := repositories.NewUserRepository(db)

	users, err := repo.FindAll()
	require.NoError(t, err)

	assert.Len(t, users, 0)
}

func TestUserRepository_Create(t *testing.T) {
	db := SetupTestDB(t)
	repo := repositories.NewUserRepository(db)

	newUser := &models.User{Name: "New User", Email: "new@example.com"}
	createdUser, err := repo.Create(newUser)
	require.NoError(t, err)
	assert.NotNil(t, createdUser)
}

func TestUserRepository_Create_Invalid(t *testing.T) {
	db := SetupTestDB(t)
	repo := repositories.NewUserRepository(db)

	newUser := &models.User{Email: "new@example.com"}
	createdUser, err := repo.Create(newUser)
	assert.Nil(t, createdUser)
	assert.Error(t, err)
}

func TestUserRepository_Create_EmailExists(t *testing.T) {
	db := SetupTestDB(t)

	testUser := models.User{Name: "Existing User", Email: "test@example.com"}
	result := db.Create(&testUser)
	require.NoError(t, result.Error)

	repo := repositories.NewUserRepository(db)

	newUser := &models.User{Email: "test@example.com"}
	createdUser, err := repo.Create(newUser)
	assert.Error(t, err)
	assert.Nil(t, createdUser)
}
