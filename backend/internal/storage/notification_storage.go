package storage

import (
	"time"

	"github.com/lileye/backend/internal/models"
	"gorm.io/gorm"
)

// NotificationStorage handles database operations for notifications
type NotificationStorage struct {
	db *gorm.DB
}

// NewNotificationStorage creates a new NotificationStorage instance
func NewNotificationStorage(db *gorm.DB) *NotificationStorage {
	return &NotificationStorage{db: db}
}

// Create stores a new notification in the database
func (s *NotificationStorage) Create(notification *models.Notification) error {
	return s.db.Create(notification).Error
}

// GetByID retrieves a notification by its ID
func (s *NotificationStorage) GetByID(id uint) (*models.Notification, error) {
	var notification models.Notification
	err := s.db.First(&notification, id).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

// GetByDeviceID retrieves notifications for a specific device
func (s *NotificationStorage) GetByDeviceID(deviceID string) ([]models.Notification, error) {
	var notifications []models.Notification
	err := s.db.Where("device_id = ?", deviceID).Find(&notifications).Error
	return notifications, err
}

// GetByDateRange retrieves notifications within a date range
func (s *NotificationStorage) GetByDateRange(deviceID string, start, end time.Time) ([]models.Notification, error) {
	var notifications []models.Notification
	err := s.db.Where("device_id = ? AND timestamp BETWEEN ? AND ?", deviceID, start, end).
		Order("timestamp desc").
		Find(&notifications).Error
	return notifications, err
}

// Search searches notifications by title, message, or from field
func (s *NotificationStorage) Search(deviceID, query string) ([]models.Notification, error) {
	var notifications []models.Notification
	err := s.db.Where("device_id = ? AND (title LIKE ? OR message LIKE ? OR \"from\" LIKE ?)",
		deviceID, "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Order("timestamp desc").
		Find(&notifications).Error
	return notifications, err
}

// GetDevices retrieves all unique device IDs
func (s *NotificationStorage) GetDevices() ([]string, error) {
	var devices []string
	err := s.db.Model(&models.Notification{}).Distinct().Pluck("device_id", &devices).Error
	return devices, err
}

// DeleteAll deletes all notifications from the database
func (s *NotificationStorage) DeleteAll() error {
	return s.db.Exec("DELETE FROM notifications").Error
} 