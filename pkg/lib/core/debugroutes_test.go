package core

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestDebugRoutes(t *testing.T) {
	t.Setenv("HM_DEBUG_ROUTES", "1")

	mux := chi.NewRouter()

	ApplyStack(mux, NewNoopLogger(), StackOptions{})
	RegisterDebugRoutes(mux)
	mux.Get("/foo", func(w http.ResponseWriter, r *http.Request) {})

	req := httptest.NewRequest(http.MethodGet, debugRoutesPath, nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if body == "" {
		t.Fatalf("expected body, got empty string")
	}
}
