package storage

import (
	"testing"
	"time"

	"github.com/lileye/backend/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Notification{}, &models.Device{})
	assert.NoError(t, err)

	return db
}

func TestNotificationStorage_Create(t *testing.T) {
	db := setupTestDB(t)
	storage := NewNotificationStorage(db).(*NotificationStorage)

	notification := &models.Notification{
		DeviceID:  "test-device",
		Title:     "Test Notification",
		Message:   "Test Message",
		From:      "Test App",
		Timestamp: time.Now(),
	}

	err := storage.Create(notification)
	assert.NoError(t, err)

	var result models.Notification
	err = db.First(&result, notification.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, notification.DeviceID, result.DeviceID)
	assert.Equal(t, notification.Title, result.Title)
	assert.Equal(t, notification.Message, result.Message)
	assert.Equal(t, notification.From, result.From)
}

func TestNotificationStorage_GetByID(t *testing.T) {
	db := setupTestDB(t)
	storage := NewNotificationStorage(db).(*NotificationStorage)

	notification := &models.Notification{
		DeviceID:  "test-device",
		Title:     "Test Notification",
		Message:   "Test Message",
		From:      "Test App",
		Timestamp: time.Now(),
	}

	err := storage.Create(notification)
	assert.NoError(t, err)

	result, err := storage.GetByID(notification.ID)
	assert.NoError(t, err)
	assert.Equal(t, notification.DeviceID, result.DeviceID)
	assert.Equal(t, notification.Title, result.Title)
	assert.Equal(t, notification.Message, result.Message)
	assert.Equal(t, notification.From, result.From)
}

func TestNotificationStorage_GetByDeviceID(t *testing.T) {
	db := setupTestDB(t)
	storage := NewNotificationStorage(db).(*NotificationStorage)

	notifications := []models.Notification{
		{
			DeviceID:  "test-device",
			Title:     "Test Notification 1",
			Message:   "Test Message 1",
			From:      "Test App",
			Timestamp: time.Now(),
		},
		{
			DeviceID:  "test-device",
			Title:     "Test Notification 2",
			Message:   "Test Message 2",
			From:      "Test App",
			Timestamp: time.Now(),
		},
	}

	for _, notification := range notifications {
		err := storage.Create(&notification)
		assert.NoError(t, err)
	}

	results, err := storage.GetByDeviceID("test-device")
	assert.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestNotificationStorage_GetByDateRange(t *testing.T) {
	db := setupTestDB(t)
	storage := NewNotificationStorage(db).(*NotificationStorage)

	now := time.Now()
	notifications := []models.Notification{
		{
			DeviceID:  "test-device",
			Title:     "Test Notification 1",
			Message:   "Test Message 1",
			From:      "Test App",
			Timestamp: now.Add(-1 * time.Hour),
		},
		{
			DeviceID:  "test-device",
			Title:     "Test Notification 2",
			Message:   "Test Message 2",
			From:      "Test App",
			Timestamp: now.Add(1 * time.Hour),
		},
	}

	for _, notification := range notifications {
		err := storage.Create(&notification)
		assert.NoError(t, err)
	}

	results, err := storage.GetByDateRange("test-device", now.Add(-2*time.Hour), now.Add(2*time.Hour))
	assert.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestNotificationStorage_Search(t *testing.T) {
	db := setupTestDB(t)
	storage := NewNotificationStorage(db).(*NotificationStorage)

	notifications := []models.Notification{
		{
			DeviceID:  "test-device",
			Title:     "Test Notification",
			Message:   "Test Message",
			From:      "Test App",
			Timestamp: time.Now(),
		},
	}

	for _, notification := range notifications {
		err := storage.Create(&notification)
		assert.NoError(t, err)
	}

	results, err := storage.Search("test-device", "Test")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}

func TestNotificationStorage_GetDevices(t *testing.T) {
	db := setupTestDB(t)
	storage := NewNotificationStorage(db).(*NotificationStorage)

	// Create a device with a custom name
	device := &models.Device{
		DeviceID:   "test-device",
		DeviceName: "Test Device",
	}
	err := db.Create(device).Error
	assert.NoError(t, err)

	// Create notifications for the device
	notifications := []models.Notification{
		{
			DeviceID:  "test-device",
			Title:     "Test Notification",
			Message:   "Test Message",
			From:      "Test App",
			Timestamp: time.Now(),
		},
	}

	for _, notification := range notifications {
		err := storage.Create(&notification)
		assert.NoError(t, err)
	}

	devices, err := storage.GetDevices()
	assert.NoError(t, err)
	assert.Len(t, devices, 1)
	assert.Equal(t, "test-device", devices[0].DeviceID)
	assert.Equal(t, "Test Device", devices[0].DeviceName)
}

func TestNotificationStorage_DeleteAll(t *testing.T) {
	db := setupTestDB(t)
	storage := NewNotificationStorage(db).(*NotificationStorage)

	notifications := []models.Notification{
		{
			DeviceID:  "test-device",
			Title:     "Test Notification",
			Message:   "Test Message",
			From:      "Test App",
			Timestamp: time.Now(),
		},
	}

	for _, notification := range notifications {
		err := storage.Create(&notification)
		assert.NoError(t, err)
	}

	err := storage.DeleteAll()
	assert.NoError(t, err)

	var count int64
	err = db.Model(&models.Notification{}).Count(&count).Error
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
} 