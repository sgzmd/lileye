package storage

import (
	"github.com/lileye/backend/internal/models"
	"gorm.io/gorm"
)

// DeviceStorage handles database operations for devices
type DeviceStorage struct {
	db *gorm.DB
}

// NewDeviceStorage creates a new DeviceStorage instance
func NewDeviceStorage(db *gorm.DB) *DeviceStorage {
	return &DeviceStorage{db: db}
}

// CreateOrUpdate creates a new device or updates an existing one
func (s *DeviceStorage) CreateOrUpdate(device *models.Device) error {
	return s.db.Save(device).Error
}

// GetByDeviceID retrieves a device by its ID
func (s *DeviceStorage) GetByDeviceID(deviceID string) (*models.Device, error) {
	var device models.Device
	err := s.db.Where("device_id = ?", deviceID).First(&device).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

// GetAll retrieves all devices
func (s *DeviceStorage) GetAll() ([]models.Device, error) {
	var devices []models.Device
	err := s.db.Find(&devices).Error
	return devices, err
}

// Delete deletes a device by its ID
func (s *DeviceStorage) Delete(deviceID string) error {
	return s.db.Where("device_id = ?", deviceID).Delete(&models.Device{}).Error
} 