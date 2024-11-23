package main

import (
	"verve/config"
	"verve/internal/handlers"
	"verve/internal/logger"
	"verve/internal/redisclient"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize logger
	logger.Initialize()

	// Load the configuration
	err := config.LoadConfig("config/config.yaml")
	if err != nil {
		logger.ConsoleLog.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize redis client
	redisclient.Initialize(config.AppConfig)

	// Create Gin router
	router := gin.Default()

	// Register endpoint
	router.GET("/api/verve/accept", handlers.AcceptHandler)

	// Start the server
	router.Run(":8080")
}
