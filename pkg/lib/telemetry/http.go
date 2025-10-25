package telemetry

import (
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/pulap/pulap/pkg/lib/core"
)

// Interfaces keep compatibility with existing tracer/metrics contracts while
// allowing telemetry to own the instrumentation helpers.
type Tracer = core.Tracer
type Span = core.Span
type Metrics = core.Metrics

func NoopTracer() Tracer { return core.NoopTracer() }

func NewNoopMetrics() Metrics { return core.NewNoopMetrics() }

// HTTP instruments HTTP handlers with tracing and metrics.
type HTTP struct {
	tracer  Tracer
	metrics Metrics
}

// Option mutates HTTP configuration.
type Option func(*HTTP)

// NewHTTP builds an HTTP instrumentation helper with optional custom deps.
func NewHTTP(opts ...Option) *HTTP {
	h := &HTTP{
		tracer:  NoopTracer(),
		metrics: NewNoopMetrics(),
	}
	for _, opt := range opts {
		if opt != nil {
			opt(h)
		}
	}
	return h
}

// WithTracer overrides the tracer implementation used by HTTP spans.
func WithTracer(t Tracer) Option {
	return func(h *HTTP) {
		if t == nil {
			t = NoopTracer()
		}
		h.tracer = t
	}
}

// WithMetrics overrides the metrics collector used by HTTP instrumentation.
func WithMetrics(m Metrics) Option {
	return func(h *HTTP) {
		if m == nil {
			m = NewNoopMetrics()
		}
		h.metrics = m
	}
}

// Start wraps the request/response pair, starting a span and returning a finish
// function that records duration and status.
func (h *HTTP) Start(w http.ResponseWriter, r *http.Request, spanName string) (http.ResponseWriter, *http.Request, func()) {
	tracer := h.tracer
	if tracer == nil {
		tracer = NoopTracer()
	}
	metrics := h.metrics
	if metrics == nil {
		metrics = NewNoopMetrics()
	}

	ctx, span := tracer.Start(r.Context(), spanName)
	rw := chimiddleware.NewWrapResponseWriter(w, r.ProtoMajor)
	start := time.Now()
	reqWithCtx := r.WithContext(ctx)

	finish := func() {
		span.End()
		metrics.ObserveHTTPRequest(reqWithCtx.URL.Path, reqWithCtx.Method, rw.Status(), time.Since(start))
	}

	return rw, reqWithCtx, finish
}
