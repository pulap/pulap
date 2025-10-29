package estate

import (
	"context"

	"github.com/google/uuid"
)

// ValidationError represents a validation error.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateCreateProperty validates a Property entity before creation.
func ValidateCreateProperty(ctx context.Context, property *Property) []ValidationError {
	var errors []ValidationError

	// ID should be set
	if property.ID == uuid.Nil {
		errors = append(errors, ValidationError{
			Field:   "id",
			Message: "ID is required",
		})
	}

	// Name is required
	if property.Name == "" {
		errors = append(errors, ValidationError{
			Field:   "name",
			Message: "Name is required",
		})
	}

	// Validate classification (basic, full validation happens via DictionaryClient)
	if classErrors := property.Classification.Validate(); len(classErrors) > 0 {
		for _, err := range classErrors {
			errors = append(errors, ValidationError{
				Field:   "classification",
				Message: err,
			})
		}
	}

	// Validate location
	if locErrors := property.Location.Validate(); len(locErrors) > 0 {
		for _, err := range locErrors {
			errors = append(errors, ValidationError{
				Field:   "location",
				Message: err,
			})
		}
	}

	// Validate features
	if featErrors := property.Features.Validate(); len(featErrors) > 0 {
		for _, err := range featErrors {
			errors = append(errors, ValidationError{
				Field:   "features",
				Message: err,
			})
		}
	}

	// Validate price
	if priceErrors := property.Price.Validate(); len(priceErrors) > 0 {
		for _, err := range priceErrors {
			errors = append(errors, ValidationError{
				Field:   "price",
				Message: err,
			})
		}
	}

	// Status should be valid
	if property.Status != "" {
		validStatuses := map[string]bool{
			"available": true,
			"sold":      true,
			"rented":    true,
			"reserved":  true,
			"draft":     true,
			"inactive":  true,
		}
		if !validStatuses[property.Status] {
			errors = append(errors, ValidationError{
				Field:   "status",
				Message: "Status must be one of: available, sold, rented, reserved, draft, inactive",
			})
		}
	}

	// CreatedAt should be set
	if property.CreatedAt.IsZero() {
		errors = append(errors, ValidationError{
			Field:   "created_at",
			Message: "Created timestamp is required",
		})
	}

	// UpdatedAt should be set
	if property.UpdatedAt.IsZero() {
		errors = append(errors, ValidationError{
			Field:   "updated_at",
			Message: "Updated timestamp is required",
		})
	}

	return errors
}

// ValidateUpdateProperty validates a Property entity before update.
func ValidateUpdateProperty(ctx context.Context, id uuid.UUID, property *Property) []ValidationError {
	var errors []ValidationError

	// ID should match parameter
	if property.ID != id {
		errors = append(errors, ValidationError{
			Field:   "id",
			Message: "ID mismatch",
		})
	}

	// Name is required
	if property.Name == "" {
		errors = append(errors, ValidationError{
			Field:   "name",
			Message: "Name is required",
		})
	}

	// Validate classification
	if classErrors := property.Classification.Validate(); len(classErrors) > 0 {
		for _, err := range classErrors {
			errors = append(errors, ValidationError{
				Field:   "classification",
				Message: err,
			})
		}
	}

	// Validate location
	if locErrors := property.Location.Validate(); len(locErrors) > 0 {
		for _, err := range locErrors {
			errors = append(errors, ValidationError{
				Field:   "location",
				Message: err,
			})
		}
	}

	// Validate features
	if featErrors := property.Features.Validate(); len(featErrors) > 0 {
		for _, err := range featErrors {
			errors = append(errors, ValidationError{
				Field:   "features",
				Message: err,
			})
		}
	}

	// Validate price
	if priceErrors := property.Price.Validate(); len(priceErrors) > 0 {
		for _, err := range priceErrors {
			errors = append(errors, ValidationError{
				Field:   "price",
				Message: err,
			})
		}
	}

	// UpdatedAt should be set
	if property.UpdatedAt.IsZero() {
		errors = append(errors, ValidationError{
			Field:   "updated_at",
			Message: "Updated timestamp is required",
		})
	}

	return errors
}

// ValidateDeleteProperty validates a Property entity before deletion.
func ValidateDeleteProperty(ctx context.Context, id uuid.UUID) []ValidationError {
	var errors []ValidationError

	// ID should be valid
	if id == uuid.Nil {
		errors = append(errors, ValidationError{
			Field:   "id",
			Message: "Valid ID is required for deletion",
		})
	}

	return errors
}
