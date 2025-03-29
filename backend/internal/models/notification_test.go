package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNotificationModel(t *testing.T) {
	// Create an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Auto migrate the schema
	err = db.AutoMigrate(&Notification{})
	assert.NoError(t, err)

	// Create a test notification
	now := time.Now()
	notification := &Notification{
		Title:       "Test Title",
		Message:     "Test Message",
		Timestamp:   now,
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    "test123",
	}

	// Test Create
	result := db.Create(notification)
	assert.NoError(t, result.Error)
	assert.NotZero(t, notification.ID)

	// Test Read
	var found Notification
	result = db.First(&found, notification.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, notification.Title, found.Title)
	assert.Equal(t, notification.Message, found.Message)
	assert.Equal(t, notification.PackageName, found.PackageName)
	assert.Equal(t, notification.From, found.From)
	assert.Equal(t, notification.DeviceID, found.DeviceID)

	// Test Update
	notification.Title = "Updated Title"
	result = db.Save(notification)
	assert.NoError(t, result.Error)

	// Verify Update
	result = db.First(&found, notification.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Updated Title", found.Title)

	// Test Delete
	result = db.Delete(&notification)
	assert.NoError(t, result.Error)

	// Verify Delete
	result = db.First(&found, notification.ID)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
} 