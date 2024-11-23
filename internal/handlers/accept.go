package handlers

import (
	"net/http"
	"sync"
	"time"

	"verve/internal/kinesis"
	"verve/internal/logger"
	"verve/internal/redisclient"
	"verve/internal/utils"

	"github.com/gin-gonic/gin"
)

var (
	requestCount int
	mu           sync.Mutex
)

func init() {
	// Periodic logging every minute
	go func() {
		for range time.Tick(1 * time.Minute) {
			processUniqueRequests(requestCount)
		}
	}()
}

// AcceptHandler handles incoming requests.
func AcceptHandler(c *gin.Context) {
	id := c.Query("id")
	endpoint := c.Query("endpoint")

	if id == "" {
		logger.ConsoleLog.Warn("ID is required")
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "id is required"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if redisclient.IsDuplicate(id) {
		logger.ConsoleLog.Infof("Duplicate request received with ID: %s", id)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	requestCount++
	logger.FileLog.Infof("Unique request received with ID: %s", id)

	if endpoint != "" {
		go utils.SendHTTPGet(endpoint, requestCount)
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func processUniqueRequests(count int) {

	// Send unique request count to Kinesis
	if err := kinesis.SendRecord(count); err != nil {
		logger.ConsoleLog.Printf("Error sending record to Kinesis: %v", err)
	}
}
