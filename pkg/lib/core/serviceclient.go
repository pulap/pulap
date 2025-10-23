package core

import (
	"context"
	"fmt"
)

type ServiceClient struct {
	baseURL string
	http    *HTTPClient
}

func NewServiceClient(baseURL string) *ServiceClient {
	httpClient := NewHTTPClient(HTTPClientConfig{
		BaseURL: baseURL,
	})

	return &ServiceClient{
		baseURL: baseURL,
		http:    httpClient,
	}
}

func (c *ServiceClient) List(ctx context.Context, resource string) (*SuccessResponse, error) {
	var resp SuccessResponse
	path := fmt.Sprintf("/%s", resource)
	if err := c.http.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *ServiceClient) Get(ctx context.Context, resource, id string) (*SuccessResponse, error) {
	var resp SuccessResponse
	path := fmt.Sprintf("/%s/%s", resource, id)
	if err := c.http.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *ServiceClient) Create(ctx context.Context, resource string, body interface{}) (*SuccessResponse, error) {
	var resp SuccessResponse
	path := fmt.Sprintf("/%s", resource)
	if err := c.http.Post(ctx, path, body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *ServiceClient) Update(ctx context.Context, resource, id string, body interface{}) (*SuccessResponse, error) {
	var resp SuccessResponse
	path := fmt.Sprintf("/%s/%s", resource, id)
	if err := c.http.Patch(ctx, path, body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *ServiceClient) Delete(ctx context.Context, resource, id string) error {
	path := fmt.Sprintf("/%s/%s", resource, id)
	if err := c.http.Delete(ctx, path); err != nil {
		return err
	}

	return nil
}

func (c *ServiceClient) Request(ctx context.Context, method, path string, body interface{}) (*SuccessResponse, error) {
	var resp SuccessResponse

	switch method {
	case "GET":
		if err := c.http.Get(ctx, path, &resp); err != nil {
			return nil, err
		}
	case "POST":
		if err := c.http.Post(ctx, path, body, &resp); err != nil {
			return nil, err
		}
	case "PATCH":
		if err := c.http.Patch(ctx, path, body, &resp); err != nil {
			return nil, err
		}
	case "DELETE":
		if err := c.http.Delete(ctx, path); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	return &resp, nil
}

func (c *ServiceClient) Ping(ctx context.Context) error {
	return c.http.Ping(ctx)
}
