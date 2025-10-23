package admin

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/services/admin/internal/config"

	authpkg "github.com/pulap/pulap/pkg/lib/auth"
)

type Handler struct {
	tmplMgr     *core.TemplateManager
	service     Service
	authZClt    *core.AuthzHTTPClient
	authnClient *core.ServiceClient
	xparams     config.XParams

	sessionValidator func(string) (string, error)
}

func NewHandler(
	tmplMgr *core.TemplateManager,
	service Service,
	autZClt *core.AuthzHTTPClient,
	authnClient *core.ServiceClient,
	xparams config.XParams,
) *Handler {
	return &Handler{
		tmplMgr:          tmplMgr,
		service:          service,
		authZClt:         autZClt,
		authnClient:      authnClient,
		xparams:          xparams,
		sessionValidator: defaultSessionValidator(),
	}
}

// RegisterRoutes registers all admin routes using Commands/Queries pattern
func (h *Handler) RegisterRoutes(r chi.Router) {
	h.Log().Info("Registering admin routes...")

	r.Get("/signin", h.ShowSignIn)
	r.Post("/signin", h.HandleSignIn)
	r.Get("/signup", h.ShowSignUp)
	r.Post("/signup", h.HandleSignUp)
	r.Post("/signout", h.HandleSignOut)

	r.Group(func(r chi.Router) {
		r.Use(SessionMiddleware(h.sessionValidator))

		h.Log().Info("Registering user management routes...")
		r.Get("/", h.Home)

		r.Get("/list-users", h.ListUsers)
		r.Get("/new-user", h.NewUser)
		r.Post("/create-user", h.CreateUser)
		r.Get("/show-user/{id}", h.ShowUser)
		r.Get("/edit-user/{id}", h.EditUser)
		r.Post("/update-user/{id}", h.UpdateUser)
		r.Post("/delete-user/{id}", h.DeleteUser)

		h.Log().Info("Registering role management routes...")
		r.Get("/list-roles", h.ListRoles)
		r.Get("/new-role", h.NewRole)
		r.Post("/create-role", h.CreateRole)
		r.Get("/show-role/{id}", h.ShowRole)
		r.Get("/edit-role/{id}", h.EditRole)
		r.Post("/update-role/{id}", h.UpdateRole)
		r.Post("/delete-role/{id}", h.DeleteRole)

		h.Log().Info("Registering grant management routes...")
		r.Get("/user-grants/{userId}", h.UserGrants)
		r.Post("/create-grant", h.CreateGrant)
		r.Post("/delete-grant/{id}", h.DeleteGrant)
	})

	h.Log().Info("Admin routes registered successfully")
}

func (h *Handler) ShowSignIn(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title":    "Sign In",
		"Next":     r.URL.Query().Get("next"),
		"Template": "signin",
		"HideNav":  true,
		"AuthPage": true,
	}

	h.renderTemplate(w, "signin.html", "base.html", data)
}

