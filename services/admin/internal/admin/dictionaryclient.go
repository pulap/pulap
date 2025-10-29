package admin

import (
	"context"

	"github.com/google/uuid"
	"github.com/pulap/pulap/pkg/lib/core"
)

// DictionaryClient provides access to dictionary service for taxonomy data.
type DictionaryClient interface {
	ListCategories(ctx context.Context) ([]DictionaryOption, error)
	ListTypesByCategory(ctx context.Context, categoryID uuid.UUID) ([]DictionaryOption, error)
	ListSubtypesByType(ctx context.Context, typeID uuid.UUID) ([]DictionaryOption, error)
}

// DictionaryOption represents a dictionary option.
type DictionaryOption struct {
	ID   uuid.UUID
	Name string
}

// FakeDictionaryClient provides hardcoded dictionary data for development.
type FakeDictionaryClient struct{}

// NewFakeDictionaryClient creates a new fake dictionary client.
func NewFakeDictionaryClient() *FakeDictionaryClient {
	return &FakeDictionaryClient{}
}

// ListCategories returns all categories.
func (c *FakeDictionaryClient) ListCategories(ctx context.Context) ([]DictionaryOption, error) {
	return []DictionaryOption{
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"), Name: "Residential"},
		{ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Name: "Commercial"},
		{ID: uuid.MustParse("33333333-3333-3333-3333-333333333333"), Name: "Industrial"},
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444444"), Name: "Land"},
		{ID: uuid.MustParse("55555555-5555-5555-5555-555555555555"), Name: "Special Purpose"},
		{ID: uuid.MustParse("66666666-6666-6666-6666-666666666666"), Name: "Mixed Use"},
	}, nil
}

// ListTypesByCategory returns types for a given category.
func (c *FakeDictionaryClient) ListTypesByCategory(ctx context.Context, categoryID uuid.UUID) ([]DictionaryOption, error) {
	// Simplified: return all types for now
	// In real implementation, filter by category_id
	return []DictionaryOption{
		{ID: uuid.MustParse("a1111111-1111-1111-1111-111111111111"), Name: "House"},
		{ID: uuid.MustParse("a2222222-2222-2222-2222-222222222222"), Name: "Apartment"},
		{ID: uuid.MustParse("a3333333-3333-3333-3333-333333333333"), Name: "Office"},
		{ID: uuid.MustParse("a4444444-4444-4444-4444-444444444444"), Name: "Retail"},
	}, nil
}

// ListSubtypesByType returns subtypes for a given type.
func (c *FakeDictionaryClient) ListSubtypesByType(ctx context.Context, typeID uuid.UUID) ([]DictionaryOption, error) {
	// Simplified: return all subtypes for now
	// In real implementation, filter by type_id
	return []DictionaryOption{
		{ID: uuid.MustParse("b1111111-1111-1111-1111-111111111111"), Name: "Bungalow"},
		{ID: uuid.MustParse("b2222222-2222-2222-2222-222222222222"), Name: "Studio"},
		{ID: uuid.MustParse("b3333333-3333-3333-3333-333333333333"), Name: "Loft"},
	}, nil
}

// APIDictionaryClient calls the real dictionary service.
type APIDictionaryClient struct {
	client *core.ServiceClient
}

// NewAPIDictionaryClient creates a new API-based dictionary client.
func NewAPIDictionaryClient(client *core.ServiceClient) *APIDictionaryClient {
	return &APIDictionaryClient{
		client: client,
	}
}

// ListCategories returns all categories from dictionary service.
func (c *APIDictionaryClient) ListCategories(ctx context.Context) ([]DictionaryOption, error) {
	// TODO: Implement when dictionary service is ready
	// For now, delegate to fake
	fake := NewFakeDictionaryClient()
	return fake.ListCategories(ctx)
}

// ListTypesByCategory returns types for a category from dictionary service.
func (c *APIDictionaryClient) ListTypesByCategory(ctx context.Context, categoryID uuid.UUID) ([]DictionaryOption, error) {
	// TODO: Implement when dictionary service is ready
	fake := NewFakeDictionaryClient()
	return fake.ListTypesByCategory(ctx, categoryID)
}

// ListSubtypesByType returns subtypes for a type from dictionary service.
func (c *APIDictionaryClient) ListSubtypesByType(ctx context.Context, typeID uuid.UUID) ([]DictionaryOption, error) {
	// TODO: Implement when dictionary service is ready
	fake := NewFakeDictionaryClient()
	return fake.ListSubtypesByType(ctx, typeID)
}

// Helper to convert options to map for templates
func DictionaryOptionsToMap(options []DictionaryOption) []map[string]string {
	result := make([]map[string]string, len(options))
	for i, opt := range options {
		result[i] = map[string]string{
			"id":   opt.ID.String(),
			"name": opt.Name,
		}
	}
	return result
}
