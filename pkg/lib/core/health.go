package core

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

type HealthCheck func(context.Context) error

type HealthChecks struct {
	Liveness  map[string]HealthCheck
	Readiness map[string]HealthCheck
}

type HealthReporter interface {
	HealthChecks() HealthChecks
}

type HealthRegistry struct {
	mu        sync.RWMutex
	liveness  map[string]HealthCheck
	readiness map[string]HealthCheck
}

func NewHealthRegistry() *HealthRegistry {
	return &HealthRegistry{
		liveness:  map[string]HealthCheck{},
		readiness: map[string]HealthCheck{},
	}
}

func (hr *HealthRegistry) RegisterChecks(checks HealthChecks) {
	for name, check := range checks.Liveness {
		hr.RegisterLiveness(name, check)
	}
	for name, check := range checks.Readiness {
		hr.RegisterReadiness(name, check)
	}
}

func (hr *HealthRegistry) RegisterLiveness(name string, check HealthCheck) {
	if check == nil {
		return
	}
	hr.mu.Lock()
	hr.liveness[name] = check
	hr.mu.Unlock()
}

func (hr *HealthRegistry) RegisterReadiness(name string, check HealthCheck) {
	if check == nil {
		return
	}
	hr.mu.Lock()
	hr.readiness[name] = check
	hr.mu.Unlock()
}

func RegisterProbes(r chi.Router, registry *HealthRegistry) {
	if registry == nil {
		registry = NewHealthRegistry()
	}

	r.Get("/healthz", makeHealthHandler(registry.liveness))
	r.Get("/livez", makeHealthHandler(registry.liveness))
	r.Get("/readyz", makeHealthHandler(registry.readiness))
	r.Get("/ping", notImplementedHandler)
	r.Get("/metrics", notImplementedHandler)
	r.Get("/version", notImplementedHandler)
}

func makeHealthHandler(checks map[string]HealthCheck) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		summary := runChecks(r.Context(), checks)
		status := http.StatusOK
		for _, res := range summary.Results {
			if res.Error != "" {
				status = http.StatusServiceUnavailable
				break
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(summary)
	}
}

func runChecks(ctx context.Context, checks map[string]HealthCheck) ProbeResponse {
	results := make([]HealthResult, 0, len(checks))
	for name, check := range checks {
		result := HealthResult{Name: name, Error: ""}
		if err := check(ctx); err != nil {
			result.Error = err.Error()
		}
		results = append(results, result)
	}

	status := "ok"
	for _, res := range results {
		if res.Error != "" {
			status = "degraded"
			break
		}
	}

	return ProbeResponse{
		Status:    status,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Results:   results,
	}
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

type HealthResult struct {
	Name  string `json:"name"`
	Error string `json:"error,omitempty"`
}

type ProbeResponse struct {
	Status    string         `json:"status"`
	Timestamp string         `json:"timestamp"`
	Results   []HealthResult `json:"results,omitempty"`
}
