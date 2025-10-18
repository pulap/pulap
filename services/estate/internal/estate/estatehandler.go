package estate

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"

	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/services/estate/internal/config"
)

const EstateMaxBodyBytes = 1 << 20

// NewEstateHandler creates a new EstateHandler for the aggregate root.
func NewEstateHandler(repo EstateRepo, xparams config.XParams) *EstateHandler {
	return &EstateHandler{
		repo:    repo,
		xparams: xparams,
	}
}

type EstateHandler struct {
	repo    EstateRepo
	xparams config.XParams
}

func (h *EstateHandler) RegisterRoutes(r chi.Router) {
	r.Route("/estates", func(r chi.Router) {
		r.Post("/", h.CreateEstate)
		r.Get("/", h.GetAllEstates)
		r.Get("/{id}", h.GetEstate)
		r.Put("/{id}", h.UpdateEstate)
		r.Delete("/{id}", h.DeleteEstate)

		// Item operations (part of the aggregate)
		r.Post("/{id}/items", h.AddItemToEstate)
		r.Put("/{id}/items/{childId}", h.UpdateItemInEstate)
		r.Delete("/{id}/items/{childId}", h.RemoveItemFromEstate)

		// Tag operations (part of the aggregate)
		r.Post("/{id}/tags", h.AddTagToEstate)
		r.Put("/{id}/tags/{childId}", h.UpdateTagInEstate)
		r.Delete("/{id}/tags/{childId}", h.RemoveTagFromEstate)

	})
}

