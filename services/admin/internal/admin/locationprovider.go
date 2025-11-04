package admin

import "context"

// LocationProvider defines the behavior required to integrate external geocoding providers.
type LocationProvider interface {
	// ProviderID returns the identifier (e.g. "google", "osm") for the provider.
	ProviderID() string
	// Autocomplete returns address suggestions for the given free-form query.
	Autocomplete(ctx context.Context, query string) ([]LocationSuggestion, error)
	// Resolve resolves a provider reference (e.g. place_id) into a structured address.
	Resolve(ctx context.Context, reference string) (*ResolvedAddress, error)
}

// LocationSuggestion represents a single autocomplete suggestion returned by a provider.
type LocationSuggestion struct {
	Text        string         `json:"text"`
	ProviderRef string         `json:"provider_ref"`
	ProviderURL string         `json:"provider_url,omitempty"`
	Raw         map[string]any `json:"raw,omitempty"`
}

// ResolvedAddress wraps the provider-normalized address payload.
type ResolvedAddress struct {
	Formatted   string         `json:"formatted"`
	Address     Address        `json:"address"`
	Coordinates Coordinates    `json:"coordinates"`
	Provider    string         `json:"provider"`
	ProviderRef string         `json:"provider_ref"`
	ProviderURL string         `json:"provider_url,omitempty"`
	Raw         map[string]any `json:"raw,omitempty"`
}

const (
	ProviderGoogle     = "google"
	ProviderOSM        = "osm"
	ProviderLocationIQ = "locationiq"
)
