package core

import (
	"context"
	"net/http"
)

// ErrorReporter captures unexpected errors and panics for external services (Sentry, Rollbar, etc.).
type ErrorReporter interface {
	ReportError(ctx context.Context, err error)
	ReportPanic(ctx context.Context, value interface{})
}

type noopErrorReporter struct{}

func (noopErrorReporter) ReportError(context.Context, error)       {}
func (noopErrorReporter) ReportPanic(context.Context, interface{}) {}

var defaultNoopReporter ErrorReporter = noopErrorReporter{}

// NewNoopErrorReporter returns an ErrorReporter that drops all signals.
func NewNoopErrorReporter() ErrorReporter {
	return defaultNoopReporter
}

// NewErrorReportingMiddleware captures panics and forwards them to the provided ErrorReporter.
func NewErrorReportingMiddleware(reporter ErrorReporter) func(http.Handler) http.Handler {
	if reporter == nil {
		reporter = NewNoopErrorReporter()
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					reporter.ReportPanic(r.Context(), rec)
					panic(rec)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// ReportError convenience helper; safe with nil reporter.
func ReportError(ctx context.Context, reporter ErrorReporter, err error) {
	if err == nil {
		return
	}
	if reporter == nil {
		reporter = NewNoopErrorReporter()
	}
	reporter.ReportError(ctx, err)
}
