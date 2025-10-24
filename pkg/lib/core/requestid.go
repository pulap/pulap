package core

import "context"

// RequestIDHeader identifies the header used to propagate request correlation IDs.
const RequestIDHeader = "X-Request-ID"

type requestIDKeyType struct{}

var requestIDKey requestIDKeyType

// WithRequestID attaches a request ID to the provided context.
func WithRequestID(ctx context.Context, id string) context.Context {
	if id == "" {
		return ctx
	}
	return context.WithValue(ctx, requestIDKey, id)
}

// RequestIDFrom extracts the request ID from the context, returning an empty
// string when the ID is not present.
func RequestIDFrom(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v, ok := ctx.Value(requestIDKey).(string); ok {
		return v
	}
	return ""
}
