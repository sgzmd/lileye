package models

import (
	"time"

	"gorm.io/gorm"
)

// Notification represents an Android notification in the system
type Notification struct {
	gorm.Model
	Title       string    `json:"title" gorm:"not null"`
	Message     string    `json:"message" gorm:"not null"`
	Timestamp   time.Time `json:"timestamp" gorm:"not null;index"`
	PackageName string    `json:"package_name" gorm:"not null;index"`
	From        string    `json:"from" gorm:"index"`
	DeviceID    string    `json:"device_id" gorm:"not null;index"`
} 