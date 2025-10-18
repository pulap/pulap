package estate

import (
	"context"

	"github.com/google/uuid"
)

// EstateRepo defines the interface for Estate aggregate operations.
// This repository manages the aggregate root and all its child entities as a single unit.
type EstateRepo interface {
	// Create creates a new Estate aggregate with all its child entities.
	Create(ctx context.Context, aggregate *Estate) error

	// Get retrieves a complete Estate aggregate by ID, including all child entities.
	Get(ctx context.Context, id uuid.UUID) (*Estate, error)

	// Save performs a unit-of-work save operation on the aggregate.
	// This will compute differences and update/insert/delete child entities as needed.
	Save(ctx context.Context, aggregate *Estate) error

	// Delete removes the entire Estate aggregate and all its child entities.
	Delete(ctx context.Context, id uuid.UUID) error

	// Estate retrieves all Estate aggregates with their child entities.
	Estate(ctx context.Context) ([]*Estate, error)
}
