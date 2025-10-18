package estate

import (
	"time"

	"github.com/google/uuid"
	"github.com/pulap/pulap/pkg/lib/core"
)

// Estate is the aggregate root for the Estate domain.
type Estate struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	Items       []Item    `json:"items"`
	Tags        []Tag     `json:"tags"`
}

// GetID returns the ID of the Estate (implements Identifiable interface).
func (a *Estate) GetID() uuid.UUID {
	return a.ID
}

// ResourceType returns the resource type for URL generation.
func (a *Estate) ResourceType() string {
	return "estate"
}

// SetID sets the ID of the Estate.
func (a *Estate) SetID(id uuid.UUID) {
	a.ID = id
}

// NewEstate creates a new Estate with a generated ID and initial version.
func NewEstate() *Estate {
	return &Estate{
		ID: core.GenerateNewID(),
	}
}

// EnsureID ensures the aggregate root has a valid ID.
func (a *Estate) EnsureID() {
	if a.ID == uuid.Nil {
		a.ID = core.GenerateNewID()
	}
}

// BeforeCreate sets creation timestamps and version.
func (a *Estate) BeforeCreate() {
	a.EnsureID()
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

// BeforeUpdate sets update timestamps and increments version.
func (a *Estate) BeforeUpdate() {
	a.UpdatedAt = time.Now()
}
