package middleware

import (
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/pulap/pulap/pkg/lib/core"
)

// StackOptions controls how the shared middleware stack behaves across
// services. All fields are optional; zero values are replaced with safe
// defaults.
type StackOptions struct {
	Timeout time.Duration
	CORS    *CORSOptions
}

// ApplyStack wires the shared middleware set onto the provided router. It keeps
// the ordering consistent so request tracing, logging, and panic recovery behave
// predictably in every service.
func ApplyStack(r *chi.Mux, logger core.Logger, opts StackOptions) {
	if logger == nil {
		logger = core.NewNoopLogger()
	}
	if opts.Timeout <= 0 {
		opts.Timeout = 30 * time.Second
	}

	r.Use(RequestIDMiddleware)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Compress(5))
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.Timeout(opts.Timeout))
	r.Use(NewRequestLogger(logger))
	r.Use(chimiddleware.AllowContentType("application/json", "application/x-www-form-urlencoded", "multipart/form-data"))

	if opts.CORS != nil {
		r.Use(CORSMiddleware(*opts.CORS))
	}
}
