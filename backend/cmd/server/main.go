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

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.Notification{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize storage and handlers
	notificationStorage := storage.NewNotificationStorage(db)
	notificationHandler := handlers.NewNotificationHandler(notificationStorage)

	// Initialize Gin router
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("./web/templates/*")

	// Register API routes
	notificationHandler.RegisterRoutes(r)

	// Serve index page
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 