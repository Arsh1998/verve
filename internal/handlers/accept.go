package handlers

import (
	"net/http"
	"sync"
	"time"

	"verve/internal/logger"
	"verve/internal/utils"

	"github.com/gin-gonic/gin"
)

var (
	uniqueRequests sync.Map
	requestCount   int
	mu             sync.Mutex
)

func init() {
	// Periodic logging every minute
	go func() {
		for range time.Tick(1 * time.Minute) {
			logUniqueRequests()
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

	_, exists := uniqueRequests.LoadOrStore(id, true)
	requestCount++

	// Log the uniqueness
	if !exists {
		logger.FileLog.Infof("Unique request received with ID: %s", id)
	} else {
		logger.ConsoleLog.Infof("Duplicate request received with ID: %s", id)
	}

	if endpoint != "" {
		go utils.SendHTTPGet(endpoint, requestCount)
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func logUniqueRequests() {
	count := 0
	uniqueRequests.Range(func(_, _ interface{}) bool {
		count++
		return true
	})

	logger.FileLog.Infof("Unique requests in the last minute: %d", count)

	// Clear the map for the next minute
	uniqueRequests = sync.Map{}
}
