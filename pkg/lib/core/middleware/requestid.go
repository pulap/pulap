package middleware

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/pulap/pulap/pkg/lib/core"
)

// RequestIDMiddleware ensures every inbound request carries a correlation
// identifier and echoes it back in the response headers.
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get(core.RequestIDHeader)
		if reqID == "" {
			reqID = uuid.NewString()
		}

		ctx := core.WithRequestID(r.Context(), reqID)
		w.Header().Set(core.RequestIDHeader, reqID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
