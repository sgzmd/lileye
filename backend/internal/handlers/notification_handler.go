package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lileye/backend/internal/models"
	"github.com/lileye/backend/internal/storage"
)

// NotificationHandler handles HTTP requests for notifications
type NotificationHandler struct {
	storage *storage.NotificationStorage
}

// NewNotificationHandler creates a new NotificationHandler instance
func NewNotificationHandler(storage *storage.NotificationStorage) *NotificationHandler {
	return &NotificationHandler{storage: storage}
}

// RegisterRoutes registers the notification routes with the Gin engine
func (h *NotificationHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/api/notifications", h.CreateNotification)
	r.GET("/api/notifications/:id", h.GetNotification)
	r.GET("/api/notifications/device/:deviceID", h.GetNotificationsByDevice)
	r.GET("/api/notifications/device/:deviceID/range", h.GetNotificationsByDateRange)
	r.GET("/api/notifications/device/:deviceID/search", h.SearchNotifications)
	r.GET("/api/devices", h.GetDevices)
}

// CreateNotification handles the creation of a new notification
func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.storage.Create(&notification); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, notification)
}

// GetNotification handles retrieving a notification by ID
func (h *NotificationHandler) GetNotification(c *gin.Context) {
	id := c.Param("id")
	var idUint uint
	if _, err := fmt.Sscanf(id, "%d", &idUint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	notification, err := h.storage.GetByID(idUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "notification not found"})
		return
	}

	c.JSON(http.StatusOK, notification)
}

// GetNotificationsByDevice handles retrieving notifications for a specific device
func (h *NotificationHandler) GetNotificationsByDevice(c *gin.Context) {
	deviceID := c.Param("deviceID")
	notifications, err := h.storage.GetByDeviceID(deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetNotificationsByDateRange handles retrieving notifications within a date range
func (h *NotificationHandler) GetNotificationsByDateRange(c *gin.Context) {
	deviceID := c.Param("deviceID")
	startStr := c.Query("start")
	endStr := c.Query("end")

	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start date format"})
		return
	}

	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end date format"})
		return
	}

	notifications, err := h.storage.GetByDateRange(deviceID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// SearchNotifications handles searching notifications
func (h *NotificationHandler) SearchNotifications(c *gin.Context) {
	deviceID := c.Param("deviceID")
	query := c.Query("q")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search query is required"})
		return
	}

	notifications, err := h.storage.Search(deviceID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetDevices handles retrieving all unique device IDs
func (h *NotificationHandler) GetDevices(c *gin.Context) {
	devices, err := h.storage.GetDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, devices)
} 