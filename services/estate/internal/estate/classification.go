package estate

import "github.com/google/uuid"

// Classification represents the hierarchical taxonomy of a property.
// It stores references (IDs) to fake options, not labels.
// The fake service owns the actual category/type/subtype data.
type Classification struct {
	CategoryID uuid.UUID `json:"category_id"` // e.g., Residential, Commercial, Land
	TypeID     uuid.UUID `json:"type_id"`     // e.g., House, Apartment, Office
	SubtypeID  uuid.UUID `json:"subtype_id"`  // e.g., Bungalow, Loft, Showroom
}

// IsZero returns true if all IDs are nil (not set).
func (c Classification) IsZero() bool {
	return c.CategoryID == uuid.Nil && c.TypeID == uuid.Nil && c.SubtypeID == uuid.Nil
}

// Validate performs basic validation on the classification.
// Full validation against the fake service should be done in the handler layer.
func (c Classification) Validate() []string {
	var errors []string

	if c.CategoryID == uuid.Nil {
		errors = append(errors, "category_id is required")
	}

	if c.TypeID == uuid.Nil {
		errors = append(errors, "type_id is required")
	}

	// Subtype is optional, but if provided must be valid
	// Full hierarchy validation (Type.parent_id = Category, Subtype.parent_id = Type)
	// is done via DictionaryClient in the handler layer

	return errors
}
