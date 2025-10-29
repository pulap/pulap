package estate

import (
	"context"

	"github.com/google/uuid"
)

// Repo defines the interface for Property aggregate operations.
// This repository manages the Property aggregate root as a single unit.
type Repo interface {
	// Create creates a new Property aggregate.
	Create(ctx context.Context, property *Property) error

	// Get retrieves a complete Property aggregate by ID.
	Get(ctx context.Context, id uuid.UUID) (*Property, error)

	// Save performs a unit-of-work save operation on the aggregate.
	Save(ctx context.Context, property *Property) error

	// Delete removes the entire Property aggregate.
	Delete(ctx context.Context, id uuid.UUID) error

	// List retrieves all Property aggregates.
	List(ctx context.Context) ([]*Property, error)

	// ListByOwner retrieves all properties for a specific owner.
	ListByOwner(ctx context.Context, ownerID string) ([]*Property, error)

	// ListByStatus retrieves all properties with a specific status.
	ListByStatus(ctx context.Context, status string) ([]*Property, error)
}
