package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

// Server represents an HTTP server with graceful shutdown capabilities.
type Server struct {
	Router *chi.Mux
	Opts   ServerOpts
	Logger Logger
	srv    *http.Server
}

// Start starts the HTTP server.
func (s *Server) Start(ctx context.Context) error {
	s.srv = &http.Server{
		Addr:    s.Opts.Port,
		Handler: s.Router,
	}

	s.Logger.Info(fmt.Sprintf("Starting server on %s...", s.Opts.Port))
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not listen and serve: %w", err)
	}
	return nil
}

// Stop gracefully stops the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	if s.srv == nil {
		return nil
	}

	s.Logger.Info("Shutting down server...")
	return s.srv.Shutdown(ctx)
}

// Serve starts an HTTP server and handles graceful shutdown.
// It also calls the provided stops functions during shutdown.
func Serve(router *chi.Mux, opts ServerOpts, stops []func(context.Context) error, log Logger) {
	srv := &http.Server{
		Addr:    opts.Port,
		Handler: router,
	}

	go func() {
		log.Info(fmt.Sprintf("Starting server on %s...", opts.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Sprintf("could not listen and serve: %v", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")
	Shutdown(srv, stops)
}

// ServerOpts holds server-related options.
type ServerOpts struct {
	Port string
}

// ProbeResponse represents the JSON response for probe endpoints.
type ProbeResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// RegisterProbes adds standard probe endpoints to the router.
// Note: Current implementation provides basic liveness/readiness checks.
// Future enhancement: aggregate health status from all dependencies (database, external services, etc.)
func RegisterProbes(r chi.Router) {
	r.Get("/healthz", healthzHandler)
	r.Get("/readyz", readyzHandler)
	r.Get("/livez", livezHandler)
	r.Get("/ping", notImplementedHandler)
	r.Get("/metrics", notImplementedHandler)
	r.Get("/version", notImplementedHandler)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	respondProbe(w, http.StatusOK, "ok")
}

func readyzHandler(w http.ResponseWriter, r *http.Request) {
	respondProbe(w, http.StatusOK, "ready")
}

func livezHandler(w http.ResponseWriter, r *http.Request) {
	respondProbe(w, http.StatusOK, "alive")
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func respondProbe(w http.ResponseWriter, status int, statusMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ProbeResponse{
		Status:    statusMsg,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}
