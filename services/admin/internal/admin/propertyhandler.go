package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// ListProperties shows all properties
func (h *Handler) ListProperties(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.ListProperties")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	properties, err := h.service.ListProperties(ctx)
	if err != nil {
		log.Error("error listing properties", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := h.tmplMgr.Get("list-properties.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":      "Properties",
		"Properties": properties,
		"ActiveNav":  "properties",
		"Template":   "list-properties-content",
	}

	if err := tmpl.ExecuteTemplate(w, "list-properties.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// NewProperty shows the form to create a new property
func (h *Handler) NewProperty(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.NewProperty")
	defer finish()
	log := h.log(r)

	ctx := r.Context()

	// Fetch fake options
	categories, err := h.dictClient.ListCategories(ctx)
	if err != nil {
		log.Error("error fetching categories", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	types, err := h.dictClient.ListTypesByCategory(ctx, uuid.Nil)
	if err != nil {
		log.Error("error fetching types", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	subtypes, err := h.dictClient.ListSubtypesByType(ctx, uuid.Nil)
	if err != nil {
		log.Error("error fetching subtypes", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := h.tmplMgr.Get("new-property.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":      "New Property",
		"ActiveNav":  "properties",
		"Template":   "new-property",
		"Categories": DictionaryOptionsToMap(categories),
		"Types":      DictionaryOptionsToMap(types),
		"Subtypes":   DictionaryOptionsToMap(subtypes),
	}

	if err := tmpl.ExecuteTemplate(w, "new-property.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// CreateProperty handles the creation of a new property
func (h *Handler) CreateProperty(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.CreateProperty")
	defer finish()
	log := h.log(r)

	ctx := r.Context()

	// Parse form data
	if err := r.ParseForm(); err != nil {
		log.Error("error parsing form", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Parse classification IDs
	categoryID, _ := uuid.Parse(r.FormValue("category_id"))
	typeID, _ := uuid.Parse(r.FormValue("type_id"))
	subtypeIDStr := r.FormValue("subtype_id")
	var subtypeID uuid.UUID
	if subtypeIDStr != "" {
		subtypeID, _ = uuid.Parse(subtypeIDStr)
	}

	// Parse numeric fields
	totalArea, _ := strconv.ParseFloat(r.FormValue("total_area"), 64)
	bedrooms, _ := strconv.Atoi(r.FormValue("bedrooms"))
	bathrooms, _ := strconv.Atoi(r.FormValue("bathrooms"))
	parking, _ := strconv.Atoi(r.FormValue("parking"))
	amount, _ := strconv.ParseFloat(r.FormValue("price_amount"), 64)

	req := &CreatePropertyRequest{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Classification: Classification{
			CategoryID: categoryID,
			TypeID:     typeID,
			SubtypeID:  subtypeID,
		},
		Location: Location{
			Address: Address{
				Street:     r.FormValue("street"),
				Number:     r.FormValue("number"),
				City:       r.FormValue("city"),
				State:      r.FormValue("state"),
				PostalCode: r.FormValue("postal_code"),
				Country:    r.FormValue("country"),
			},
			Region: r.FormValue("region"),
		},
		Features: Features{
			TotalArea: totalArea,
			Bedrooms:  bedrooms,
			Bathrooms: bathrooms,
			Parking:   parking,
		},
		Price: Price{
			Amount:   amount,
			Currency: r.FormValue("currency"),
			Type:     r.FormValue("price_type"),
		},
		Status:  r.FormValue("status"),
		OwnerID: r.FormValue("owner_id"),
	}

	property, err := h.service.CreateProperty(ctx, req)
	if err != nil {
		log.Error("error creating property", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Info("property created successfully", "id", property.ID)
	http.Redirect(w, r, "/list-properties", http.StatusSeeOther)
}

// ShowProperty displays a single property
func (h *Handler) ShowProperty(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.ShowProperty")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid property id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	property, err := h.service.GetProperty(ctx, id)
	if err != nil {
		log.Error("error getting property", "error", err, "id", id)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := h.tmplMgr.Get("show-property.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     fmt.Sprintf("Property: %s", property.Name),
		"Property":  property,
		"ActiveNav": "properties",
		"Template":  "show-property",
	}

	if err := tmpl.ExecuteTemplate(w, "show-property.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// EditProperty shows the form to edit a property
func (h *Handler) EditProperty(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.EditProperty")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid property id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	property, err := h.service.GetProperty(ctx, id)
	if err != nil {
		log.Error("error getting property", "error", err, "id", id)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Fetch fake options
	categories, err := h.dictClient.ListCategories(ctx)
	if err != nil {
		log.Error("error fetching categories", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	types, err := h.dictClient.ListTypesByCategory(ctx, uuid.Nil)
	if err != nil {
		log.Error("error fetching types", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	subtypes, err := h.dictClient.ListSubtypesByType(ctx, uuid.Nil)
	if err != nil {
		log.Error("error fetching subtypes", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := h.tmplMgr.Get("edit-property.html")
	if err != nil {
		log.Error("error getting template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":      fmt.Sprintf("Edit: %s", property.Name),
		"Property":   property,
		"ActiveNav":  "properties",
		"Template":   "edit-property",
		"Categories": DictionaryOptionsToMap(categories),
		"Types":      DictionaryOptionsToMap(types),
		"Subtypes":   DictionaryOptionsToMap(subtypes),
	}

	if err := tmpl.ExecuteTemplate(w, "edit-property.html", data); err != nil {
		log.Error("error executing template", "error", err)
	}
}

// UpdateProperty handles updating an existing property
func (h *Handler) UpdateProperty(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.UpdateProperty")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid property id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Error("error parsing form", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Parse classification IDs
	categoryID, _ := uuid.Parse(r.FormValue("category_id"))
	typeID, _ := uuid.Parse(r.FormValue("type_id"))
	subtypeIDStr := r.FormValue("subtype_id")
	var subtypeID uuid.UUID
	if subtypeIDStr != "" {
		subtypeID, _ = uuid.Parse(subtypeIDStr)
	}

	// Parse numeric fields
	totalArea, _ := strconv.ParseFloat(r.FormValue("total_area"), 64)
	bedrooms, _ := strconv.Atoi(r.FormValue("bedrooms"))
	bathrooms, _ := strconv.Atoi(r.FormValue("bathrooms"))
	parking, _ := strconv.Atoi(r.FormValue("parking"))
	amount, _ := strconv.ParseFloat(r.FormValue("price_amount"), 64)

	req := &UpdatePropertyRequest{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Classification: Classification{
			CategoryID: categoryID,
			TypeID:     typeID,
			SubtypeID:  subtypeID,
		},
		Location: Location{
			Address: Address{
				Street:     r.FormValue("street"),
				Number:     r.FormValue("number"),
				City:       r.FormValue("city"),
				State:      r.FormValue("state"),
				PostalCode: r.FormValue("postal_code"),
				Country:    r.FormValue("country"),
			},
			Region: r.FormValue("region"),
		},
		Features: Features{
			TotalArea: totalArea,
			Bedrooms:  bedrooms,
			Bathrooms: bathrooms,
			Parking:   parking,
		},
		Price: Price{
			Amount:   amount,
			Currency: r.FormValue("currency"),
			Type:     r.FormValue("price_type"),
		},
		Status:  r.FormValue("status"),
		OwnerID: r.FormValue("owner_id"),
	}

	_, err = h.service.UpdateProperty(ctx, id, req)
	if err != nil {
		log.Error("error updating property", "error", err, "id", id)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Info("property updated successfully", "id", id)
	http.Redirect(w, r, fmt.Sprintf("/show-property/%s", id), http.StatusSeeOther)
}

// DeleteProperty handles deleting a property
func (h *Handler) DeleteProperty(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.http.Start(w, r, "Handler.DeleteProperty")
	defer finish()
	log := h.log(r)

	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("invalid property id", "id", idStr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteProperty(ctx, id); err != nil {
		log.Error("error deleting property", "error", err, "id", id)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Info("property deleted successfully", "id", id)
	http.Redirect(w, r, "/list-properties", http.StatusSeeOther)
}
