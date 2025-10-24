package core

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

// StackParams exposes the minimum surface required by the router helpers.
type StackParams interface {
	Log() Logger
}

// NewRouter returns a chi router preconfigured with the standard middleware
// stack. Optional StackOptions can be provided to tweak defaults (timeout,
// CORS, …).
func NewRouter(params StackParams, opts ...StackOptions) *chi.Mux {
	r := chi.NewRouter()

	var stack StackOptions
	if len(opts) > 0 {
		stack = opts[0]
	}

	ApplyStack(r, params.Log(), stack)
	return r
}

// NewWebRouter builds upon NewRouter with extra helpers suited for browser
// frontends (no-cache headers + fallback redirect for unknown routes).
func NewWebRouter(params StackParams, fallback string, opts ...StackOptions) *chi.Mux {
	r := NewRouter(params, opts...)

	r.Use(chimiddleware.NoCache)
	RedirectNotFound(r, fallback)

	return r
}