func (h *Handler) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	if h.authnClient == nil {
		h.Log().Error("authn client not configured")
		http.Error(w, "Authentication service unavailable", http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
		h.Log().Error("error parsing signin form", "error", err)
		h.renderSignInWithError(w, r, "Could not parse credentials")
		return
	}

	email := strings.TrimSpace(r.FormValue("email"))
	password := r.FormValue("password")
	nextURL := r.FormValue("next")

	if email == "" || password == "" {
		h.renderSignInWithError(w, r, "Email and password are required")
		return
	}

	payload := map[string]string{
		"email":    email,
		"password": password,
	}

	resp, err := h.authnClient.Request(r.Context(), http.MethodPost, "/authn/signin", payload)
	if err != nil {
		h.handleSignInError(w, r, err)
		return
	}

	if resp == nil {
		h.Log().Error("authn signin returned empty response")
		h.renderSignInWithError(w, r, "Authentication service error")
		return
	}

	userData, userID, err := extractUserID(resp.Data)
	if err != nil {
		h.Log().Error("cannot extract user from signin response", "error", err, "user_data", userData)
		http.Error(w, "Authentication service error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    userID,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	h.Log().Info("user signed in", "user_id", userID, "email", email)

	target := sanitizeRedirect(nextURL)
	http.Redirect(w, r, target, http.StatusFound)
}

func (h *Handler) HandleSignOut(w http.ResponseWriter, r *http.Request) {
	clearSessionCookie(w)
	h.Log().Debug("user signed out")
	http.Redirect(w, r, "/signin", http.StatusFound)
}

func (h *Handler) ShowSignUp(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title":    "Sign Up",
		"Template": "signup",
		"HideNav":  true,
		"AuthPage": true,
	}
	h.renderTemplate(w, "signup.html", "base.html", data)
}

func (h *Handler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Sign up is not available in this environment", http.StatusNotImplemented)
}

func (h *Handler) renderTemplate(w http.ResponseWriter, templateName, layout string, data map[string]interface{}) {
	tmpl, err := h.tmplMgr.Get(templateName)
	if err != nil {
		h.Log().Error("error loading template", "error", err, "template", templateName)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, layout, data); err != nil {
		h.Log().Error("error rendering template", "error", err, "layout", layout)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) renderSignInWithError(w http.ResponseWriter, r *http.Request, message string) {
	data := map[string]interface{}{
		"Title":    "Sign In",
		"Next":     r.FormValue("next"),
		"Error":    message,
		"Email":    r.FormValue("email"),
		"Template": "signin",
		"HideNav":  true,
		"AuthPage": true,
	}

	h.renderTemplate(w, "signin.html", "base.html", data)
}

func (h *Handler) handleSignInError(w http.ResponseWriter, r *http.Request, err error) {
	if httpErr, ok := err.(*core.HTTPError); ok {
		h.Log().Debug("authn signin failed", "status", httpErr.StatusCode, "message", httpErr.Message)

		switch httpErr.StatusCode {
		case http.StatusUnauthorized, http.StatusForbidden:
			h.renderSignInWithError(w, r, "Invalid credentials")
			return
		case http.StatusBadRequest:
			h.renderSignInWithError(w, r, "Invalid request")
			return
		default:
			h.renderSignInWithError(w, r, "Authentication service unavailable")
			return
		}
	}

	h.Log().Error("authn signin request failed", "error", err)
	h.renderSignInWithError(w, r, "Authentication service error")
}

func extractUserID(data interface{}) (map[string]interface{}, string, error) {
	payload, ok := data.(map[string]interface{})
	if !ok {
		return nil, "", fmt.Errorf("unexpected signin payload type %T", data)
	}

	userRaw, ok := payload["user"]
	if !ok {
		return payload, "", fmt.Errorf("signin payload missing user field")
	}

	userData, ok := userRaw.(map[string]interface{})
	if !ok {
		return payload, "", fmt.Errorf("unexpected user payload type %T", userRaw)
	}

	idRaw, ok := userData["id"]
	if !ok {
		return userData, "", fmt.Errorf("signin user payload missing id")
	}

	idStr, ok := idRaw.(string)
	if !ok {
		return userData, "", fmt.Errorf("signin user id is not string")
	}

	return userData, idStr, nil
}

func sanitizeRedirect(target string) string {
	if target == "" {
		return "/"
	}

	if !strings.HasPrefix(target, "/") || strings.HasPrefix(target, "//") {
		return "/"
	}

	return target
}

// Home renders the admin dashboard
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if code, err := h.pfc(r, "dashboard:read", "*"); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	tmpl, err := h.tmplMgr.Get("home.html")
	if err != nil {
		h.Log().Error("error getting home template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "Dashboard",
		"ActiveNav": "dashboard",
	}

	if err := tmpl.ExecuteTemplate(w, "home.html", data); err != nil {
		h.Log().Error("error executing home template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	if code, err := h.pfc(r, "user:list", "*"); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	tmpl, err := h.tmplMgr.Get("users.html")
	if err != nil {
		h.Log().Error("error getting users template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		h.Log().Error("error fetching users", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "Users",
		"ActiveNav": "users",
		"Template":  "users-content",
		"Users":     users,
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		h.Log().Error("error executing users template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) NewUser(w http.ResponseWriter, r *http.Request) {
	if code, err := h.pfc(r, "user:create", "*"); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	tmpl, err := h.tmplMgr.Get("new-user.html")
	if err != nil {
		h.Log().Error("error getting new-user template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "New User",
		"ActiveNav": "users",
		"Template":  "new-user",
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		h.Log().Error("error executing base template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if code, err := h.pfc(r, "user:create", "*"); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	if err := r.ParseForm(); err != nil {
		h.Log().Error("error parsing form", "error", err)
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

	user, err := h.service.CreateUser(r.Context(), req)
	if err != nil {
		h.Log().Error("error creating user", "error", err)
		http.Error(w, "Cannot create user: "+err.Error(), http.StatusBadRequest)
		return
	}

	h.Log().Info("User created successfully", "id", user.ID, "email", user.Email)

	w.Header().Set("HX-Redirect", "/list-users")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ShowUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	if code, err := h.pfc(r, "user:read", idStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		h.Log().Error("error fetching user", "error", err, "id", id)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	tmpl, err := h.tmplMgr.Get("show-user.html")
	if err != nil {
		h.Log().Error("error getting show-user template", "error", err)
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
		h.Log().Error("error executing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) EditUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	if code, err := h.pfc(r, "user:update", idStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		h.Log().Error("error fetching user", "error", err, "id", id)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	tmpl, err := h.tmplMgr.Get("edit-user.html")
	if err != nil {
		h.Log().Error("error getting edit-user template", "error", err)
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
		h.Log().Error("error executing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if code, err := h.pfc(r, "user:update", idStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

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

	user, err := h.service.UpdateUser(r.Context(), id, req)
	if err != nil {
		h.Log().Error("error updating user", "error", err)
		http.Error(w, "Cannot update user: "+err.Error(), http.StatusBadRequest)
		return
	}

	h.Log().Info("User updated successfully", "id", user.ID)
	w.Header().Set("HX-Redirect", "/list-users")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	if code, err := h.pfc(r, "user:delete", idStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.Log().Error("invalid user ID", "error", err, "id", idStr)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteUser(r.Context(), id); err != nil {
		h.Log().Error("error deleting user", "error", err, "id", id)
		http.Error(w, "Cannot delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListRoles(w http.ResponseWriter, r *http.Request) {
	if code, err := h.pfc(r, "role:list", "*"); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	tmpl, err := h.tmplMgr.Get("roles.html")
	if err != nil {
		h.Log().Error("error getting roles template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	roles, err := h.service.ListRoles(r.Context())
	if err != nil {
		h.Log().Error("error fetching roles", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "Roles",
		"ActiveNav": "roles",
		"Template":  "roles-content",
		"Roles":     roles,
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		h.Log().Error("error executing roles template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) NewRole(w http.ResponseWriter, r *http.Request) {
	if code, err := h.pfc(r, "role:create", "*"); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	tmpl, err := h.tmplMgr.Get("new-role.html")
	if err != nil {
		h.Log().Error("error getting new-role template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":              "New Role",
		"ActiveNav":          "roles",
		"Template":           "new-role",
		"PermissionRegistry": authpkg.PermissionRegistry,
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		h.Log().Error("error executing base template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) CreateRole(w http.ResponseWriter, r *http.Request) {
	if code, err := h.pfc(r, "role:create", "*"); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	if err := r.ParseForm(); err != nil {
		h.Log().Error("error parsing form", "error", err)
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	req := &CreateRoleRequest{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Permissions: r.Form["permissions"],
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	role, err := h.service.CreateRole(r.Context(), req)
	if err != nil {
		h.Log().Error("error creating role", "error", err)
		http.Error(w, "Cannot create role: "+err.Error(), http.StatusBadRequest)
		return
	}

	h.Log().Info("Role created successfully", "id", role.ID, "name", role.Name)

	w.Header().Set("HX-Redirect", "/list-roles")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ShowRole(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing role ID", http.StatusBadRequest)
		return
	}
	if code, err := h.pfc(r, "role:read", idStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}

	role, err := h.service.GetRole(r.Context(), id)
	if err != nil {
		h.Log().Error("error fetching role", "error", err, "id", id)
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}

	tmpl, err := h.tmplMgr.Get("show-role.html")
	if err != nil {
		h.Log().Error("error getting show-role template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":     "View Role",
		"ActiveNav": "roles",
		"Template":  "show-role",
		"Role":      role,
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		h.Log().Error("error executing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) EditRole(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing role ID", http.StatusBadRequest)
		return
	}
	if code, err := h.pfc(r, "role:update", idStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}

	role, err := h.service.GetRole(r.Context(), id)
	if err != nil {
		h.Log().Error("error fetching role", "error", err, "id", id)
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}

	tmpl, err := h.tmplMgr.Get("edit-role.html")
	if err != nil {
		h.Log().Error("error getting edit-role template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":              "Edit Role",
		"ActiveNav":          "roles",
		"Template":           "edit-role",
		"Role":               role,
		"PermissionRegistry": authpkg.PermissionRegistry,
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		h.Log().Error("error executing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if code, err := h.pfc(r, "role:update", idStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	req := &UpdateRoleRequest{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Permissions: r.Form["permissions"],
		Status:      r.FormValue("status"),
	}

	role, err := h.service.UpdateRole(r.Context(), id, req)
	if err != nil {
		h.Log().Error("error updating role", "error", err)
		http.Error(w, "Cannot update role: "+err.Error(), http.StatusBadRequest)
		return
	}

	h.Log().Info("Role updated successfully", "id", role.ID)
	w.Header().Set("HX-Redirect", "/list-roles")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing role ID", http.StatusBadRequest)
		return
	}
	if code, err := h.pfc(r, "role:delete", idStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.Log().Error("invalid role ID", "error", err, "id", idStr)
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteRole(r.Context(), id); err != nil {
		h.Log().Error("error deleting role", "error", err, "id", id)
		http.Error(w, "Cannot delete role", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UserGrants(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userId")
	if userIDStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	if code, err := h.pfc(r, "grant:list", userIDStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		h.Log().Error("error fetching user", "error", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	grants, err := h.service.ListGrants(r.Context())
	if err != nil {
		h.Log().Error("error fetching grants", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	roles, err := h.service.ListRoles(r.Context())
	if err != nil {
		h.Log().Error("error fetching roles", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := h.tmplMgr.Get("user-grants.html")
	if err != nil {
		h.Log().Error("error getting user-grants template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":              "User Grants",
		"ActiveNav":          "users",
		"Template":           "user-grants",
		"User":               user,
		"Grants":             grants,
		"Roles":              roles,
		"PermissionRegistry": authpkg.PermissionRegistry,
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		h.Log().Error("error executing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) CreateGrant(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.Log().Error("error parsing form", "error", err)
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	userIDStr := r.FormValue("user_id")
	if userIDStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	if code, err := h.pfc(r, "grant:create", userIDStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	req := &CreateGrantRequest{
		UserID:    userID,
		GrantType: r.FormValue("grant_type"),
		Value:     r.FormValue("value"),
		Scope: Scope{
			Type: r.FormValue("scope_type"),
			ID:   r.FormValue("scope_id"),
		},
	}

	grant, err := h.service.CreateGrant(r.Context(), req)
	if err != nil {
		h.Log().Error("error creating grant", "error", err)
		http.Error(w, "Cannot create grant: "+err.Error(), http.StatusBadRequest)
		return
	}

	h.Log().Info("Grant created successfully", "id", grant.ID, "user_id", grant.UserID)

	w.Header().Set("HX-Redirect", "/user-grants/"+userID.String())
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteGrant(w http.ResponseWriter, r *http.Request) {
	grantIDStr := chi.URLParam(r, "id")
	if grantIDStr == "" {
		http.Error(w, "Missing grant ID", http.StatusBadRequest)
		return
	}
	if code, err := h.pfc(r, "grant:delete", grantIDStr); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	id, err := uuid.Parse(grantIDStr)
	if err != nil {
		h.Log().Error("invalid grant ID", "error", err, "id", grantIDStr)
		http.Error(w, "Invalid grant ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteGrant(r.Context(), id); err != nil {
		h.Log().Error("error deleting grant", "error", err, "id", id)
		http.Error(w, "Cannot delete grant", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Log() core.Logger {
	return h.xparams.Log
}

// SetSessionValidator overrides the default session validation logic.
func (h *Handler) SetSessionValidator(validator func(string) (string, error)) {
	if validator == nil {
		return
	}
	h.sessionValidator = validator
}

func defaultSessionValidator() func(string) (string, error) {
	return func(sessionID string) (string, error) {
		if strings.TrimSpace(sessionID) == "" {
			return "", errors.New("invalid session")
		}
		// TODO: Integrate with AuthN service to resolve user ID from session token.
		return sessionID, nil
	}
}
