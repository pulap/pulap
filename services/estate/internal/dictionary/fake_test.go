package dictionary

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/pulap/pulap/services/estate/internal/estate"
)

func TestFakeGetOption(t *testing.T) {
	fake := NewFake()
	ctx := context.Background()

	// Test getting existing option
	residentialID := uuid.MustParse("00000000-0000-0000-0001-000000000001")
	opt, err := fake.GetOption(ctx, residentialID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if opt.Key != "residential" {
		t.Errorf("expected key 'residential', got '%s'", opt.Key)
	}
	if opt.Label != "Residential" {
		t.Errorf("expected label 'Residential', got '%s'", opt.Label)
	}
	if opt.ParentID != nil {
		t.Errorf("expected nil parent for category, got %v", opt.ParentID)
	}

	// Test getting non-existent option
	nonExistent := uuid.New()
	_, err = fake.GetOption(ctx, nonExistent)
	if err != ErrOptionNotFound {
		t.Errorf("expected ErrOptionNotFound, got %v", err)
	}
}

func TestFakeListOptionsByParent(t *testing.T) {
	fake := NewFake()
	ctx := context.Background()

	// Test listing root-level categories (no parent)
	categories, err := fake.ListOptionsByParent(ctx, "estate_category", nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(categories) != 6 {
		t.Errorf("expected 6 categories, got %d", len(categories))
	}

	// Test listing types under residential category
	residentialID := uuid.MustParse("00000000-0000-0000-0001-000000000001")
	types, err := fake.ListOptionsByParent(ctx, "estate_type", &residentialID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(types) != 2 {
		t.Errorf("expected 2 residential types (house, apartment), got %d", len(types))
	}

	// Verify the types are house and apartment
	keys := make(map[string]bool)
	for _, typ := range types {
		keys[typ.Key] = true
		if typ.ParentID == nil || *typ.ParentID != residentialID {
			t.Errorf("expected parent_id to be residential category, got %v", typ.ParentID)
		}
	}
	if !keys["house"] || !keys["apartment"] {
		t.Errorf("expected 'house' and 'apartment', got keys: %v", keys)
	}

	// Test listing subtypes under house type
	houseID := uuid.MustParse("00000000-0000-0000-0002-000000000001")
	subtypes, err := fake.ListOptionsByParent(ctx, "estate_subtype", &houseID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(subtypes) != 1 {
		t.Errorf("expected 1 subtype under house (bungalow), got %d", len(subtypes))
	}
	if len(subtypes) > 0 && subtypes[0].Key != "bungalow" {
		t.Errorf("expected subtype 'bungalow', got '%s'", subtypes[0].Key)
	}
}

func TestFakeValidateClassificationValid(t *testing.T) {
	fake := NewFake()
	ctx := context.Background()

	// Valid classification: Residential -> House -> Bungalow
	classification := estate.Classification{
		CategoryID: uuid.MustParse("00000000-0000-0000-0001-000000000001"), // residential
		TypeID:     uuid.MustParse("00000000-0000-0000-0002-000000000001"), // house
		SubtypeID:  uuid.MustParse("00000000-0000-0000-0003-000000000001"), // bungalow
	}

	valid, errs, err := fake.ValidateClassification(ctx, classification)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !valid {
		t.Errorf("expected valid classification, got errors: %v", errs)
	}
	if len(errs) != 0 {
		t.Errorf("expected no validation errors, got: %v", errs)
	}
}

func TestFakeValidateClassificationValidWithoutSubtype(t *testing.T) {
	fake := NewFake()
	ctx := context.Background()

	// Valid classification without subtype: Commercial -> Office
	classification := estate.Classification{
		CategoryID: uuid.MustParse("00000000-0000-0000-0001-000000000002"), // commercial
		TypeID:     uuid.MustParse("00000000-0000-0000-0002-000000000003"), // office
		SubtypeID:  uuid.Nil,                                               // no subtype
	}

	valid, errs, err := fake.ValidateClassification(ctx, classification)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !valid {
		t.Errorf("expected valid classification, got errors: %v", errs)
	}
	if len(errs) != 0 {
		t.Errorf("expected no validation errors, got: %v", errs)
	}
}

func TestFakeValidateClassificationInvalidHierarchy(t *testing.T) {
	fake := NewFake()
	ctx := context.Background()

	// Invalid: Type doesn't belong to Category (house under commercial)
	classification := estate.Classification{
		CategoryID: uuid.MustParse("00000000-0000-0000-0001-000000000002"), // commercial
		TypeID:     uuid.MustParse("00000000-0000-0000-0002-000000000001"), // house (belongs to residential)
		SubtypeID:  uuid.Nil,
	}

	valid, errs, err := fake.ValidateClassification(ctx, classification)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if valid {
		t.Error("expected invalid classification")
	}
	if len(errs) == 0 {
		t.Error("expected validation errors")
	}

	// Check for specific error
	found := false
	for _, e := range errs {
		if e == "type does not belong to the selected category" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'type does not belong to the selected category' error, got: %v", errs)
	}
}

func TestFakeValidateClassificationInvalidSubtypeHierarchy(t *testing.T) {
	fake := NewFake()
	ctx := context.Background()

	// Invalid: Subtype doesn't belong to Type (loft under house, should be under apartment)
	classification := estate.Classification{
		CategoryID: uuid.MustParse("00000000-0000-0000-0001-000000000001"), // residential
		TypeID:     uuid.MustParse("00000000-0000-0000-0002-000000000001"), // house
		SubtypeID:  uuid.MustParse("00000000-0000-0000-0003-000000000002"), // loft (belongs to apartment)
	}

	valid, errs, err := fake.ValidateClassification(ctx, classification)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if valid {
		t.Error("expected invalid classification")
	}

	// Check for specific error
	found := false
	for _, e := range errs {
		if e == "subtype does not belong to the selected type" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'subtype does not belong to the selected type' error, got: %v", errs)
	}
}

func TestFakeValidateClassificationNonExistentIDs(t *testing.T) {
	fake := NewFake()
	ctx := context.Background()

	// Non-existent category
	classification := estate.Classification{
		CategoryID: uuid.New(),
		TypeID:     uuid.MustParse("00000000-0000-0000-0002-000000000001"),
		SubtypeID:  uuid.Nil,
	}

	valid, errs, err := fake.ValidateClassification(ctx, classification)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if valid {
		t.Error("expected invalid classification")
	}
	if len(errs) == 0 {
		t.Error("expected validation errors for non-existent category")
	}
}
