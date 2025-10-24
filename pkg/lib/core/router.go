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
// stack using default StackOptions.
func NewRouter(xparams StackParams) *chi.Mux {
	return NewRouterWithOptions(StackOptions{}, xparams)
}

// NewRouterWithOptions behaves like NewRouter but allows customising the stack.
func NewRouterWithOptions(opts StackOptions, xparams StackParams) *chi.Mux {
	r := chi.NewRouter()
	ApplyStack(r, xparams.Log(), opts)
	return r
}

// NewWebRouter builds upon NewRouter with extra helpers suited for browser
// frontends (NoCache + fallback redirect for unknown routes).
func NewWebRouter(fallback string, xparams StackParams) *chi.Mux {
	return NewWebRouterWithOptions(fallback, StackOptions{}, xparams)
}

// NewWebRouterWithOptions mirrors NewRouterWithOptions for web frontends.
func NewWebRouterWithOptions(fallback string, opts StackOptions, xparams StackParams) *chi.Mux {
	r := NewRouterWithOptions(opts, xparams)

	r.Use(chimiddleware.NoCache)
	RedirectNotFound(r, fallback)

	return r
}
