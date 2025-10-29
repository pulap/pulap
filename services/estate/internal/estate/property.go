package estate

import (
	"time"

	"github.com/google/uuid"
	"github.com/pulap/pulap/pkg/lib/core"
)

// Property is the aggregate root for the real estate domain.
// It represents a real estate property (house, apartment, commercial space, land, etc.)
// with its classification, location, physical features, and pricing information.
type Property struct {
	ID             uuid.UUID      `json:"id"`
	Name           string         `json:"name"`               // Short name/title for the property
	Description    string         `json:"description"`        // Detailed description
	Classification Classification `json:"classification"`     // Category, Type, Subtype (fake refs)
	Location       Location       `json:"location"`           // Address and coordinates
	Features       Features       `json:"features"`           // Physical characteristics
	Price          Price          `json:"price"`              // Pricing information
	Status         string         `json:"status"`             // e.g., "available", "sold", "rented", "reserved"
	OwnerID        string         `json:"owner_id,omitempty"` // Reference to owner/user
	CreatedAt      time.Time      `json:"created_at"`
	CreatedBy      string         `json:"created_by"`
	UpdatedAt      time.Time      `json:"updated_at"`
	UpdatedBy      string         `json:"updated_by"`
}

// Price represents pricing information for a property.
type Price struct {
	Amount     float64 `json:"amount"`     // Price amount
	Currency   string  `json:"currency"`   // e.g., "USD", "EUR", "ARS"
	Type       string  `json:"type"`       // e.g., "sale", "rent_monthly", "rent_daily"
	Negotiable bool    `json:"negotiable"` // Whether price is negotiable
}

// Validate performs basic validation on the price.
func (p Price) Validate() []string {
	var errors []string

	if p.Amount < 0 {
		errors = append(errors, "price.amount cannot be negative")
	}

	if p.Currency == "" {
		errors = append(errors, "price.currency is required")
	}

	if p.Type == "" {
		errors = append(errors, "price.type is required")
	}

	// Validate price type
	validPriceTypes := map[string]bool{
		"sale":         true,
		"rent_monthly": true,
		"rent_daily":   true,
		"rent_weekly":  true,
		"rent_yearly":  true,
	}
	if !validPriceTypes[p.Type] {
		errors = append(errors, "price.type must be one of: sale, rent_monthly, rent_daily, rent_weekly, rent_yearly")
	}

	return errors
}

// GetID returns the ID of the Property (implements Identifiable interface).
func (p *Property) GetID() uuid.UUID {
	return p.ID
}

// ResourceType returns the resource type for URL generation.
func (p *Property) ResourceType() string {
	return "estate"
}

// SetID sets the ID of the Property.
func (p *Property) SetID(id uuid.UUID) {
	p.ID = id
}

// New creates a new Property with a generated ID.
func New() *Property {
	return &Property{
		ID:     core.GenerateNewID(),
		Status: "available",
		Price: Price{
			Currency: "USD",
			Type:     "sale",
		},
	}
}

// EnsureID ensures the aggregate root has a valid ID.
func (p *Property) EnsureID() {
	if p.ID == uuid.Nil {
		p.ID = core.GenerateNewID()
	}
}

// BeforeCreate sets creation timestamps.
func (p *Property) BeforeCreate() {
	p.EnsureID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	if p.Status == "" {
		p.Status = "available"
	}
}

// BeforeUpdate sets update timestamps.
func (p *Property) BeforeUpdate() {
	p.UpdatedAt = time.Now()
}
