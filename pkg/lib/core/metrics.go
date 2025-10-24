package core

import "time"

// Metrics captures high-level service telemetry like request durations.
type Metrics interface {
	ObserveHTTPRequest(path, method string, status int, duration time.Duration)
}

type noopMetrics struct{}

// ObserveHTTPRequest implements Metrics with no side effects.
func (noopMetrics) ObserveHTTPRequest(string, string, int, time.Duration) {}

var defaultNoopMetrics Metrics = noopMetrics{}

// NewNoopMetrics returns a Metrics implementation that drops all observations.
func NewNoopMetrics() Metrics {
	return defaultNoopMetrics
}
*** End Patch
