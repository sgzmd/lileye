package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lileye/backend/internal/storage"
)

type DeviceHandler struct {
	deviceStorage *storage.DeviceStorage
}

func NewDeviceHandler(deviceStorage *storage.DeviceStorage) *DeviceHandler {
	return &DeviceHandler{deviceStorage: deviceStorage}
}

// SetDeviceName handles setting a custom name for a device
func (h *DeviceHandler) SetDeviceName(c *gin.Context) {
	var req struct {
		DeviceID   string `json:"device_id" binding:"required"`
		DeviceName string `json:"device_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.deviceStorage.CreateOrUpdate(req.DeviceID, req.DeviceName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Device name updated successfully"})
}

// GetDeviceName handles getting a custom name for a device
func (h *DeviceHandler) GetDeviceName(c *gin.Context) {
	deviceID := c.Param("device_id")
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_id is required"})
		return
	}

	deviceName, err := h.deviceStorage.GetDeviceName(deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"device_name": deviceName})
}

// GetAllDevices handles getting all devices with their custom names
func (h *DeviceHandler) GetAllDevices(c *gin.Context) {
	devices, err := h.deviceStorage.GetAllDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, devices)
}

// DeleteDevice handles removing a device name mapping
func (h *DeviceHandler) DeleteDevice(c *gin.Context) {
	deviceID := c.Param("device_id")
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_id is required"})
		return
	}

	if err := h.deviceStorage.DeleteDevice(deviceID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Device deleted successfully"})
} 