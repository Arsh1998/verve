package main

import (
	"verve/internal/handlers"
	"verve/internal/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize logger
	logger.Initialize()

	// Create Gin router
	router := gin.Default()

	// Register endpoint
	router.GET("/api/verve/accept", handlers.AcceptHandler)

	// Start the server
	router.Run(":8080")
}
