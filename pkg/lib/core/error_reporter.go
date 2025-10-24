package core

import "context"

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
