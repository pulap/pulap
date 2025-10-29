package dictionary

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/pkg/lib/telemetry"
	"github.com/pulap/pulap/services/dictionary/internal/config"
)

const MaxBodyBytes = 1 << 20 // 1 MB

// Handler handles HTTP requests for the Dictionary service.
type Handler struct {
	setRepo    SetRepo
	optionRepo OptionRepo
	xparams    config.XParams
	tlm        *telemetry.HTTP
}

// NewHandler creates a new Handler for Dictionary operations.
func NewHandler(setRepo SetRepo, optionRepo OptionRepo, xparams config.XParams) *Handler {
	return &Handler{
		setRepo:    setRepo,
		optionRepo: optionRepo,
		xparams:    xparams,
		tlm: telemetry.NewHTTP(
			telemetry.WithTracer(xparams.Tracer()),
			telemetry.WithMetrics(xparams.Metrics()),
		),
	}
}

// RegisterRoutes registers all routes for the fake service.
func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/fake", func(r chi.Router) {
		// Set routes
		r.Route("/sets", func(r chi.Router) {
			r.Post("/", h.CreateSet)
			r.Get("/", h.ListSets)
			r.Get("/{id}", h.GetSet)
			r.Put("/{id}", h.UpdateSet)
			r.Delete("/{id}", h.DeleteSet)
			r.Get("/name/{name}", h.GetSetByName)
		})

		// Option routes
		r.Route("/options", func(r chi.Router) {
			r.Post("/", h.CreateOption)
			r.Get("/", h.ListOptions)
			r.Get("/{id}", h.GetOption)
			r.Put("/{id}", h.UpdateOption)
			r.Delete("/{id}", h.DeleteOption)
			r.Get("/set/{setName}", h.ListOptionsBySetName)
			r.Get("/set/{setName}/parent/{parentID}", h.ListOptionsBySetAndParent)
		})
	})
}

// Set Handlers

