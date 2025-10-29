package estate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/pkg/lib/telemetry"
	"github.com/pulap/pulap/services/estate/internal/config"
)

const MaxBodyBytes = 1 << 20 // 1 MB

// Handler handles HTTP requests for the Property aggregate.
type Handler struct {
	repo       Repo
	dictClient Client
	xparams    config.XParams
	tlm        *telemetry.HTTP
}

// NewHandler creates a new Handler for Property operations.
func NewHandler(repo Repo, dictClient Client, xparams config.XParams) *Handler {
	return &Handler{
		repo:       repo,
		dictClient: dictClient,
		xparams:    xparams,
		tlm: telemetry.NewHTTP(
			telemetry.WithTracer(xparams.Tracer()),
			telemetry.WithMetrics(xparams.Metrics()),
		),
	}
}

// RegisterRoutes registers all routes for the estate service.
func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/estates", func(r chi.Router) {
		r.Post("/", h.CreateProperty)
		r.Get("/", h.ListProperties)
		r.Get("/{id}", h.GetProperty)
		r.Put("/{id}", h.UpdateProperty)
		r.Delete("/{id}", h.DeleteProperty)
	})
}

// CreateProperty handles POST /estates
func (h *Handler) CreateProperty(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.CreateProperty")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	property, ok := h.decodePropertyPayload(w, r, log)
	if !ok {
		return
	}

	property.EnsureID()
	property.BeforeCreate()

	// Basic validation
	if validationErrors := ValidateCreateProperty(ctx, property); len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Validate classification against dictionary service
	if valid, errs, err := h.dictClient.ValidateClassification(ctx, property.Classification); err != nil {
		log.Error("dictionary service error", "error", err)
		core.RespondError(w, http.StatusBadGateway, "Could not validate classification")
		return
	} else if !valid {
		log.Debug("classification validation failed", "errors", errs)
		core.RespondError(w, http.StatusBadRequest, fmt.Sprintf("Invalid classification: %v", errs))
		return
	}

	// Create in repository
	if err := h.repo.Create(ctx, property); err != nil {
		log.Error("cannot create property", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not create property")
		return
	}

	links := core.RESTfulLinksFor(property)
	w.WriteHeader(http.StatusCreated)
	core.RespondSuccess(w, property, links...)
}

// GetProperty handles GET /estates/{id}
func (h *Handler) GetProperty(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.GetProperty")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	property, err := h.repo.Get(ctx, id)
	if err != nil {
		log.Error("error loading property", "error", err, "id", id.String())
		core.RespondError(w, http.StatusNotFound, "Property not found")
		return
	}

	if property == nil {
		core.RespondError(w, http.StatusNotFound, "Property not found")
		return
	}

	links := core.RESTfulLinksFor(property)
	core.RespondSuccess(w, property, links...)
}

// ListProperties handles GET /estates
func (h *Handler) ListProperties(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.ListProperties")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	// Check for query parameters
	ownerID := r.URL.Query().Get("owner_id")
	status := r.URL.Query().Get("status")

	var properties []*Property
	var err error

	if ownerID != "" {
		properties, err = h.repo.ListByOwner(ctx, ownerID)
	} else if status != "" {
		properties, err = h.repo.ListByStatus(ctx, status)
	} else {
		properties, err = h.repo.List(ctx)
	}

	if err != nil {
		log.Error("error retrieving properties", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve properties")
		return
	}

	core.RespondCollection(w, properties, "estate")
}

// UpdateProperty handles PUT /estates/{id}
func (h *Handler) UpdateProperty(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.UpdateProperty")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	property, ok := h.decodePropertyPayload(w, r, log)
	if !ok {
		return
	}

	property.SetID(id)
	property.BeforeUpdate()

	// Basic validation
	if validationErrors := ValidateUpdateProperty(ctx, id, property); len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Validate classification against dictionary service
	if valid, errs, err := h.dictClient.ValidateClassification(ctx, property.Classification); err != nil {
		log.Error("dictionary service error", "error", err)
		core.RespondError(w, http.StatusBadGateway, "Could not validate classification")
		return
	} else if !valid {
		log.Debug("classification validation failed", "errors", errs)
		core.RespondError(w, http.StatusBadRequest, fmt.Sprintf("Invalid classification: %v", errs))
		return
	}

	// Update in repository
	if err := h.repo.Save(ctx, property); err != nil {
		log.Error("cannot update property", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not update property")
		return
	}

	links := core.RESTfulLinksFor(property)
	core.RespondSuccess(w, property, links...)
}

// DeleteProperty handles DELETE /estates/{id}
func (h *Handler) DeleteProperty(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.DeleteProperty")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	if validationErrors := ValidateDeleteProperty(ctx, id); len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	if err := h.repo.Delete(ctx, id); err != nil {
		log.Error("cannot delete property", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not delete property")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Helper methods

func (h *Handler) log(r *http.Request) core.Logger {
	return h.xparams.Log().With("request_id", r.Context().Value("request_id"))
}

func (h *Handler) parseIDParam(w http.ResponseWriter, r *http.Request, log core.Logger) (uuid.UUID, bool) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		log.Debug("missing id parameter")
		core.RespondError(w, http.StatusBadRequest, "Missing id parameter")
		return uuid.Nil, false
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Debug("invalid id parameter", "id", idStr, "error", err)
		core.RespondError(w, http.StatusBadRequest, "Invalid id parameter")
		return uuid.Nil, false
	}

	return id, true
}

func (h *Handler) decodePropertyPayload(w http.ResponseWriter, r *http.Request, log core.Logger) (*Property, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Debug("error reading request body", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Could not read request body")
		return nil, false
	}

	var property Property
	if err := json.Unmarshal(body, &property); err != nil {
		log.Debug("error decoding JSON", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Invalid JSON payload")
		return nil, false
	}

	return &property, true
}
