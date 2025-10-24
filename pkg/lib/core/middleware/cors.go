package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pulap/pulap/pkg/lib/core"
)

// CORSOptions describes the configuration for the lightweight CORS middleware
// bundled with the core package. It intentionally mirrors the most common
// options from github.com/go-chi/cors so we can swap implementations later if
// needed.
type CORSOptions struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	ExposedHeaders   []string
	AllowCredentials bool
	MaxAge           time.Duration
}

// DefaultCORSOptions returns permissive defaults suited for internal services
// and development environments.
func DefaultCORSOptions() CORSOptions {
	return CORSOptions{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-Requested-With", "User-Agent", core.RequestIDHeader},
		ExposedHeaders: []string{"Content-Length", core.RequestIDHeader},
		MaxAge:         10 * time.Minute,
	}
}

// CORSMiddleware produces a middleware that applies the provided CORS options.
func CORSMiddleware(opts CORSOptions) func(http.Handler) http.Handler {
	allowedMethods := strings.Join(opts.AllowedMethods, ", ")
	allowedHeaders := strings.Join(opts.AllowedHeaders, ", ")
	exposedHeaders := strings.Join(opts.ExposedHeaders, ", ")
	maxAge := int(opts.MaxAge.Seconds())

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin == "" {
				next.ServeHTTP(w, r)
				return
			}

			if !isOriginAllowed(origin, opts.AllowedOrigins) {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			setVaryHeaders(w.Header())
			applyCORSHeaders(w.Header(), origin, opts.AllowedOrigins, allowedMethods, allowedHeaders, exposedHeaders, opts.AllowCredentials, maxAge)

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func isOriginAllowed(requestOrigin string, allowed []string) bool {
	if len(allowed) == 0 {
		return false
	}

	requestOrigin = strings.ToLower(requestOrigin)
	for _, origin := range allowed {
		origin = strings.ToLower(origin)
		if origin == "*" || origin == requestOrigin {
			return true
		}
	}

	return false
}

func setVaryHeaders(h http.Header) {
	vary := h.Values("Vary")
	needed := map[string]struct{}{"Origin": {}, "Access-Control-Request-Method": {}, "Access-Control-Request-Headers": {}}

	existing := map[string]struct{}{}
	for _, v := range vary {
		for _, token := range strings.Split(v, ",") {
			token = strings.TrimSpace(token)
			if token != "" {
				existing[token] = struct{}{}
			}
		}
	}

	for token := range needed {
		if _, ok := existing[token]; !ok {
			h.Add("Vary", token)
		}
	}
}

func applyCORSHeaders(h http.Header, requestOrigin string, allowedOrigins []string, methods, headers, exposed string, allowCredentials bool, maxAge int) {
	originValue := originHeaderValue(requestOrigin, allowedOrigins)
	if originValue == "*" && allowCredentials {
		originValue = requestOrigin
	}
	h.Set("Access-Control-Allow-Origin", originValue)

	if methods != "" {
		h.Set("Access-Control-Allow-Methods", methods)
	}
	if headers != "" {
		h.Set("Access-Control-Allow-Headers", headers)
	}
	if exposed != "" {
		h.Set("Access-Control-Expose-Headers", exposed)
	}

	if allowCredentials {
		h.Set("Access-Control-Allow-Credentials", "true")
	}

	if maxAge > 0 {
		h.Set("Access-Control-Max-Age", strconv.Itoa(maxAge))
	}
}

func originHeaderValue(requestOrigin string, allowed []string) string {
	for _, origin := range allowed {
		if origin == "*" {
			return "*"
		}
		if strings.EqualFold(origin, requestOrigin) {
			return requestOrigin
		}
	}
	return requestOrigin
}
