package estate

// Features represents the physical characteristics and amenities of a property.
type Features struct {
	// Basic measurements
	TotalArea    float64 `json:"total_area"`     // Total area in square meters
	CoveredArea  float64 `json:"covered_area"`   // Covered/built area in square meters
	LandArea     float64 `json:"land_area"`      // Land/lot area in square meters

	// Rooms
	Bedrooms   int `json:"bedrooms"`
	Bathrooms  int `json:"bathrooms"`
	HalfBaths  int `json:"half_baths"`  // Toilets without shower/tub
	Rooms      int `json:"rooms"`       // Total rooms

	// Parking
	Parking        int  `json:"parking"`          // Number of parking spaces
	CoveredParking int  `json:"covered_parking"`  // Covered/garage spaces

	// Building details
	Floors      int    `json:"floors"`       // Number of floors in the property
	Floor       int    `json:"floor"`        // Floor number (for apartments)
	YearBuilt   int    `json:"year_built"`
	Condition   string `json:"condition"`    // e.g., "new", "excellent", "good", "fair", "needs_work"

	// Amenities (boolean flags)
	Pool          bool `json:"pool"`
	Garden        bool `json:"garden"`
	Balcony       bool `json:"balcony"`
	Terrace       bool `json:"terrace"`
	Elevator      bool `json:"elevator"`
	AirConditioning bool `json:"air_conditioning"`
	Heating       bool `json:"heating"`
	Furnished     bool `json:"furnished"`
	PetFriendly   bool `json:"pet_friendly"`
	Storage       bool `json:"storage"`
	Laundry       bool `json:"laundry"`
	Fireplace     bool `json:"fireplace"`

	// Additional amenities as flexible list
	Amenities []string `json:"amenities,omitempty"` // e.g., ["gym", "security", "concierge"]
}

// Validate performs basic validation on the features.
func (f Features) Validate() []string {
	var errors []string

	if f.TotalArea <= 0 {
		errors = append(errors, "total_area must be greater than 0")
	}

	if f.CoveredArea < 0 {
		errors = append(errors, "covered_area cannot be negative")
	}

	if f.LandArea < 0 {
		errors = append(errors, "land_area cannot be negative")
	}

	if f.CoveredArea > f.TotalArea {
		errors = append(errors, "covered_area cannot be greater than total_area")
	}

	if f.Bedrooms < 0 {
		errors = append(errors, "bedrooms cannot be negative")
	}

	if f.Bathrooms < 0 {
		errors = append(errors, "bathrooms cannot be negative")
	}

	if f.Parking < 0 {
		errors = append(errors, "parking cannot be negative")
	}

	if f.CoveredParking > f.Parking {
		errors = append(errors, "covered_parking cannot be greater than total parking")
	}

	if f.YearBuilt < 0 {
		errors = append(errors, "year_built cannot be negative")
	}

	if f.Floor < 0 {
		errors = append(errors, "floor cannot be negative")
	}

	if f.Floors < 0 {
		errors = append(errors, "floors cannot be negative")
	}

	// Validate condition if provided
	if f.Condition != "" {
		validConditions := map[string]bool{
			"new":        true,
			"excellent":  true,
			"good":       true,
			"fair":       true,
			"needs_work": true,
			"renovation": true,
		}
		if !validConditions[f.Condition] {
			errors = append(errors, "condition must be one of: new, excellent, good, fair, needs_work, renovation")
		}
	}

	return errors
}
