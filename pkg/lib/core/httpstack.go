package core

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

// StackOptions controls how the shared middleware stack behaves across
// services. All fields are optional; zero values are replaced with safe
// defaults.
type StackOptions struct {
	Timeout time.Duration
	CORS    *CORSOptions
	Metrics Metrics
	Errors  ErrorReporter
}

// ApplyStack wires the shared middleware set onto the provided router. It keeps
// the ordering consistent so request tracing, logging, and panic recovery behave
// predictably in every service.
func ApplyStack(r *chi.Mux, logger Logger, opts StackOptions) {
	if logger == nil {
		logger = NewNoopLogger()
	}
	if opts.Timeout <= 0 {
		opts.Timeout = 30 * time.Second
	}
	if opts.Metrics == nil {
		opts.Metrics = NewNoopMetrics()
	}
	if opts.Errors == nil {
		opts.Errors = NewNoopErrorReporter()
	}

	r.Use(RequestIDMiddleware)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Compress(5))
	r.Use(chimiddleware.Recoverer)
	r.Use(NewErrorReportingMiddleware(opts.Errors))
	r.Use(chimiddleware.Timeout(opts.Timeout))
	r.Use(NewRequestLogger(logger))
	r.Use(NewMetricsMiddleware(opts.Metrics))
	r.Use(chimiddleware.AllowContentType("application/json", "application/x-www-form-urlencoded", "multipart/form-data"))

	if opts.CORS != nil {
		r.Use(CORSMiddleware(*opts.CORS))
	}
}

// RedirectNotFound configures the router to send unknown routes to a fallback
// endpoint. Useful for web frontends where a generic NOT FOUND page feels too raw.
func RedirectNotFound(r *chi.Mux, target string) {
	if target == "" {
		target = "/"
	}

	r.NotFound(func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, target, http.StatusFound)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, target, http.StatusFound)
	})
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
