package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jszymanowski/alive/fixtures"
	"github.com/jszymanowski/alive/models"
	"github.com/jszymanowski/alive/repositories"
	"github.com/jszymanowski/alive/utilities"
)

func TestUserRepository_FindByID_Found(t *testing.T) {
	db := utilities.SetupTestDB(t)

	testUser := fixtures.BuildUser()
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
	db := utilities.SetupTestDB(t)

	repo := repositories.NewUserRepository(db)

	user, err := repo.FindByID(9999999999)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "record not found")
}

func TestUserRepository_FindAll(t *testing.T) {
	db := utilities.SetupTestDB(t)

	testUsers := []*models.User{
		fixtures.BuildUser(models.User{Name: "Test User 1", Email: "test1@example.com"}),
		fixtures.BuildUser(models.User{Name: "Test User 2", Email: "test2@example.com"}),
	}
	result := db.Create(&testUsers)
	require.NoError(t, result.Error)

	repo := repositories.NewUserRepository(db)

	users, total, err := repo.FindAll(1, 10)
	require.NoError(t, err)

	assert.Len(t, users, 2)
	assert.Equal(t, int64(2), total)
	assert.Equal(t, "Test User 1", users[0].Name)
	assert.Equal(t, "test1@example.com", users[0].Email)
	assert.Equal(t, "Test User 2", users[1].Name)
	assert.Equal(t, "test2@example.com", users[1].Email)
}

func TestUserRepository_FindAll_WithNone(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewUserRepository(db)

	users, total, err := repo.FindAll(1, 10)
	require.NoError(t, err)

	assert.Len(t, users, 0)
	assert.Equal(t, int64(0), total)
}

func TestUserRepository_Create(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewUserRepository(db)

	newUser := fixtures.BuildUser()
	createdUser, err := repo.Create(newUser)
	require.NoError(t, err)
	assert.NotNil(t, createdUser)
}

func TestUserRepository_Create_InvalidName(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewUserRepository(db)

	newUser := &models.User{Email: "new@example.com"}
	createdUser, err := repo.Create(newUser)
	assert.Nil(t, createdUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag")

}

func TestUserRepository_Create_InvalidEmail(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewUserRepository(db)

	newUser := &models.User{Name: "Test User", Email: "newexamplecom"}
	createdUser, err := repo.Create(newUser)
	assert.Nil(t, createdUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag")
}

func TestUserRepository_Create_EmailExists(t *testing.T) {
	db := utilities.SetupTestDB(t)

	testUser := models.User{Name: "Existing User", Email: "test@example.com"}
	result := db.Create(&testUser)
	require.NoError(t, result.Error)

	repo := repositories.NewUserRepository(db)

	newUser := &models.User{Name: "Test User", Email: "test@example.com"}
	createdUser, err := repo.Create(newUser)
	assert.Error(t, err)
	assert.Nil(t, createdUser)
	assert.Contains(t, err.Error(), "UNIQUE constraint failed: users.email")

}
