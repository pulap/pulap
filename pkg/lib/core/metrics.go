package core

import (
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

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

// NewMetricsMiddleware measures request durations and reports them through the
// provided Metrics implementation.
func NewMetricsMiddleware(metrics Metrics) func(http.Handler) http.Handler {
	if metrics == nil {
		metrics = NewNoopMetrics()
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			rw := chimiddleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(rw, r)

			duration := time.Since(start)
			metrics.ObserveHTTPRequest(r.URL.Path, r.Method, rw.Status(), duration)
		})
	}
}
