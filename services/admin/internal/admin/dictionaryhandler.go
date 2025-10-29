package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// ListSets shows all fake sets
func (h *Handler) ListSets(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.ListSets")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	sets, err := h.dictClient.ListSets(ctx)
	if err != nil {
		log.Error("error listing sets", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := h.tmplMgr.Get("list-sets.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "Dictionary Sets",
		"Sets":      sets,
		"ActiveNav": "fake",
		"Template":  "list-sets-content",
	}

	if err := tmpl.ExecuteTemplate(w, "list-sets.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// NewSet shows the form to create a new set
func (h *Handler) NewSet(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.NewSet")
	defer finish()
	log := h.log(r)

	tmpl, err := h.tmplMgr.Get("new-set.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "New Set",
		"ActiveNav": "fake",
		"Template":  "new-set",
	}

	if err := tmpl.ExecuteTemplate(w, "new-set.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// CreateSet handles the creation of a new set
func (h *Handler) CreateSet(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.CreateSet")
	defer finish()
	log := h.log(r)

	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		log.Error("error parsing form", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	req := &CreateSetRequest{
		Name:        r.FormValue("name"),
		Label:       r.FormValue("label"),
		Description: r.FormValue("description"),
		Active:      r.FormValue("active") == "true",
	}

	set, err := h.dictClient.CreateSet(ctx, req)
	if err != nil {
		log.Error("error creating set", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Info("set created successfully", "id", set.ID)
	http.Redirect(w, r, "/list-sets", http.StatusSeeOther)
}

// ShowSet displays a single set
func (h *Handler) ShowSet(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.ShowSet")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid set id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	set, err := h.dictClient.GetSet(ctx, id)
	if err != nil {
		log.Error("error getting set", "error", err, "id", id)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := h.tmplMgr.Get("show-set.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     fmt.Sprintf("Set: %s", set.Label),
		"Set":       set,
		"ActiveNav": "fake",
		"Template":  "show-set",
	}

	if err := tmpl.ExecuteTemplate(w, "show-set.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// EditSet shows the form to edit a set
func (h *Handler) EditSet(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.EditSet")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid set id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	set, err := h.dictClient.GetSet(ctx, id)
	if err != nil {
		log.Error("error getting set", "error", err, "id", id)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := h.tmplMgr.Get("edit-set.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     fmt.Sprintf("Edit: %s", set.Label),
		"Set":       set,
		"ActiveNav": "fake",
		"Template":  "edit-set",
	}

	if err := tmpl.ExecuteTemplate(w, "edit-set.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// UpdateSet handles updating an existing set
func (h *Handler) UpdateSet(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.UpdateSet")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid set id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Error("error parsing form", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	req := &UpdateSetRequest{
		Name:        r.FormValue("name"),
		Label:       r.FormValue("label"),
		Description: r.FormValue("description"),
		Active:      r.FormValue("active") == "true",
	}

	_, err = h.dictClient.UpdateSet(ctx, id, req)
	if err != nil {
		log.Error("error updating set", "error", err, "id", id)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Info("set updated successfully", "id", id)
	http.Redirect(w, r, fmt.Sprintf("/show-set/%s", id), http.StatusSeeOther)
}

// DeleteSet handles deleting a set
func (h *Handler) DeleteSet(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.DeleteSet")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid set id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := h.dictClient.DeleteSet(ctx, id); err != nil {
		log.Error("error deleting set", "error", err, "id", id)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Info("set deleted successfully", "id", id)
	http.Redirect(w, r, "/list-sets", http.StatusSeeOther)
}

// ListOptions shows all fake options
func (h *Handler) ListOptions(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.ListOptions")
	defer finish()
	log := h.log(r)

	ctx := r.Context()

	// Get filter parameters
	setIDStr := r.URL.Query().Get("set_id")
	var setID *uuid.UUID
	if setIDStr != "" {
		id, err := uuid.Parse(setIDStr)
		if err == nil {
			setID = &id
		}
	}

	options, err := h.dictClient.ListOptions(ctx, setID)
	if err != nil {
		log.Error("error listing options", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get all sets for filter dropdown
	sets, err := h.dictClient.ListSets(ctx)
	if err != nil {
		log.Error("error listing sets", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := h.tmplMgr.Get("list-options.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":         "Dictionary Options",
		"Options":       options,
		"Sets":          sets,
		"SelectedSetID": setIDStr,
		"ActiveNav":     "fake",
		"Template":      "list-options-content",
	}

	if err := tmpl.ExecuteTemplate(w, "list-options.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// NewOption shows the form to create a new option
func (h *Handler) NewOption(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.NewOption")
	defer finish()
	log := h.log(r)

	ctx := r.Context()

	// Get all sets for dropdown
	sets, err := h.dictClient.ListSets(ctx)
	if err != nil {
		log.Error("error listing sets", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := h.tmplMgr.Get("new-option.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "New Option",
		"Sets":      sets,
		"ActiveNav": "fake",
		"Template":  "new-option",
	}

	if err := tmpl.ExecuteTemplate(w, "new-option.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// CreateOption handles the creation of a new option
func (h *Handler) CreateOption(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.CreateOption")
	defer finish()
	log := h.log(r)

	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		log.Error("error parsing form", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	setID, _ := uuid.Parse(r.FormValue("set_id"))
	order, _ := strconv.Atoi(r.FormValue("order"))

	var parentID *uuid.UUID
	if parentIDStr := r.FormValue("parent_id"); parentIDStr != "" {
		pid, _ := uuid.Parse(parentIDStr)
		parentID = &pid
	}

	req := &CreateOptionRequest{
		Set:         setID,
		ParentID:    parentID,
		ShortCode:   r.FormValue("short_code"),
		Key:         r.FormValue("key"),
		Label:       r.FormValue("label"),
		Description: r.FormValue("description"),
		Value:       r.FormValue("value"),
		Order:       order,
		Active:      r.FormValue("active") == "true",
	}

	option, err := h.dictClient.CreateOption(ctx, req)
	if err != nil {
		log.Error("error creating option", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Info("option created successfully", "id", option.ID)
	http.Redirect(w, r, "/list-options", http.StatusSeeOther)
}

// ShowOption displays a single option
func (h *Handler) ShowOption(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.ShowOption")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid option id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	option, err := h.dictClient.GetOption(ctx, id)
	if err != nil {
		log.Error("error getting option", "error", err, "id", id)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := h.tmplMgr.Get("show-option.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":       fmt.Sprintf("Option: %s", option.Label),
		"Option":      option,
		"SetName":     option.SetName,
		"ParentLabel": option.ParentLabel,
		"ActiveNav":   "fake",
		"Template":    "show-option",
	}

	if err := tmpl.ExecuteTemplate(w, "show-option.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// EditOption shows the form to edit an option
func (h *Handler) EditOption(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.EditOption")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid option id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	option, err := h.dictClient.GetOption(ctx, id)
	if err != nil {
		log.Error("error getting option", "error", err, "id", id)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Get all sets for dropdown
	sets, err := h.dictClient.ListSets(ctx)
	if err != nil {
		log.Error("error listing sets", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get potential parent options from the same set
	parentOptions, _ := h.dictClient.ListOptions(ctx, &option.Set)

	tmpl, err := h.tmplMgr.Get("edit-option.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":         fmt.Sprintf("Edit: %s", option.Label),
		"Option":        option,
		"Sets":          sets,
		"ParentOptions": parentOptions,
		"ActiveNav":     "fake",
		"Template":      "edit-option",
	}

	if err := tmpl.ExecuteTemplate(w, "edit-option.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// UpdateOption handles updating an existing option
func (h *Handler) UpdateOption(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.UpdateOption")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid option id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Error("error parsing form", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	setID, _ := uuid.Parse(r.FormValue("set_id"))
	order, _ := strconv.Atoi(r.FormValue("order"))

	var parentID *uuid.UUID
	if parentIDStr := r.FormValue("parent_id"); parentIDStr != "" {
		pid, _ := uuid.Parse(parentIDStr)
		parentID = &pid
	}

	req := &UpdateOptionRequest{
		Set:         setID,
		ParentID:    parentID,
		ShortCode:   r.FormValue("short_code"),
		Key:         r.FormValue("key"),
		Label:       r.FormValue("label"),
		Description: r.FormValue("description"),
		Value:       r.FormValue("value"),
		Order:       order,
		Active:      r.FormValue("active") == "true",
	}

	_, err = h.dictClient.UpdateOption(ctx, id, req)
	if err != nil {
		log.Error("error updating option", "error", err, "id", id)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Info("option updated successfully", "id", id)
	http.Redirect(w, r, fmt.Sprintf("/show-option/%s", id), http.StatusSeeOther)
}

// DeleteOption handles deleting an option
func (h *Handler) DeleteOption(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.DeleteOption")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid option id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := h.dictClient.DeleteOption(ctx, id); err != nil {
		log.Error("error deleting option", "error", err, "id", id)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Info("option deleted successfully", "id", id)
	http.Redirect(w, r, "/list-options", http.StatusSeeOther)
}
