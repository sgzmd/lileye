package storage

import (
	"github.com/lileye/backend/internal/models"

	"gorm.io/gorm"
)

type DeviceStorage struct {
	db *gorm.DB
}

func NewDeviceStorage(db *gorm.DB) *DeviceStorage {
	return &DeviceStorage{db: db}
}

// CreateOrUpdate creates a new device name or updates an existing one
func (s *DeviceStorage) CreateOrUpdate(deviceID, deviceName string) error {
	return s.db.Where(models.Device{DeviceID: deviceID}).
		Assign(models.Device{DeviceName: deviceName}).
		FirstOrCreate(&models.Device{DeviceID: deviceID, DeviceName: deviceName}).Error
}

// GetDeviceName returns the custom name for a device, or the device ID if no name is set
func (s *DeviceStorage) GetDeviceName(deviceID string) (string, error) {
	var device models.Device
	err := s.db.Where("device_id = ?", deviceID).First(&device).Error
	if err == nil {
		return device.DeviceName, nil
	}
	if err == gorm.ErrRecordNotFound {
		return deviceID, nil
	}
	return "", err
}

// GetAllDevices returns all devices with their custom names
func (s *DeviceStorage) GetAllDevices() ([]models.Device, error) {
	var devices []models.Device
	err := s.db.Find(&devices).Error
	return devices, err
}

// DeleteDevice removes a device name mapping
func (s *DeviceStorage) DeleteDevice(deviceID string) error {
	return s.db.Where("device_id = ?", deviceID).Delete(&models.Device{}).Error
} 