package estate

// Location represents the physical location of a property.
type Location struct {
	Address     Address     `json:"address"`
	Coordinates Coordinates `json:"coordinates"`
	Region      string      `json:"region,omitempty"` // e.g., "LATAM", "North America"
}

// Address represents a structured physical address.
type Address struct {
	Street     string `json:"street"`
	Number     string `json:"number,omitempty"`
	Unit       string `json:"unit,omitempty"`       // Apartment, suite, etc.
	City       string `json:"city"`
	State      string `json:"state,omitempty"`      // Province, state, department
	PostalCode string `json:"postal_code,omitempty"`
	Country    string `json:"country"`
}

// Coordinates represents geographic coordinates.
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// IsZero returns true if coordinates are not set.
func (c Coordinates) IsZero() bool {
	return c.Latitude == 0 && c.Longitude == 0
}

// Validate performs basic validation on the location.
func (l Location) Validate() []string {
	var errors []string

	if l.Address.Street == "" {
		errors = append(errors, "address.street is required")
	}

	if l.Address.City == "" {
		errors = append(errors, "address.city is required")
	}

	if l.Address.Country == "" {
		errors = append(errors, "address.country is required")
	}

	// Coordinates are optional, but if provided must be valid
	if !l.Coordinates.IsZero() {
		if l.Coordinates.Latitude < -90 || l.Coordinates.Latitude > 90 {
			errors = append(errors, "coordinates.latitude must be between -90 and 90")
		}
		if l.Coordinates.Longitude < -180 || l.Coordinates.Longitude > 180 {
			errors = append(errors, "coordinates.longitude must be between -180 and 180")
		}
	}

	return errors
}

// FullAddress returns a formatted full address string.
func (a Address) FullAddress() string {
	result := a.Street
	if a.Number != "" {
		result += " " + a.Number
	}
	if a.Unit != "" {
		result += ", " + a.Unit
	}
	result += ", " + a.City
	if a.State != "" {
		result += ", " + a.State
	}
	if a.PostalCode != "" {
		result += " " + a.PostalCode
	}
	result += ", " + a.Country
	return result
}