func (h *EstateHandler) CreateEstate(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	estate, ok := h.decodeEstatePayload(w, r, log)
	if !ok {
		return
	}

	estate.EnsureID()
	estate.BeforeCreate()

	validationErrors := ValidateCreateEstate(ctx, estate)
	if len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	if err := h.repo.Create(ctx, &estate); err != nil {
		log.Error("cannot create estate", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not create estate")
		return
	}

	// Standard links
	links := core.RESTfulLinksFor(&estate)

	// Child collection links
	links = append(links, core.Link{
		Rel:  "items",
		Href: fmt.Sprintf("/estates/%s/items", estate.ID),
	})

	// Child collection links
	links = append(links, core.Link{
		Rel:  "tags",
		Href: fmt.Sprintf("/estates/%s/tags", estate.ID),
	})

	w.WriteHeader(http.StatusCreated)
	core.RespondSuccess(w, estate, links...)
}

func (h *EstateHandler) GetEstate(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	estate, err := h.repo.Get(ctx, id)
	if err != nil {
		log.Error("error loading estate", "error", err, "id", id.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve estate")
		return
	}

	if estate == nil {
		core.RespondError(w, http.StatusNotFound, "Estate not found")
		return
	}

	// Standard links
	links := core.RESTfulLinksFor(estate)

	// Child collection links
	links = append(links, core.Link{
		Rel:  "items",
		Href: fmt.Sprintf("/estates/%s/items", estate.ID),
	})

	// Child collection links
	links = append(links, core.Link{
		Rel:  "tags",
		Href: fmt.Sprintf("/estates/%s/tags", estate.ID),
	})

	// Child links
	for _, item := range estate.Items {
		childLinks := core.ChildLinksFor(estate, &item)
		// Child entity link
		links = append(links, core.Link{
			Rel:  "item",
			Href: childLinks[0].Href,
		})
	}

	// Child links
	for _, tag := range estate.Tags {
		childLinks := core.ChildLinksFor(estate, &tag)
		// Child entity link
		links = append(links, core.Link{
			Rel:  "tag",
			Href: childLinks[0].Href,
		})
	}

	core.RespondSuccess(w, estate, links...)
}

func (h *EstateHandler) GetAllEstates(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	estates, err := h.repo.Estate(ctx)
	if err != nil {
		log.Error("error retrieving estates", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not estate all estates")
		return
	}

	// Collection response
	core.RespondCollection(w, estates, "estate")
}

func (h *EstateHandler) UpdateEstate(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	estate, ok := h.decodeEstatePayload(w, r, log)
	if !ok {
		return
	}

	estate.SetID(id)
	estate.BeforeUpdate()

	validationErrors := ValidateUpdateEstate(ctx, id, estate)
	if len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors, "id", id.String())
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	if err := h.repo.Save(ctx, &estate); err != nil {
		log.Error("cannot save estate", "error", err, "id", id.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not update estate")
		return
	}

	// Standard links
	links := core.RESTfulLinksFor(&estate)

	// Child collection links
	links = append(links, core.Link{
		Rel:  "items",
		Href: fmt.Sprintf("/estates/%s/items", estate.ID),
	})

	// Child collection links
	links = append(links, core.Link{
		Rel:  "tags",
		Href: fmt.Sprintf("/estates/%s/tags", estate.ID),
	})

	core.RespondSuccess(w, estate, links...)
}

func (h *EstateHandler) DeleteEstate(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	id, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	validationErrors := ValidateDeleteEstate(ctx, id)
	if len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors, "id", id.String())
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	if err := h.repo.Delete(ctx, id); err != nil {
		log.Error("error deleting estate", "error", err, "id", id.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not delete estate")
		return
	}

	// Post-deletion links
	links := core.CollectionLinksFor("estate")
	w.WriteHeader(http.StatusNoContent)
	core.RespondSuccess(w, nil, links...)
}

// Child entity operations (Items)
func (h *EstateHandler) AddItemToEstate(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	estateID, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	item, ok := h.decodeItemPayload(w, r, log)
	if !ok {
		return
	}

	// Load the aggregate
	estate, err := h.repo.Get(ctx, estateID)
	if err != nil {
		log.Error("cannot load estate for adding item", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve estate")
		return
	}

	if estate == nil {
		core.RespondError(w, http.StatusNotFound, "Estate not found")
		return
	}

	// Add item to aggregate
	item.EnsureID()
	item.BeforeCreate()
	estate.Items = append(estate.Items, item)

	// Save the entire aggregate
	if err := h.repo.Save(ctx, estate); err != nil {
		log.Error("error saving estate with new item", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not add item to estate")
		return
	}

	// Child response
	w.WriteHeader(http.StatusCreated)
	core.RespondChild(w, estate, &item)
}

func (h *EstateHandler) UpdateItemInEstate(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	estateID, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	itemID, ok := h.parseItemIDParam(w, r, log)
	if !ok {
		return
	}

	item, ok := h.decodeItemPayload(w, r, log)
	if !ok {
		return
	}

	// Load the aggregate
	estate, err := h.repo.Get(ctx, estateID)
	if err != nil {
		log.Error("cannot load estate for updating item", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve estate")
		return
	}

	if estate == nil {
		core.RespondError(w, http.StatusNotFound, "Estate not found")
		return
	}

	// Find and update item in aggregate
	found := false
	for i, existingItem := range estate.Items {
		if existingItem.ID == itemID {
			item.SetID(itemID)
			item.BeforeUpdate()
			estate.Items[i] = item
			found = true
			break
		}
	}

	if !found {
		core.RespondError(w, http.StatusNotFound, "Item not found in estate")
		return
	}

	// Save the entire aggregate
	if err := h.repo.Save(ctx, estate); err != nil {
		log.Error("error saving estate with updated item", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not update item in estate")
		return
	}

	// Child response
	core.RespondChild(w, estate, &item)
}

func (h *EstateHandler) RemoveItemFromEstate(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	estateID, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	itemID, ok := h.parseItemIDParam(w, r, log)
	if !ok {
		return
	}

	// Load the aggregate
	estate, err := h.repo.Get(ctx, estateID)
	if err != nil {
		log.Error("cannot load estate for removing item", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve estate")
		return
	}

	if estate == nil {
		core.RespondError(w, http.StatusNotFound, "Estate not found")
		return
	}

	// Remove item from aggregate
	found := false
	for i, existingItem := range estate.Items {
		if existingItem.ID == itemID {
			estate.Items = append(estate.Items[:i], estate.Items[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		core.RespondError(w, http.StatusNotFound, "Item not found in estate")
		return
	}

	// Save the entire aggregate
	if err := h.repo.Save(ctx, estate); err != nil {
		log.Error("error saving estate after removing item", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not remove item from estate")
		return
	}

	// Post-deletion links
	links := []core.Link{
		{Rel: "estate", Href: fmt.Sprintf("/estates/%s", estateID)},
		{Rel: "collection", Href: fmt.Sprintf("/estates/%s/items", estateID)},
		{Rel: "create", Href: fmt.Sprintf("/estates/%s/items", estateID)},
	}

	w.WriteHeader(http.StatusNoContent)
	core.RespondSuccess(w, nil, links...)
}

func (h *EstateHandler) parseItemIDParam(w http.ResponseWriter, r *http.Request, log core.Logger) (uuid.UUID, bool) {
	rawID := strings.TrimSpace(chi.URLParam(r, "childId"))
	if rawID == "" {
		core.RespondError(w, http.StatusBadRequest, "Missing childId path parameter")
		return uuid.Nil, false
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		log.Debug("invalid childId parameter", "childId", rawID, "error", err)
		core.RespondError(w, http.StatusBadRequest, "Invalid childId format")
		return uuid.Nil, false
	}

	return id, true
}

func (h *EstateHandler) decodeItemPayload(w http.ResponseWriter, r *http.Request, log core.Logger) (Item, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, EstateMaxBodyBytes)
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var item Item
	if err := dec.Decode(&item); err != nil {
		log.Error("cannot decode item request body", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Item request body could not be decoded")
		return Item{}, false
	}

	return item, true
}

// Child entity operations (Tags)
func (h *EstateHandler) AddTagToEstate(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	estateID, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	tag, ok := h.decodeTagPayload(w, r, log)
	if !ok {
		return
	}

	// Load the aggregate
	estate, err := h.repo.Get(ctx, estateID)
	if err != nil {
		log.Error("cannot load estate for adding tag", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve estate")
		return
	}

	if estate == nil {
		core.RespondError(w, http.StatusNotFound, "Estate not found")
		return
	}

	// Add tag to aggregate
	tag.EnsureID()
	tag.BeforeCreate()
	estate.Tags = append(estate.Tags, tag)

	// Save the entire aggregate
	if err := h.repo.Save(ctx, estate); err != nil {
		log.Error("error saving estate with new tag", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not add tag to estate")
		return
	}

	// Child response
	w.WriteHeader(http.StatusCreated)
	core.RespondChild(w, estate, &tag)
}

func (h *EstateHandler) UpdateTagInEstate(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	estateID, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	tagID, ok := h.parseTagIDParam(w, r, log)
	if !ok {
		return
	}

	tag, ok := h.decodeTagPayload(w, r, log)
	if !ok {
		return
	}

	// Load the aggregate
	estate, err := h.repo.Get(ctx, estateID)
	if err != nil {
		log.Error("cannot load estate for updating tag", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve estate")
		return
	}

	if estate == nil {
		core.RespondError(w, http.StatusNotFound, "Estate not found")
		return
	}

	// Find and update tag in aggregate
	found := false
	for i, existingTag := range estate.Tags {
		if existingTag.ID == tagID {
			tag.SetID(tagID)
			tag.BeforeUpdate()
			estate.Tags[i] = tag
			found = true
			break
		}
	}

	if !found {
		core.RespondError(w, http.StatusNotFound, "Tag not found in estate")
		return
	}

	// Save the entire aggregate
	if err := h.repo.Save(ctx, estate); err != nil {
		log.Error("error saving estate with updated tag", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not update tag in estate")
		return
	}

	// Child response
	core.RespondChild(w, estate, &tag)
}

func (h *EstateHandler) RemoveTagFromEstate(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	estateID, ok := h.parseIDParam(w, r, log)
	if !ok {
		return
	}

	tagID, ok := h.parseTagIDParam(w, r, log)
	if !ok {
		return
	}

	// Load the aggregate
	estate, err := h.repo.Get(ctx, estateID)
	if err != nil {
		log.Error("cannot load estate for removing tag", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not retrieve estate")
		return
	}

	if estate == nil {
		core.RespondError(w, http.StatusNotFound, "Estate not found")
		return
	}

	// Remove tag from aggregate
	found := false
	for i, existingTag := range estate.Tags {
		if existingTag.ID == tagID {
			estate.Tags = append(estate.Tags[:i], estate.Tags[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		core.RespondError(w, http.StatusNotFound, "Tag not found in estate")
		return
	}

	// Save the entire aggregate
	if err := h.repo.Save(ctx, estate); err != nil {
		log.Error("error saving estate after removing tag", "error", err, "estateId", estateID.String())
		core.RespondError(w, http.StatusInternalServerError, "Could not remove tag from estate")
		return
	}

	// Post-deletion links
	links := []core.Link{
		{Rel: "estate", Href: fmt.Sprintf("/estates/%s", estateID)},
		{Rel: "collection", Href: fmt.Sprintf("/estates/%s/tags", estateID)},
		{Rel: "create", Href: fmt.Sprintf("/estates/%s/tags", estateID)},
	}

	w.WriteHeader(http.StatusNoContent)
	core.RespondSuccess(w, nil, links...)
}

func (h *EstateHandler) parseTagIDParam(w http.ResponseWriter, r *http.Request, log core.Logger) (uuid.UUID, bool) {
	rawID := strings.TrimSpace(chi.URLParam(r, "childId"))
	if rawID == "" {
		core.RespondError(w, http.StatusBadRequest, "Missing childId path parameter")
		return uuid.Nil, false
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		log.Debug("invalid childId parameter", "childId", rawID, "error", err)
		core.RespondError(w, http.StatusBadRequest, "Invalid childId format")
		return uuid.Nil, false
	}

	return id, true
}

func (h *EstateHandler) decodeTagPayload(w http.ResponseWriter, r *http.Request, log core.Logger) (Tag, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, EstateMaxBodyBytes)
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var tag Tag
	if err := dec.Decode(&tag); err != nil {
		log.Error("cannot decode tag request body", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Tag request body could not be decoded")
		return Tag{}, false
	}

	return tag, true
}

// Helper methods
func (h *EstateHandler) parseIDParam(w http.ResponseWriter, r *http.Request, log core.Logger) (uuid.UUID, bool) {
	rawID := strings.TrimSpace(chi.URLParam(r, "id"))
	if rawID == "" {
		core.RespondError(w, http.StatusBadRequest, "Missing id path parameter")
		return uuid.Nil, false
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		log.Debug("invalid id parameter", "id", rawID, "error", err)
		core.RespondError(w, http.StatusBadRequest, "Invalid id format")
		return uuid.Nil, false
	}

	return id, true
}

func (h *EstateHandler) decodeEstatePayload(w http.ResponseWriter, r *http.Request, log core.Logger) (Estate, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, EstateMaxBodyBytes)
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var estate Estate
	if err := dec.Decode(&estate); err != nil {
		log.Error("cannot decode request body", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Request body could not be decoded")
		return Estate{}, false
	}

	if err := ensureEstateSingleJSONValue(dec); err != nil {
		log.Error("request body contains extra data", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Request body contains unexpected data")
		return Estate{}, false
	}

	return estate, true
}

func ensureEstateSingleJSONValue(dec *json.Decoder) error {
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return errors.New("additional JSON values detected")
		}
		return err
	}
	return nil
}

func (h *EstateHandler) Log() core.Logger {
	return h.xparams.Log
}

func (h *EstateHandler) logForRequest(r *http.Request) core.Logger {
	return h.xparams.Log.With(
		"request_id", middleware.GetReqID(r.Context()),
		"method", r.Method,
		"path", r.URL.Path,
	)
}

// ValidationError represents a validation error.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateCreateEstate validates a Estate entity before creation.
// TODO: Implement validation logic
func ValidateCreateEstate(ctx context.Context, estate Estate) []ValidationError {
	// TODO: Add validation logic here
	return []ValidationError{}
}

// ValidateUpdateEstate validates a Estate entity before update.
// TODO: Implement validation logic
func ValidateUpdateEstate(ctx context.Context, id uuid.UUID, estate Estate) []ValidationError {
	// TODO: Add validation logic here
	return []ValidationError{}
}

// ValidateDeleteEstate validates a Estate entity before deletion.
// TODO: Implement validation logic
func ValidateDeleteEstate(ctx context.Context, id uuid.UUID) []ValidationError {
	// TODO: Add validation logic here
	return []ValidationError{}
}
