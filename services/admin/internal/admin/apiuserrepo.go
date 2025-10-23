package admin

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pulap/pulap/pkg/lib/core"
)

// APIUserRepo implements UserRepo using ServiceClient to call authn service
type APIUserRepo struct {
	client *core.ServiceClient
}

// NewAPIUserRepo creates a new API-based user repository
func NewAPIUserRepo(client *core.ServiceClient) *APIUserRepo {
	return &APIUserRepo{
		client: client,
	}
}

// List retrieves all users from authn service
func (r *APIUserRepo) List(ctx context.Context) ([]*User, error) {
	resp, err := r.client.List(ctx, "users")
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	// Parse response data
	usersData, ok := resp.Data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	users := make([]*User, 0, len(usersData))
	for _, item := range usersData {
		userData, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		idStr := stringField(userData, "id")
		if idStr == "" {
			continue
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}

		user := &User{
			ID:     id,
			Email:  stringField(userData, "email"),
			Name:   stringField(userData, "name"),
			Status: stringField(userData, "status"),
		}
		users = append(users, user)
	}

	return users, nil
}

// Get retrieves a user by ID from authn service
func (r *APIUserRepo) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	resp, err := r.client.Get(ctx, "users", id.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	userData, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	idStr := stringField(userData, "id")
	if idStr == "" {
		return nil, fmt.Errorf("missing user id in response")
	}

	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	user := &User{
		ID:     parsedID,
		Email:  stringField(userData, "email"),
		Name:   stringField(userData, "name"),
		Status: stringField(userData, "status"),
	}

	return user, nil
}

// Create creates a new user via authn service
func (r *APIUserRepo) Create(ctx context.Context, req *CreateUserRequest) (*User, error) {
	_, err := r.client.Create(ctx, "users", req)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// TODO: Return created user from response
	return &User{
		Email:  req.Email,
		Name:   req.Name,
		Status: "active",
	}, nil
}

// Update updates an existing user via authn service
func (r *APIUserRepo) Update(ctx context.Context, id uuid.UUID, req *UpdateUserRequest) (*User, error) {
	_, err := r.client.Update(ctx, "users", id.String(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// TODO: Return updated user from response
	return &User{
		ID:     id,
		Email:  req.Email,
		Name:   req.Name,
		Status: req.Status,
	}, nil
}

// Delete removes a user via authn service
func (r *APIUserRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.client.Delete(ctx, "users", id.String()); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// ListByStatus retrieves users filtered by status
func (r *APIUserRepo) ListByStatus(ctx context.Context, status string) ([]*User, error) {
	// For now, fetch all and filter client-side
	// TODO: Add query parameters support to ServiceClient
	allUsers, err := r.List(ctx)
	if err != nil {
		return nil, err
	}

	filtered := make([]*User, 0)
	for _, user := range allUsers {
		if user.Status == status {
			filtered = append(filtered, user)
		}
	}

	return filtered, nil
}

func stringField(data map[string]interface{}, key string) string {
	value, ok := data[key]
	if !ok || value == nil {
		return ""
	}

	if str, ok := value.(string); ok {
		return str
	}

	return fmt.Sprintf("%v", value)
}
