package core

import (
	"time"

	"github.com/google/uuid"
)

// Identifiable provides ID and resource type for models.
type Identifiable interface {
	GetID() uuid.UUID
	ResourceType() string
}

// Lifecycle hooks for models.
type Lifecycle interface {
	BeforeCreate()
	BeforeUpdate()
}

// GenerateNewID generates a new UUID.
func GenerateNewID() uuid.UUID {
	return uuid.New()
}

// SetAuditFieldsBeforeCreate sets the initial timestamps and createdBy/updatedBy for a model.
// It expects pointers to the model's audit fields.
func SetAuditFieldsBeforeCreate(
	createdAt, updatedAt *time.Time,
	createdBy, updatedBy *uuid.UUID,
) {
	now := time.Now().UTC()
	*createdAt = now
	*updatedAt = now
}

// SetAuditFieldsBeforeUpdate updates the UpdatedAt timestamp and UpdatedBy for a model.
// It expects pointers to the model's audit fields.
func SetAuditFieldsBeforeUpdate(
	updatedAt *time.Time,
	updatedBy *uuid.UUID,
) {
	*updatedAt = time.Now().UTC()
}
