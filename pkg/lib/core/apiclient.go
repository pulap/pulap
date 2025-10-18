package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// APIClient provides HTTP client functionality for microservice communication
type APIClient struct {
	baseURL    string
	httpClient *http.Client
	timeout    time.Duration
}

// NewAPIClient creates a new API client with default timeout
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		timeout: 30 * time.Second,
	}
}

// SetTimeout configures request timeout
func (c *APIClient) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
	c.httpClient.Timeout = timeout
}

// Get performs GET request and unmarshals response
func (c *APIClient) Get(ctx context.Context, path string, result interface{}) error {
	url := c.baseURL + path
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %d %s", resp.StatusCode, resp.Status)
	}
	
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("error decoding response: %w", err)
		}
	}
	
	return nil
}

// Post performs POST request with JSON payload
func (c *APIClient) Post(ctx context.Context, path string, payload interface{}, result interface{}) error {
	return c.doJSONRequest(ctx, http.MethodPost, path, payload, result)
}

// Put performs PUT request with JSON payload
func (c *APIClient) Put(ctx context.Context, path string, payload interface{}, result interface{}) error {
	return c.doJSONRequest(ctx, http.MethodPut, path, payload, result)
}

// Delete performs DELETE request
func (c *APIClient) Delete(ctx context.Context, path string) error {
	url := c.baseURL + path
	
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %d %s", resp.StatusCode, resp.Status)
	}
	
	return nil
}

// doJSONRequest handles POST/PUT requests with JSON payload
func (c *APIClient) doJSONRequest(ctx context.Context, method, path string, payload interface{}, result interface{}) error {
	url := c.baseURL + path
	
	var body *bytes.Buffer
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("error encoding payload: %w", err)
		}
		body = bytes.NewBuffer(jsonData)
	} else {
		body = bytes.NewBuffer(nil)
	}
	
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %d %s", resp.StatusCode, resp.Status)
	}
	
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("error decoding response: %w", err)
		}
	}
	
	return nil
}

// CheckHealth performs health check on the service
func (c *APIClient) CheckHealth(ctx context.Context) error {
	return c.Get(ctx, "/health", nil)
}