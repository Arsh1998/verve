package utils

import (
	"fmt"
	"net/http"
	"time"
	"verve/internal/logger"
)

// SendHTTPGet sends an HTTP GET request with query parameters.
func SendHTTPGet(endpoint string, count int) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		logger.FileLog.Errorf("Failed to create HTTP request: %v", err)
		return
	}

	// Add query parameter
	q := req.URL.Query()
	q.Add("count", fmt.Sprintf("%d", count))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		logger.FileLog.Errorf("HTTP request failed: %v", err)
		return
	}
	defer resp.Body.Close()

	logger.FileLog.Infof("HTTP GET to %s returned status code %d", endpoint, resp.StatusCode)
}
