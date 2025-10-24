package core

import (
	"context"
	"net/http"
)

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
