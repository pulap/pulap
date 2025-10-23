package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	BaseURL    string
	HTTPClient *http.Client
	MaxRetries int
	RetryDelay time.Duration
}

type HTTPClientConfig struct {
	BaseURL    string
	Timeout    time.Duration
	MaxRetries int
	RetryDelay time.Duration
}

func NewHTTPClient(config HTTPClientConfig) *HTTPClient {
	if config.Timeout == 0 {
		config.Timeout = 10 * time.Second
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}
	if config.RetryDelay == 0 {
		config.RetryDelay = 1 * time.Second
	}

	return &HTTPClient{
		BaseURL: config.BaseURL,
		HTTPClient: &http.Client{
			Timeout: config.Timeout,
		},
		MaxRetries: config.MaxRetries,
		RetryDelay: config.RetryDelay,
	}
}

func (c *HTTPClient) Get(ctx context.Context, path string, result interface{}) error {
	return c.doWithRetry(ctx, "GET", path, nil, result)
}

func (c *HTTPClient) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.doWithRetry(ctx, "POST", path, body, result)
}

func (c *HTTPClient) Patch(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.doWithRetry(ctx, "PATCH", path, body, result)
}

func (c *HTTPClient) Delete(ctx context.Context, path string) error {
	return c.doWithRetry(ctx, "DELETE", path, nil, nil)
}

func (c *HTTPClient) doWithRetry(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	var lastErr error

	for attempt := 0; attempt <= c.MaxRetries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(c.RetryDelay * time.Duration(attempt)):
			}
		}

		err := c.do(ctx, method, path, body, result)
		if err == nil {
			return nil
		}

		lastErr = err

		if !c.shouldRetry(err) {
			return err
		}
	}

	return fmt.Errorf("max retries (%d) exceeded: %w", c.MaxRetries, lastErr)
}

func (c *HTTPClient) do(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	url := c.BaseURL + path

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Message:    string(bodyBytes),
		}
	}

	if result != nil && resp.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

func (c *HTTPClient) shouldRetry(err error) bool {
	if err == nil {
		return false
	}

	httpErr, ok := err.(*HTTPError)
	if !ok {
		return true
	}

	switch httpErr.StatusCode {
	case http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusConflict:
		return false
	case http.StatusTooManyRequests,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return true
	default:
		return httpErr.StatusCode >= 500
	}
}

type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

func (e *HTTPError) IsNotFound() bool {
	return e.StatusCode == http.StatusNotFound
}

func (e *HTTPError) IsUnauthorized() bool {
	return e.StatusCode == http.StatusUnauthorized
}

func (e *HTTPError) IsForbidden() bool {
	return e.StatusCode == http.StatusForbidden
}

func (c *HTTPClient) Ping(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+"/healthz", nil)
	if err != nil {
		return fmt.Errorf("failed to create ping request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("service unhealthy: status %d", resp.StatusCode)
	}

	return nil
}
