package storage

import (
	"testing"
	"time"

	"github.com/lileye/backend/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, *NotificationStorage) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Notification{})
	assert.NoError(t, err)

	return db, NewNotificationStorage(db)
}

func createTestNotification(t *testing.T, storage *NotificationStorage, deviceID string) *models.Notification {
	notification := &models.Notification{
		Title:       "Test Title",
		Message:     "Test Message",
		Timestamp:   time.Now(),
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    deviceID,
	}

	err := storage.Create(notification)
	assert.NoError(t, err)
	return notification
}

func TestNotificationStorage_Create(t *testing.T) {
	_, storage := setupTestDB(t)
	notification := createTestNotification(t, storage, "device1")
	assert.NotZero(t, notification.ID)
}

func TestNotificationStorage_GetByID(t *testing.T) {
	_, storage := setupTestDB(t)
	created := createTestNotification(t, storage, "device1")

	found, err := storage.GetByID(created.ID)
	assert.NoError(t, err)
	assert.Equal(t, created.Title, found.Title)
	assert.Equal(t, created.DeviceID, found.DeviceID)
}

func TestNotificationStorage_GetByDeviceID(t *testing.T) {
	_, storage := setupTestDB(t)
	notification1 := createTestNotification(t, storage, "device1")
	notification2 := createTestNotification(t, storage, "device1")
	_ = createTestNotification(t, storage, "device2")

	notifications, err := storage.GetByDeviceID("device1")
	assert.NoError(t, err)
	assert.Len(t, notifications, 2)
	assert.Contains(t, []uint{notification1.ID, notification2.ID}, notifications[0].ID)
}

func TestNotificationStorage_GetByDateRange(t *testing.T) {
	_, storage := setupTestDB(t)
	now := time.Now()
	
	notification := &models.Notification{
		Title:       "Test Title",
		Message:     "Test Message",
		Timestamp:   now,
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    "device1",
	}
	err := storage.Create(notification)
	assert.NoError(t, err)

	notifications, err := storage.GetByDateRange("device1", now.Add(-time.Hour), now.Add(time.Hour))
	assert.NoError(t, err)
	assert.Len(t, notifications, 1)
	assert.Equal(t, notification.ID, notifications[0].ID)
}

func TestNotificationStorage_Search(t *testing.T) {
	_, storage := setupTestDB(t)
	notification := createTestNotification(t, storage, "device1")

	// Search by title
	results, err := storage.Search("device1", "Test Title")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, notification.ID, results[0].ID)

	// Search by message
	results, err = storage.Search("device1", "Test Message")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, notification.ID, results[0].ID)

	// Search by from
	results, err = storage.Search("device1", "Test User")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, notification.ID, results[0].ID)
}

func TestNotificationStorage_GetDevices(t *testing.T) {
	_, storage := setupTestDB(t)
	_ = createTestNotification(t, storage, "device1")
	_ = createTestNotification(t, storage, "device2")
	_ = createTestNotification(t, storage, "device1")

	devices, err := storage.GetDevices()
	assert.NoError(t, err)
	assert.Len(t, devices, 2)
	assert.Contains(t, devices, "device1")
	assert.Contains(t, devices, "device2")
} 