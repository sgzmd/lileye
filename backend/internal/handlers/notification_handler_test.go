package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lileye/backend/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockNotificationStorage is a mock implementation of NotificationStorageInterface
type MockNotificationStorage struct {
	mock.Mock
}

func (m *MockNotificationStorage) Create(notification *models.Notification) error {
	args := m.Called(notification)
	return args.Error(0)
}

func (m *MockNotificationStorage) GetByID(id uint) (*models.Notification, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Notification), args.Error(1)
}

func (m *MockNotificationStorage) GetByDeviceID(deviceID string) ([]models.Notification, error) {
	args := m.Called(deviceID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Notification), args.Error(1)
}

func (m *MockNotificationStorage) GetByDateRange(deviceID string, start, end time.Time) ([]models.Notification, error) {
	args := m.Called(deviceID, start, end)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Notification), args.Error(1)
}

func (m *MockNotificationStorage) Search(deviceID, query string) ([]models.Notification, error) {
	args := m.Called(deviceID, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Notification), args.Error(1)
}

func (m *MockNotificationStorage) GetDevices() ([]models.Device, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Device), args.Error(1)
}

func (m *MockNotificationStorage) DeleteAll() error {
	args := m.Called()
	return args.Error(0)
}

func setupTestRouter() (*gin.Engine, *MockNotificationStorage) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	mockStorage := new(MockNotificationStorage)
	handler := NewNotificationHandler(mockStorage)
	handler.RegisterRoutes(r)
	return r, mockStorage
}

func TestCreateNotification(t *testing.T) {
	r, mockStorage := setupTestRouter()

	notification := models.Notification{
		DeviceID:  "test-device",
		Title:     "Test Notification",
		Message:   "Test Message",
		From:      "Test App",
		Timestamp: time.Now(),
	}

	mockStorage.On("Create", mock.AnythingOfType("*models.Notification")).Return(nil)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(notification)
	req, _ := http.NewRequest("POST", "/api/notifications", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockStorage.AssertExpectations(t)
}

func TestGetNotification(t *testing.T) {
	r, mockStorage := setupTestRouter()

	notification := &models.Notification{
		Model:     gorm.Model{ID: 1},
		DeviceID:  "test-device",
		Title:     "Test Notification",
		Message:   "Test Message",
		From:      "Test App",
		Timestamp: time.Now(),
	}

	mockStorage.On("GetByID", uint(1)).Return(notification, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/notifications/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockStorage.AssertExpectations(t)
}

func TestGetNotificationsByDevice(t *testing.T) {
	r, mockStorage := setupTestRouter()

	notifications := []models.Notification{
		{
			Model:     gorm.Model{ID: 1},
			DeviceID:  "test-device",
			Title:     "Test Notification 1",
			Message:   "Test Message 1",
			From:      "Test App",
			Timestamp: time.Now(),
		},
		{
			Model:     gorm.Model{ID: 2},
			DeviceID:  "test-device",
			Title:     "Test Notification 2",
			Message:   "Test Message 2",
			From:      "Test App",
			Timestamp: time.Now(),
		},
	}

	mockStorage.On("GetByDeviceID", "test-device").Return(notifications, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/notifications/device/test-device", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockStorage.AssertExpectations(t)
}

func TestGetNotificationsByDateRange(t *testing.T) {
	r, mockStorage := setupTestRouter()

	start := time.Now().Add(-24 * time.Hour)
	end := time.Now()

	notifications := []models.Notification{
		{
			Model:     gorm.Model{ID: 1},
			DeviceID:  "test-device",
			Title:     "Test Notification",
			Message:   "Test Message",
			From:      "Test App",
			Timestamp: time.Now(),
		},
	}

	mockStorage.On("GetByDateRange", "test-device", mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(notifications, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/notifications/device/test-device/range?start="+start.Format(time.RFC3339)+"&end="+end.Format(time.RFC3339), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockStorage.AssertExpectations(t)
}

func TestSearchNotifications(t *testing.T) {
	r, mockStorage := setupTestRouter()

	notifications := []models.Notification{
		{
			Model:     gorm.Model{ID: 1},
			DeviceID:  "test-device",
			Title:     "Test Notification",
			Message:   "Test Message",
			From:      "Test App",
			Timestamp: time.Now(),
		},
	}

	mockStorage.On("Search", "test-device", "test").Return(notifications, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/notifications/device/test-device/search?q=test", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockStorage.AssertExpectations(t)
}

func TestDeleteAllNotifications(t *testing.T) {
	r, mockStorage := setupTestRouter()

	mockStorage.On("DeleteAll").Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/notifications/all", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockStorage.AssertExpectations(t)
} 