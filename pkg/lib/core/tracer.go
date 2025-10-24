package core

import "context"

// Span represents an in-flight unit of work produced by a Tracer.
type Span interface {
	End()
}

// Tracer starts spans used to trace operations through the system.
type Tracer interface {
	Start(ctx context.Context, spanName string) (context.Context, Span)
}

type noopTracer struct{}

type noopSpan struct{}

func (noopTracer) Start(ctx context.Context, _ string) (context.Context, Span) {
	return ctx, noopSpan{}
}

func (noopSpan) End() {}

var defaultNoopTracer Tracer = noopTracer{}

// NoopTracer returns a Tracer implementation that does nothing.
func NoopTracer() Tracer {
	return defaultNoopTracer
}
