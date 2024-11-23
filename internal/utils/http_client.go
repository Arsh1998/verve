package utils

import (
	"bytes"
	"encoding/json"
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

// PostRequestBody represents the structure of the POST request body
type PostRequestBody struct {
	UniqueRequestCount int    `json:"unique_request_count"`
	Message            string `json:"message"`
	Timestamp          int64  `json:"timestamp"`
}

// SendHTTPPost sends an HTTP POST request to the provided endpoint with a JSON body.
func SendHTTPPost(endpoint string, count int) {
	payload := PostRequestBody{
		UniqueRequestCount: count,
		Message:            "Request count data",
		Timestamp:          time.Now().Unix(),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		logger.ConsoleLog.Errorf("Failed to serialize POST payload: %v", err)
		return
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		logger.FileLog.Errorf("Failed to create HTTP POST request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logger.FileLog.Errorf("HTTP POST request failed: %v", err)
		return
	}
	defer resp.Body.Close()

	logger.FileLog.Infof("HTTP POST to %s returned status code %d", endpoint, resp.StatusCode)
}
