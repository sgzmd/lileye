package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lileye/backend/internal/models"
	"github.com/lileye/backend/internal/storage"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestHandler(t *testing.T) (*gin.Engine, *NotificationHandler) {
	gin.SetMode(gin.TestMode)
	
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Notification{})
	assert.NoError(t, err)

	storage := storage.NewNotificationStorage(db)
	handler := NewNotificationHandler(storage)
	
	r := gin.Default()
	handler.RegisterRoutes(r)

	return r, handler
}

func TestCreateNotification(t *testing.T) {
	r, _ := setupTestHandler(t)

	notification := models.Notification{
		Title:       "Test Title",
		Message:     "Test Message",
		Timestamp:   time.Now(),
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    "test123",
	}

	body, err := json.Marshal(notification)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/notifications", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Notification
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotZero(t, response.ID)
	assert.Equal(t, notification.Title, response.Title)
}

func TestGetNotification(t *testing.T) {
	r, h := setupTestHandler(t)

	// Create a test notification
	notification := models.Notification{
		Title:       "Test Title",
		Message:     "Test Message",
		Timestamp:   time.Now(),
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    "test123",
	}
	err := h.storage.Create(&notification)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/notifications/%d", notification.ID), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Notification
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, notification.ID, response.ID)
	assert.Equal(t, notification.Title, response.Title)
}

func TestGetNotificationsByDevice(t *testing.T) {
	r, h := setupTestHandler(t)

	// Create test notifications
	deviceID := "test123"
	notification1 := models.Notification{
		Title:       "Test Title 1",
		Message:     "Test Message 1",
		Timestamp:   time.Now(),
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    deviceID,
	}
	err := h.storage.Create(&notification1)
	assert.NoError(t, err)

	notification2 := models.Notification{
		Title:       "Test Title 2",
		Message:     "Test Message 2",
		Timestamp:   time.Now(),
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    deviceID,
	}
	err = h.storage.Create(&notification2)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/notifications/device/%s", deviceID), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Notification
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
}

func TestGetNotificationsByDateRange(t *testing.T) {
	r, h := setupTestHandler(t)

	now := time.Now()
	deviceID := "test123"
	notification := models.Notification{
		Title:       "Test Title",
		Message:     "Test Message",
		Timestamp:   now,
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    deviceID,
	}
	err := h.storage.Create(&notification)
	assert.NoError(t, err)

	start := now.Add(-time.Hour).Format(time.RFC3339)
	end := now.Add(time.Hour).Format(time.RFC3339)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/notifications/device/%s/range?start=%s&end=%s", deviceID, start, end), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Notification
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 1)
}

func TestSearchNotifications(t *testing.T) {
	r, h := setupTestHandler(t)

	deviceID := "test123"
	notification := models.Notification{
		Title:       "Test Title",
		Message:     "Test Message",
		Timestamp:   time.Now(),
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    deviceID,
	}
	err := h.storage.Create(&notification)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/notifications/device/%s/search?q=Test", deviceID), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Notification
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 1)
}

func TestGetDevices(t *testing.T) {
	r, h := setupTestHandler(t)

	// Create notifications for different devices
	notification1 := models.Notification{
		Title:       "Test Title 1",
		Message:     "Test Message 1",
		Timestamp:   time.Now(),
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    "device1",
	}
	err := h.storage.Create(&notification1)
	assert.NoError(t, err)

	notification2 := models.Notification{
		Title:       "Test Title 2",
		Message:     "Test Message 2",
		Timestamp:   time.Now(),
		PackageName: "com.test.app",
		From:        "Test User",
		DeviceID:    "device2",
	}
	err = h.storage.Create(&notification2)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/devices", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Contains(t, response, "device1")
	assert.Contains(t, response, "device2")
} 