package dictionary

import (
	"time"

	"github.com/google/uuid"
	"github.com/pulap/pulap/pkg/lib/core"
)

// Set is the aggregate root for a fake set.
// A set is a container for related options (e.g., "estate_category", "estate_type").
type Set struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`  // Unique name, e.g., "estate_category", "estate_type"
	Label       string    `json:"label"` // Human-readable label
	Description string    `json:"description,omitempty"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
}

// GetID returns the ID of the Set (implements Identifiable interface).
func (s *Set) GetID() uuid.UUID {
	return s.ID
}

// ResourceType returns the resource type for URL generation.
func (s *Set) ResourceType() string {
	return "fake/set"
}

// SetID sets the ID of the Set.
func (s *Set) SetID(id uuid.UUID) {
	s.ID = id
}

// New creates a new Set with a generated ID.
func NewSet() *Set {
	return &Set{
		ID:     core.GenerateNewID(),
		Active: true,
	}
}

// EnsureID ensures the aggregate root has a valid ID.
func (s *Set) EnsureID() {
	if s.ID == uuid.Nil {
		s.ID = core.GenerateNewID()
	}
}

// BeforeCreate sets creation timestamps.
func (s *Set) BeforeCreate() {
	s.EnsureID()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	if !s.Active {
		s.Active = true
	}
}

// BeforeUpdate sets update timestamps.
func (s *Set) BeforeUpdate() {
	s.UpdatedAt = time.Now()
}
