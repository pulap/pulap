package core

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

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

// RequestIDMiddleware ensures every inbound request carries a correlation
// identifier and echoes it back in the response headers.
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get(RequestIDHeader)
		if reqID == "" {
			reqID = uuid.NewString()
		}

		ctx := WithRequestID(r.Context(), reqID)
		w.Header().Set(RequestIDHeader, reqID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
