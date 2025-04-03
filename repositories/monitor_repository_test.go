package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jszymanowski/alive/models"
	"github.com/jszymanowski/alive/repositories"
	"github.com/jszymanowski/alive/utilities"
)

func TestMonitorRepository_FindByID_Found(t *testing.T) {
	db := utilities.SetupTestDB(t)

	testMonitor := models.Monitor{Name: "Test Monitor", Email: "test@example.com"}
	result := db.Create(&testMonitor)
	require.NoError(t, result.Error)

	repo := repositories.NewMonitorRepository(db)

	monitor, err := repo.FindByID(testMonitor.ID)
	require.NoError(t, err)
	assert.Equal(t, testMonitor.ID, monitor.ID)
	assert.Equal(t, "Test Monitor", monitor.Name)
	assert.Equal(t, "test@example.com", monitor.Email)
}

func TestMonitorRepository_FindByID_NotFound(t *testing.T) {
	db := utilities.SetupTestDB(t)

	repo := repositories.NewMonitorRepository(db)

	monitor, err := repo.FindByID(9999999999)
	assert.Error(t, err)
	assert.Nil(t, monitor)
	assert.Contains(t, err.Error(), "record not found")
}

func TestMonitorRepository_FindAll(t *testing.T) {
	db := utilities.SetupTestDB(t)

	testMonitors := []*models.Monitor{
		{Name: "Test Monitor 1", Email: "test1@example.com"},
		{Name: "Test Monitor 2", Email: "test2@example.com"},
	}
	result := db.Create(&testMonitors)
	require.NoError(t, result.Error)

	repo := repositories.NewMonitorRepository(db)

	monitors, err := repo.FindAll()
	require.NoError(t, err)

	assert.Len(t, monitors, 2)
	assert.Equal(t, "Test Monitor 1", monitors[0].Name)
	assert.Equal(t, "test1@example.com", monitors[0].Email)
	assert.Equal(t, "Test Monitor 2", monitors[1].Name)
	assert.Equal(t, "test2@example.com", monitors[1].Email)
}

func TestMonitorRepository_FindAll_WithNone(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewMonitorRepository(db)

	monitors, err := repo.FindAll()
	require.NoError(t, err)

	assert.Len(t, monitors, 0)
}

func TestMonitorRepository_Create(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewMonitorRepository(db)

	newMonitor := &models.Monitor{Name: "New Monitor", Email: "new@example.com"}
	createdMonitor, err := repo.Create(newMonitor)
	require.NoError(t, err)
	assert.NotNil(t, createdMonitor)
}

func TestMonitorRepository_Create_InvalidName(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewMonitorRepository(db)

	newMonitor := &models.Monitor{Email: "new@example.com"}
	createdMonitor, err := repo.Create(newMonitor)
	assert.Nil(t, createdMonitor)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'Monitor.Name' Error:Field validation for 'Name' failed on the 'required' tag")

}

func TestMonitorRepository_Create_InvalidEmail(t *testing.T) {
	db := SetupTestDB(t)
	repo := repositories.NewMonitorRepository(db)

	newMonitor := &models.Monitor{Name: "Test Monitor", Email: "newexamplecom"}
	createdMonitor, err := repo.Create(newMonitor)
	assert.Nil(t, createdMonitor)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'Monitor.Email' Error:Field validation for 'Email' failed on the 'email' tag")
}

func TestMonitorRepository_Create_EmailExists(t *testing.T) {
	db := SetupTestDB(t)

	testMonitor := models.Monitor{Name: "Existing Monitor", Email: "test@example.com"}
	result := db.Create(&testMonitor)
	require.NoError(t, result.Error)

	repo := repositories.NewMonitorRepository(db)

	newMonitor := &models.Monitor{Name: "Test Monitor", Email: "test@example.com"}
	createdMonitor, err := repo.Create(newMonitor)
	assert.Error(t, err)
	assert.Nil(t, createdMonitor)
	assert.Contains(t, err.Error(), "UNIQUE constraint failed: monitors.email")

}