// CreateSet handles POST /fake/sets
func (h *Handler) CreateSet(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.CreateSet")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	set, ok := h.decodeSetPayload(w, r, log)
	if !ok {
		return
	}

	set.EnsureID()
	set.BeforeCreate()

	// Validation
	if validationErrors := ValidateCreateSet(ctx, set); len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Create in repository
	if err := h.setRepo.Create(ctx, set); err != nil {
		log.Error("cannot create set", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not create set")
		return
	}

	links := core.RESTfulLinksFor(set)
	w.WriteHeader(http.StatusCreated)
	core.RespondSuccess(w, set, links...)
}

// GetSet handles GET /fake/sets/{id}
func (h *Handler) GetSet(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.GetSet")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	set, err := h.setRepo.Get(ctx, id)
	if err != nil {
		log.Error("error loading set", "error", err, "id", id.String())
		core.RespondError(w, http.StatusNotFound, "Set not found")
		return
	}

	if set == nil {
		core.RespondError(w, http.StatusNotFound, "Set not found")
		return
	}

	links := core.RESTfulLinksFor(set)
	core.RespondSuccess(w, set, links...)
}

// GetSetByName handles GET /fake/sets/name/{name}
func (h *Handler) GetSetByName(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.GetSetByName")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	name := chi.URLParam(r, "name")
	if name == "" {
		log.Debug("missing name parameter")
		core.RespondError(w, http.StatusBadRequest, "Missing name parameter")
		return
	}

	set, err := h.setRepo.GetByName(ctx, name)
	if err != nil {
		log.Error("error loading set by name", "error", err, "name", name)
		core.RespondError(w, http.StatusNotFound, "Set not found")
		return
	}

	if set == nil {
		core.RespondError(w, http.StatusNotFound, "Set not found")
		return
	}

	links := core.RESTfulLinksFor(set)
	core.RespondSuccess(w, set, links...)
}

// ListSets handles GET /fake/sets
func (h *Handler) ListSets(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.ListSets")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	// Check for active query parameter
	activeOnly := r.URL.Query().Get("active") == "true"

	var sets []*Set
	var err error

	if activeOnly {
		sets, err = h.setRepo.ListActive(ctx)
	} else {
		sets, err = h.setRepo.List(ctx)
	}

	if err != nil {
		log.Error("error retrieving sets", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve sets")
		return
	}

	core.RespondCollection(w, sets, "fake/set")
}

// UpdateSet handles PUT /fake/sets/{id}
func (h *Handler) UpdateSet(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.UpdateSet")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	set, ok := h.decodeSetPayload(w, r, log)
	if !ok {
		return
	}

	set.SetID(id)
	set.BeforeUpdate()

	// Validation
	if validationErrors := ValidateUpdateSet(ctx, id, set); len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Update in repository
	if err := h.setRepo.Save(ctx, set); err != nil {
		log.Error("cannot update set", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not update set")
		return
	}

	links := core.RESTfulLinksFor(set)
	core.RespondSuccess(w, set, links...)
}

// DeleteSet handles DELETE /fake/sets/{id}
func (h *Handler) DeleteSet(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.DeleteSet")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	if validationErrors := ValidateDeleteSet(ctx, id); len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	if err := h.setRepo.Delete(ctx, id); err != nil {
		log.Error("cannot delete set", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not delete set")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Option Handlers

// CreateOption handles POST /fake/options
func (h *Handler) CreateOption(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.CreateOption")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	option, ok := h.decodeOptionPayload(w, r, log)
	if !ok {
		return
	}

	option.EnsureID()
	option.BeforeCreate()

	// Validation
	if validationErrors := ValidateCreateOption(ctx, option); len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Create in repository
	if err := h.optionRepo.Create(ctx, option); err != nil {
		log.Error("cannot create option", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not create option")
		return
	}

	links := core.RESTfulLinksFor(option)
	w.WriteHeader(http.StatusCreated)
	core.RespondSuccess(w, option, links...)
}

// GetOption handles GET /fake/options/{id}
func (h *Handler) GetOption(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.GetOption")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	option, err := h.optionRepo.Get(ctx, id)
	if err != nil {
		log.Error("error loading option", "error", err, "id", id.String())
		core.RespondError(w, http.StatusNotFound, "Option not found")
		return
	}

	if option == nil {
		core.RespondError(w, http.StatusNotFound, "Option not found")
		return
	}

	links := core.RESTfulLinksFor(option)
	core.RespondSuccess(w, option, links...)
}

// ListOptions handles GET /fake/options
func (h *Handler) ListOptions(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.ListOptions")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	// Check for active query parameter
	activeOnly := r.URL.Query().Get("active") == "true"

	var options []*Option
	var err error

	if activeOnly {
		options, err = h.optionRepo.ListActive(ctx)
	} else {
		options, err = h.optionRepo.List(ctx)
	}

	if err != nil {
		log.Error("error retrieving options", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve options")
		return
	}

	core.RespondCollection(w, options, "fake/option")
}

// ListOptionsBySetName handles GET /fake/options/set/{setName}
func (h *Handler) ListOptionsBySetName(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.ListOptionsBySetName")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	setName := chi.URLParam(r, "setName")
	if setName == "" {
		log.Debug("missing setName parameter")
		core.RespondError(w, http.StatusBadRequest, "Missing setName parameter")
		return
	}

	options, err := h.optionRepo.ListBySetName(ctx, setName)
	if err != nil {
		log.Error("error retrieving options by set name", "error", err, "setName", setName)
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve options")
		return
	}

	core.RespondCollection(w, options, "fake/option")
}

// ListOptionsBySetAndParent handles GET /fake/options/set/{setName}/parent/{parentID}
func (h *Handler) ListOptionsBySetAndParent(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.ListOptionsBySetAndParent")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	setName := chi.URLParam(r, "setName")
	if setName == "" {
		log.Debug("missing setName parameter")
		core.RespondError(w, http.StatusBadRequest, "Missing setName parameter")
		return
	}

	parentIDStr := chi.URLParam(r, "parentID")
	var parentID *uuid.UUID
	if parentIDStr != "" && parentIDStr != "null" {
		pid, err := uuid.Parse(parentIDStr)
		if err != nil {
			log.Debug("invalid parentID parameter", "parentID", parentIDStr, "error", err)
			core.RespondError(w, http.StatusBadRequest, "Invalid parentID parameter")
			return
		}
		parentID = &pid
	}

	options, err := h.optionRepo.ListBySetAndParent(ctx, setName, parentID)
	if err != nil {
		log.Error("error retrieving options by set and parent", "error", err, "setName", setName, "parentID", parentIDStr)
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve options")
		return
	}

	core.RespondCollection(w, options, "fake/option")
}

// UpdateOption handles PUT /fake/options/{id}
func (h *Handler) UpdateOption(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.UpdateOption")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	option, ok := h.decodeOptionPayload(w, r, log)
	if !ok {
		return
	}

	option.SetID(id)
	option.BeforeUpdate()

	// Validation
	if validationErrors := ValidateUpdateOption(ctx, id, option); len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Update in repository
	if err := h.optionRepo.Save(ctx, option); err != nil {
		log.Error("cannot update option", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not update option")
		return
	}

	links := core.RESTfulLinksFor(option)
	core.RespondSuccess(w, option, links...)
}

// DeleteOption handles DELETE /fake/options/{id}
func (h *Handler) DeleteOption(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "Handler.DeleteOption")
	defer finish()
	log := h.log(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	if validationErrors := ValidateDeleteOption(ctx, id); len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	if err := h.optionRepo.Delete(ctx, id); err != nil {
		log.Error("cannot delete option", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not delete option")
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

func (h *Handler) decodeSetPayload(w http.ResponseWriter, r *http.Request, log core.Logger) (*Set, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Debug("error reading request body", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Could not read request body")
		return nil, false
	}

	var set Set
	if err := json.Unmarshal(body, &set); err != nil {
		log.Debug("error decoding JSON", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Invalid JSON payload")
		return nil, false
	}

	return &set, true
}

func (h *Handler) decodeOptionPayload(w http.ResponseWriter, r *http.Request, log core.Logger) (*Option, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Debug("error reading request body", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Could not read request body")
		return nil, false
	}

	var option Option
	if err := json.Unmarshal(body, &option); err != nil {
		log.Debug("error decoding JSON", "error", err)
		core.RespondError(w, http.StatusBadRequest, fmt.Sprintf("Invalid JSON payload: %v", err))
		return nil, false
	}

	return &option, true
}
