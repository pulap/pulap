package core

import (
	"net/http"

	"github.com/google/uuid"
)

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
