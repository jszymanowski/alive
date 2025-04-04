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

	monitors, total, err := repo.FindAll(1, 10)
	require.NoError(t, err)

	assert.Len(t, monitors, 2)
	assert.Equal(t, int64(2), total)
	assert.Equal(t, "Test Monitor 1", monitors[0].Name)
	assert.Equal(t, "test-monitor-1", monitors[0].Slug)
	assert.Equal(t, "Test Monitor 2", monitors[1].Name)
	assert.Equal(t, "test-monitor-2", monitors[1].Slug)
}

func TestMonitorRepository_FindAll_WithNone(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewMonitorRepository(db)

	monitors, total, err := repo.FindAll(1, 10)
	require.NoError(t, err)

	assert.Len(t, monitors, 0)
	assert.Equal(t, int64(0), total)
}

func TestMonitorRepository_Create(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewMonitorRepository(db)

	newMonitor := fixtures.BuildMonitor()
	createdMonitor, err := repo.Create(newMonitor)
	require.NoError(t, err)

	assert.NotNil(t, createdMonitor)
	assert.NotZero(t, createdMonitor.ID)
	assert.Equal(t, newMonitor.Name, createdMonitor.Name)
	assert.Equal(t, newMonitor.Slug, createdMonitor.Slug)
	assert.Equal(t, newMonitor.Status, createdMonitor.Status)
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

func TestMonitorRepository_Create_TooShortName(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewMonitorRepository(db)

	newMonitor := fixtures.BuildMonitor(fixtures.WithName("to"))
	createdMonitor, err := repo.Create(newMonitor)

	assert.Nil(t, createdMonitor)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'Monitor.Name' Error:Field validation for 'Name' failed on the 'min' tag")
}

func TestMonitorRepository_Create_TooShortSlug(t *testing.T) {
	db := utilities.SetupTestDB(t)
	repo := repositories.NewMonitorRepository(db)

	newMonitor := fixtures.BuildMonitor(fixtures.WithSlug("xy"))
	createdMonitor, err := repo.Create(newMonitor)

	assert.Nil(t, createdMonitor)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'Monitor.Slug' Error:Field validation for 'Slug' failed on the 'min' tag")
}
