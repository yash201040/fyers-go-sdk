package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client holds the configuration for making requests to the Fyers API.
type Client struct {
	AppID       string
	AccessToken string
	BaseURL     string
	HttpClient  *http.Client
}

// NewClient initializes and returns a new API client.
func NewClient(appID, accessToken, baseURL string) *Client {
	return &Client{
		AppID:       appID,
		AccessToken: accessToken,
		BaseURL:     baseURL,
		HttpClient:  &http.Client{},
	}
}

func CheckHTTPResponse(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK HTTP status: %d", resp.StatusCode)
	}

	// Check for empty or invalid body
	if resp.ContentLength == 0 {
		return fmt.Errorf("empty response body")
	}

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return fmt.Errorf("failed to decode response body: %w", err)
	}

	// Rewind the body for future use
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal response body: %w", err)
	}
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Check if the "status" field is present and is "ok"
	if status, ok := body["status"].(string); !ok || status != "ok" {
		return fmt.Errorf("received non-OK status in body: %v", body)
	}

	return nil
}
