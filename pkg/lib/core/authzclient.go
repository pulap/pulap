package core

import (
	"context"
	"fmt"
)

type AuthzHTTPClient struct {
	http *HTTPClient
}

func NewAuthZHTTPClient(baseURL string) *AuthzHTTPClient {
	httpClient := NewHTTPClient(HTTPClientConfig{
		BaseURL: baseURL,
	})

	return &AuthzHTTPClient{
		http: httpClient,
	}
}

func (c *AuthzHTTPClient) CheckPermission(ctx context.Context, userID, permission, resource string) (bool, error) {
	req := PermissionRequest{
		UserID:     userID,
		Permission: permission,
		Scope: AuthzScope{
			Type: "global",
			ID:   "",
		},
	}

	if resource != "" && resource != "*" {
		req.Scope = AuthzScope{
			Type: "resource",
			ID:   resource,
		}
	}

	var resp SuccessResponse
	err := c.http.Post(ctx, "/authz/policy/evaluate", req, &resp)
	if err != nil {
		return false, fmt.Errorf("authz check failed: %w", err)
	}

	permResp, ok := resp.Data.(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("invalid response format")
	}

	allowed, ok := permResp["allowed"].(bool)
	if !ok {
		return false, fmt.Errorf("invalid allowed field in response")
	}

	return allowed, nil
}

type PermissionRequest struct {
	UserID     string     `json:"user_id"`
	Permission string     `json:"permission"`
	Scope      AuthzScope `json:"scope"`
}

type AuthzScope struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
