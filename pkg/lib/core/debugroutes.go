package core

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

// TODO: Move to config values
const debugRoutesEnv = "HM_DEBUG_ROUTES"
const debugRoutesPath = "/debug/routes"

// RegisterDebugRoutes exposes a routes debugger when HM_DEBUG_ROUTES is
// enabled. The handler enumerates every route currently registered on the
// router and returns them as JSON.
func RegisterDebugRoutes(r chi.Router) {
	if !debugRoutesEnabled() {
		return
	}

	_ = attachDebugRoutesMiddleware(r)
}

func attachDebugRoutesMiddleware(r chi.Router) (ok bool) {
	defer func() {
		if rec := recover(); rec != nil {
			ok = false
		}
	}()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.Method == http.MethodGet && req.URL.Path == debugRoutesPath {
				writeDebugRoutes(w, r)
				return
			}
			next.ServeHTTP(w, req)
		})
	})

	return true
}

func writeDebugRoutes(w http.ResponseWriter, r chi.Router) {
	var out []map[string]string
	chi.Walk(r, func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		out = append(out, map[string]string{
			"method": method,
			"path":   route,
		})
		return nil
	})

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(out)
}

func debugRoutesEnabled() bool {
	value := strings.TrimSpace(strings.ToLower(os.Getenv(debugRoutesEnv)))
	switch value {
	case "1", "true", "yes", "on":
		return true
	}
	return false
}
