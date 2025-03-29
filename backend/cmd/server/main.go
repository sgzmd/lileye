package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lileye/backend/internal/handlers"
	"github.com/lileye/backend/internal/models"
	"github.com/lileye/backend/internal/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("notifications.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&models.Notification{}, &models.Device{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize storage
	notificationStorage := storage.NewNotificationStorage(db)
	deviceStorage := storage.NewDeviceStorage(db)

	// Initialize handlers
	notificationHandler := handlers.NewNotificationHandler(notificationStorage)
	deviceHandler := handlers.NewDeviceHandler(deviceStorage)

	// Create Gin router
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")

	// Register notification routes
	notificationHandler.RegisterRoutes(r)

	// Device name routes
	api := r.Group("/api")
	{
		api.POST("/devices", deviceHandler.SetDeviceName)
		api.GET("/devices", deviceHandler.GetAllDevices)
		api.GET("/devices/:device_id", deviceHandler.GetDeviceName)
		api.DELETE("/devices/:device_id", deviceHandler.DeleteDevice)
	}

	// Web interface route
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 