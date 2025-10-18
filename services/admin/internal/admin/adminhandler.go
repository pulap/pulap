package admin

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/services/admin/internal/config"
)

// AdminHandler manages the web interface for system administration
type AdminHandler struct {
	xparams  config.XParams
	tmplMgr  *core.TemplateManager
	userRepo UserRepo
}

func NewAdminHandler(tmplMgr *core.TemplateManager, userRepo UserRepo, xparams config.XParams) *AdminHandler {
	return &AdminHandler{
		xparams:  xparams,
		tmplMgr:  tmplMgr,
		userRepo: userRepo,
	}
}

// RegisterRoutes registers all admin routes using Commands/Queries pattern
func (h *AdminHandler) RegisterRoutes(r chi.Router) {
	h.xparams.Log.Info("Registering admin routes...")

	r.Get("/", h.Home)
	r.Get("/health", h.Health)

	h.xparams.Log.Info("Registering user management routes...")
	r.Get("/list-users", h.ListUsers)
	r.Get("/new-user", h.NewUser)
	r.Post("/create-user", h.CreateUser)
	r.Get("/show-user/{id}", h.ShowUser)
	r.Get("/edit-user/{id}", h.EditUser)
	r.Post("/update-user/{id}", h.UpdateUser)
	r.Post("/delete-user/{id}", h.DeleteUser)
	h.xparams.Log.Info("Admin routes registered successfully")
}

// Health check endpoint
func (h *AdminHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Home renders the admin dashboard
func (h *AdminHandler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmplMgr.Get("home.html")
	if err != nil {
		h.xparams.Log.Error("error getting home template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "Dashboard",
		"ActiveNav": "dashboard",
	}

	if err := tmpl.ExecuteTemplate(w, "home.html", data); err != nil {
		h.xparams.Log.Error("error executing home template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *AdminHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmplMgr.Get("users.html")
	if err != nil {
		h.xparams.Log.Error("error getting users template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	users, err := h.userRepo.List(r.Context())
	if err != nil {
		h.xparams.Log.Error("error fetching users", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "Users",
		"ActiveNav": "users",
		"Users":     users,
	}

	if err := tmpl.ExecuteTemplate(w, "users.html", data); err != nil {
		h.xparams.Log.Error("error executing users template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *AdminHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmplMgr.Get("new-user.html")
	if err != nil {
		h.xparams.Log.Error("error getting new-user template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "New User",
		"ActiveNav": "users",
		"Template":  "new-user",
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		h.xparams.Log.Error("error executing base template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *AdminHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.xparams.Log.Error("error parsing form", "error", err)
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	req := &CreateUserRequest{
		Email:    r.FormValue("email"),
		Name:     r.FormValue("name"),
		Password: r.FormValue("password"),
	}

	if req.Email == "" || req.Name == "" || req.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.Create(r.Context(), req)
	if err != nil {
		h.xparams.Log.Error("error creating user", "error", err)
		http.Error(w, "Cannot create user: "+err.Error(), http.StatusBadRequest)
		return
	}

	h.xparams.Log.Info("User created successfully", "id", user.ID, "email", user.Email)

	w.Header().Set("HX-Redirect", "/list-users")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) ShowUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.Get(r.Context(), id)
	if err != nil {
		h.xparams.Log.Error("error fetching user", "error", err, "id", id)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	tmpl, err := h.tmplMgr.Get("show-user.html")
	if err != nil {
		h.xparams.Log.Error("error getting show-user template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "View User",
		"ActiveNav": "users",
		"Template":  "show-user",
		"User":      user,
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		h.xparams.Log.Error("error executing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *AdminHandler) EditUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.Get(r.Context(), id)
	if err != nil {
		h.xparams.Log.Error("error fetching user", "error", err, "id", id)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	tmpl, err := h.tmplMgr.Get("edit-user.html")
	if err != nil {
		h.xparams.Log.Error("error getting edit-user template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "Edit User",
		"ActiveNav": "users",
		"Template":  "edit-user",
		"User":      user,
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		h.xparams.Log.Error("error executing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *AdminHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	req := &UpdateUserRequest{
		Email:  r.FormValue("email"),
		Name:   r.FormValue("name"),
		Status: r.FormValue("status"),
	}

	user, err := h.userRepo.Update(r.Context(), id, req)
	if err != nil {
		h.xparams.Log.Error("error updating user", "error", err)
		http.Error(w, "Cannot update user: "+err.Error(), http.StatusBadRequest)
		return
	}

	h.xparams.Log.Info("User updated successfully", "id", user.ID)
	w.Header().Set("HX-Redirect", "/list-users")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.xparams.Log.Error("invalid user ID", "error", err, "id", idStr)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.userRepo.Delete(r.Context(), id); err != nil {
		h.xparams.Log.Error("error deleting user", "error", err, "id", id)
		http.Error(w, "Cannot delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
