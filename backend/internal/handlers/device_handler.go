package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lileye/backend/internal/models"
	"github.com/lileye/backend/internal/storage"
)

// DeviceHandler handles HTTP requests for devices
type DeviceHandler struct {
	storage *storage.DeviceStorage
}

// NewDeviceHandler creates a new DeviceHandler instance
func NewDeviceHandler(storage *storage.DeviceStorage) *DeviceHandler {
	return &DeviceHandler{storage: storage}
}

// RegisterRoutes registers the device routes with the Gin engine
func (h *DeviceHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/api/devices/names", h.GetAllDevices)
	r.POST("/api/devices/names", h.CreateOrUpdateDevice)
	r.DELETE("/api/devices/names/:deviceID", h.DeleteDevice)
}

// GetAllDevices handles retrieving all devices with their custom names
func (h *DeviceHandler) GetAllDevices(c *gin.Context) {
	devices, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, devices)
}

// CreateOrUpdateDevice handles creating or updating a device's custom name
func (h *DeviceHandler) CreateOrUpdateDevice(c *gin.Context) {
	var device models.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if device.DeviceID == "" || device.CustomName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_id and custom_name are required"})
		return
	}

	if err := h.storage.CreateOrUpdate(&device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, device)
}

// DeleteDevice handles deleting a device's custom name
func (h *DeviceHandler) DeleteDevice(c *gin.Context) {
	deviceID := c.Param("deviceID")
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_id is required"})
		return
	}

	if err := h.storage.Delete(deviceID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Device deleted successfully"})
} 