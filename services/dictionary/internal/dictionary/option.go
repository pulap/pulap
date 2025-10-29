package dictionary

import (
	"time"

	"github.com/google/uuid"
	"github.com/pulap/pulap/pkg/lib/core"
)

// Option is the aggregate root for a fake option.
// An option represents a single entry within a set (e.g., "Residential", "House", "Bungalow").
// Options can have hierarchical relationships via ParentID.
type Option struct {
	ID          uuid.UUID  `json:"id"`
	Set         uuid.UUID  `json:"set_id"`                // Reference to the Set this option belongs to
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`   // Optional parent option for hierarchy
	ShortCode   string     `json:"short_code"`            // Short code for the option (e.g., "res", "com")
	Key         string     `json:"key"`                   // Unique key within the set (e.g., "residential", "commercial")
	Label       string     `json:"label"`                 // Human-readable label
	Description string     `json:"description,omitempty"` // Optional description
	Value       string     `json:"value"`                 // The actual value
	Order       int        `json:"order"`                 // Display order
	Active      bool       `json:"active"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   string     `json:"updated_by"`
}

// GetID returns the ID of the Option (implements Identifiable interface).
func (o *Option) GetID() uuid.UUID {
	return o.ID
}

// ResourceType returns the resource type for URL generation.
func (o *Option) ResourceType() string {
	return "fake/option"
}

// SetID sets the ID of the Option.
func (o *Option) SetID(id uuid.UUID) {
	o.ID = id
}

// New creates a new Option with a generated ID.
func NewOption() *Option {
	return &Option{
		ID:     core.GenerateNewID(),
		Active: true,
		Order:  0,
	}
}

// EnsureID ensures the aggregate root has a valid ID.
func (o *Option) EnsureID() {
	if o.ID == uuid.Nil {
		o.ID = core.GenerateNewID()
	}
}

// BeforeCreate sets creation timestamps.
func (o *Option) BeforeCreate() {
	o.EnsureID()
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	if !o.Active {
		o.Active = true
	}
}

// BeforeUpdate sets update timestamps.
func (o *Option) BeforeUpdate() {
	o.UpdatedAt = time.Now()
}
