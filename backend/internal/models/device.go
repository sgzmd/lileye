package models

import (
	"gorm.io/gorm"
)

// Device represents a device with its custom name
type Device struct {
	gorm.Model
	DeviceID   string `json:"device_id" gorm:"uniqueIndex;not null"`
	CustomName string `json:"custom_name" gorm:"not null"`
}

func (Device) TableName() string {
	return "devices"
} 