package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// APIResponse represents the structure of the API response
type APIResponse struct {
	// Add fields based on the actual API response structure
	// Example:
	// UserID   string `json:"user_id"`
	// Username string `json:"username"`
	// Role     string `json:"role"`
	Message string `json:"message"`
}

// CallWhoAmIAPI calls the whoami API and returns the response
func CallWhoAmIAPI() (*APIResponse, error) {
	apiURL := "https://10.11.0.190/ha/api/v1/whoami"

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Add headers if needed (e.g., Authorization header)
	// req.Header.Add("Authorization", "Bearer your-token-here")

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse JSON response
	var apiResponse APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return &apiResponse, nil
}

// Example usage
func main() {
	fmt.Println("Calling whoami API...")

	response, err := CallWhoAmIAPI()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("API Response: %+v\n", response)
}
