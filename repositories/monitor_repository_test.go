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

func TestMonitorRepository_FindByID_Found(t *testing.T) {
	db := utilities.SetupTestDB(t)

	testMonitor := fixtures.BuildMonitor()
	result := db.Create(&testMonitor)
	require.NoError(t, result.Error)

	repo := repositories.NewMonitorRepository(db)

	monitor, err := repo.FindByID(testMonitor.ID)
	require.NoError(t, err)
	assert.Equal(t, testMonitor.ID, monitor.ID)
	assert.Equal(t, "Some Monitor", monitor.Name)
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
		fixtures.BuildMonitor(fixtures.WithName("Test Monitor 1"), fixtures.WithSlug("test-monitor-1")),
		fixtures.BuildMonitor(fixtures.WithName("Test Monitor 2"), fixtures.WithSlug("test-monitor-2")),
	}
	result := db.Create(&testMonitors)
	require.NoError(t, result.Error)

	repo := repositories.NewMonitorRepository(db)

	monitors, err := repo.FindAll()
	require.NoError(t, err)

	assert.Len(t, monitors, 2)
	assert.Equal(t, "Test Monitor 1", monitors[0].Name)
	assert.Equal(t, "Test Monitor 2", monitors[1].Name)
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

	newMonitor := fixtures.BuildMonitor()
	createdMonitor, err := repo.Create(newMonitor)
	require.NoError(t, err)
	assert.NotNil(t, createdMonitor)
}

func TestMonitorRepository_Create_InvalidName(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewMonitorRepository(db)

	newMonitor := fixtures.BuildMonitor(fixtures.WithName(""))

	createdMonitor, err := repo.Create(newMonitor)

	assert.Nil(t, createdMonitor)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'Monitor.Name' Error:Field validation for 'Name' failed on the 'required' tag")
}
