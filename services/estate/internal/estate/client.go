package estate

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Client is the interface for interacting with the Dictionary service.
// It provides methods to retrieve and validate dictionary options.
type Client interface {
	// GetOption retrieves a single option by ID.
	GetOption(ctx context.Context, id uuid.UUID) (*Option, error)

	// GetOptionsByParent lists all options in a set filtered by parent ID.
	// If parentID is nil, returns root-level options (no parent).
	ListOptionsByParent(ctx context.Context, setName string, parentID *uuid.UUID) ([]Option, error)

	// ValidateClassification validates that a classification is valid:
	// - All IDs exist and are active
	// - Type.parent_id equals CategoryID
	// - Subtype.parent_id equals TypeID (if subtype is provided)
	// Returns true if valid, along with validation error messages if any.
	ValidateClassification(ctx context.Context, c Classification) (bool, []string, error)
}

// Option represents a dictionary option (category, type, or subtype).
// This is a data type for the Dictionary service's Option entity.
type Option struct {
	ID          uuid.UUID  `json:"id"`
	SetID       uuid.UUID  `json:"set_id"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
	ShortCode   string     `json:"short_code"`
	Key         string     `json:"key"`
	Label       string     `json:"label"`
	Description string     `json:"description,omitempty"`
	Value       string     `json:"value"`
	Order       int        `json:"order"`
	Active      bool       `json:"active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Set represents a dictionary set (container for options).
// This is a DTO for the Dictionary service's Set entity.
type Set struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`  // e.g., "estate_category", "estate_type"
	Label       string    `json:"label"` // Human-readable label
	Description string    `json:"description,omitempty"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
