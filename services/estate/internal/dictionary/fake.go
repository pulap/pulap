package dictionary

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/pulap/pulap/services/estate/internal/estate"
)

var (
	// ErrOptionNotFound is returned when an option is not found.
	ErrOptionNotFound = errors.New("option not found")
)

// Fake is a fake implementation of estate.Client for testing and development.
// It uses hardcoded data from the ADR seed.
type Fake struct {
	options map[uuid.UUID]*estate.Option
	sets    map[string]*estate.Set
}

// NewFake creates a new Fake dictionary client with seed data from the ADR.
func NewFake() *Fake {
	f := &Fake{
		options: make(map[uuid.UUID]*estate.Option),
		sets:    make(map[string]*estate.Set),
	}
	f.seedData()
	return f
}

// GetOption retrieves a single option by ID.
func (f *Fake) GetOption(ctx context.Context, id uuid.UUID) (*estate.Option, error) {
	opt, ok := f.options[id]
	if !ok {
		return nil, ErrOptionNotFound
	}
	return opt, nil
}

// ListOptionsByParent lists all options in a set filtered by parent ID.
func (f *Fake) ListOptionsByParent(ctx context.Context, setName string, parentID *uuid.UUID) ([]estate.Option, error) {
	set, ok := f.sets[setName]
	if !ok {
		return nil, errors.New("set not found")
	}

	var result []estate.Option
	for _, opt := range f.options {
		if opt.SetID != set.ID {
			continue
		}
		// Filter by parent
		if parentID == nil && opt.ParentID == nil {
			result = append(result, *opt)
		} else if parentID != nil && opt.ParentID != nil && *opt.ParentID == *parentID {
			result = append(result, *opt)
		}
	}
	return result, nil
}

// ValidateClassification validates a classification against the dictionary.
func (f *Fake) ValidateClassification(ctx context.Context, c estate.Classification) (bool, []string, error) {
	var errors []string

	// Validate category exists and is active
	category, err := f.GetOption(ctx, c.CategoryID)
	if err != nil {
		errors = append(errors, "category_id not found")
		return false, errors, nil
	}
	if !category.Active {
		errors = append(errors, "category is not active")
	}

	// Validate type exists and is active
	typeOpt, err := f.GetOption(ctx, c.TypeID)
	if err != nil {
		errors = append(errors, "type_id not found")
		return false, errors, nil
	}
	if !typeOpt.Active {
		errors = append(errors, "type is not active")
	}

	// Validate hierarchy: Type.parent_id must equal CategoryID
	if typeOpt.ParentID == nil || *typeOpt.ParentID != c.CategoryID {
		errors = append(errors, "type does not belong to the selected category")
	}

	// Validate subtype if provided
	if c.SubtypeID != uuid.Nil {
		subtype, err := f.GetOption(ctx, c.SubtypeID)
		if err != nil {
			errors = append(errors, "subtype_id not found")
			return false, errors, nil
		}
		if !subtype.Active {
			errors = append(errors, "subtype is not active")
		}

		// Validate hierarchy: Subtype.parent_id must equal TypeID
		if subtype.ParentID == nil || *subtype.ParentID != c.TypeID {
			errors = append(errors, "subtype does not belong to the selected type")
		}
	}

	return len(errors) == 0, errors, nil
}

// seedData populates the fake with data from the ADR.
func (f *Fake) seedData() {
	now := time.Now()

	// Create sets
	categorySet := &estate.Set{
		ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		Name:      "estate_category",
		Label:     "Estate Category",
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	typeSet := &estate.Set{
		ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
		Name:      "estate_type",
		Label:     "Estate Type",
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	subtypeSet := &estate.Set{
		ID:        uuid.MustParse("00000000-0000-0000-0000-000000000003"),
		Name:      "estate_subtype",
		Label:     "Estate Subtype",
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	f.sets["estate_category"] = categorySet
	f.sets["estate_type"] = typeSet
	f.sets["estate_subtype"] = subtypeSet

	// Categories (no parent)
	residential := f.addOption("00000000-0000-0000-0001-000000000001", categorySet.ID, nil, "res", "residential", "Residential", "Residential", 1)
	commercial := f.addOption("00000000-0000-0000-0001-000000000002", categorySet.ID, nil, "com", "commercial", "Commercial", "Commercial", 2)
	land := f.addOption("00000000-0000-0000-0001-000000000003", categorySet.ID, nil, "land", "land", "Land", "Land", 3)
	agricultural := f.addOption("00000000-0000-0000-0001-000000000004", categorySet.ID, nil, "agr", "agricultural", "Agricultural", "Agricultural", 4)
	mixedUse := f.addOption("00000000-0000-0000-0001-000000000005", categorySet.ID, nil, "mix", "mixed_use", "Mixed-use", "Mixed-use", 5)
	specialPurpose := f.addOption("00000000-0000-0000-0001-000000000006", categorySet.ID, nil, "spc", "special_purpose", "Special Purpose", "Special Purpose", 6)

	// Types (parent = category)
	house := f.addOption("00000000-0000-0000-0002-000000000001", typeSet.ID, &residential, "house", "house", "House", "House", 1)
	apartment := f.addOption("00000000-0000-0000-0002-000000000002", typeSet.ID, &residential, "apt", "apartment", "Apartment", "Apartment", 2)
	office := f.addOption("00000000-0000-0000-0002-000000000003", typeSet.ID, &commercial, "off", "office", "Office", "Office", 3)
	retail := f.addOption("00000000-0000-0000-0002-000000000004", typeSet.ID, &commercial, "rtl", "retail", "Retail", "Retail", 4)

	// Subtypes (parent = type)
	f.addOption("00000000-0000-0000-0003-000000000001", subtypeSet.ID, &house, "bglw", "bungalow", "Bungalow", "Bungalow", 1)
	f.addOption("00000000-0000-0000-0003-000000000002", subtypeSet.ID, &apartment, "loft", "loft", "Loft", "Loft", 2)
	f.addOption("00000000-0000-0000-0003-000000000003", subtypeSet.ID, &retail, "shw", "showroom", "Showroom", "Showroom", 3)

	// Prevent unused variable warnings
	_ = land
	_ = agricultural
	_ = mixedUse
	_ = specialPurpose
	_ = office
}

// addOption is a helper to add an option to the fake.
func (f *Fake) addOption(idStr string, setID uuid.UUID, parentID *uuid.UUID, shortCode, key, label, value string, order int) uuid.UUID {
	id := uuid.MustParse(idStr)
	now := time.Now()

	opt := &estate.Option{
		ID:        id,
		SetID:     setID,
		ParentID:  parentID,
		ShortCode: shortCode,
		Key:       key,
		Label:     label,
		Value:     value,
		Order:     order,
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	f.options[id] = opt
	return id
}
