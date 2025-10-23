package authn

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	authpkg "github.com/pulap/pulap/pkg/lib/auth"
	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/services/authn/internal/config"
)

// SystemHandler manages system-level operations like bootstrap
type SystemHandler struct {
	userRepo UserRepo
	xparams  config.XParams
}

// BootstrapStatusResponse represents the current bootstrap status
type BootstrapStatusResponse struct {
	NeedsBootstrap bool   `json:"needs_bootstrap"`
	SuperadminID   string `json:"superadmin_id,omitempty"` // Only if !needs_bootstrap
}

// BootstrapResponse represents the result of bootstrap operation
type BootstrapResponse struct {
	SuperadminID string `json:"superadmin_id"`
	Email        string `json:"email"`
	Password     string `json:"password"` // Generated password
}

const SuperadminEmail = "superadmin@system"

func NewSystemHandler(userRepo UserRepo, xparams config.XParams) *SystemHandler {
	return &SystemHandler{
		userRepo: userRepo,
		xparams:  xparams,
	}
}

// RegisterRoutes registers system management routes
func (h *SystemHandler) RegisterRoutes(r chi.Router) {
	h.Log().Info("Registering system routes...")

	r.Get("/system/bootstrap-status", h.GetBootstrapStatus)
	r.Post("/system/bootstrap", h.Bootstrap)

	h.Log().Info("System routes registered successfully")
}

// GetBootstrapStatus checks if the system needs bootstrap
func (h *SystemHandler) GetBootstrapStatus(w http.ResponseWriter, r *http.Request) {
	signingKey := []byte(h.xparams.Cfg.Auth.SigningKey)
	normalizedEmail := authpkg.NormalizeEmail(SuperadminEmail)
	lookupHash := authpkg.ComputeLookupHash(normalizedEmail, signingKey)

	superadmin, err := h.userRepo.GetByEmailLookup(r.Context(), lookupHash)
	if err != nil || superadmin == nil {
		if err != nil {
			h.Log().Error("failed to check superadmin user", "error", err)
		}
		response := BootstrapStatusResponse{
			NeedsBootstrap: true,
		}
		core.RespondSuccess(w, response)
		return
	}

	response := BootstrapStatusResponse{
		NeedsBootstrap: false,
		SuperadminID:   superadmin.ID.String(),
	}
	core.RespondSuccess(w, response)
}

// Bootstrap creates the superadmin user if it doesn't exist
func (h *SystemHandler) Bootstrap(w http.ResponseWriter, r *http.Request) {
	encKey := []byte(h.xparams.Cfg.Auth.EncryptionKey)
	signingKey := []byte(h.xparams.Cfg.Auth.SigningKey)
	normalizedEmail := authpkg.NormalizeEmail(SuperadminEmail)
	lookupHash := authpkg.ComputeLookupHash(normalizedEmail, signingKey)

	existing, err := h.userRepo.GetByEmailLookup(r.Context(), lookupHash)
	if err == nil && existing != nil {
		response := BootstrapResponse{
			SuperadminID: existing.ID.String(),
			Email:        SuperadminEmail,
			Password:     "",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	if err != nil {
		h.Log().Error("failed to check existing superadmin", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Failed to check bootstrap state")
		return
	}

	password := generateSecurePassword(32)

	encryptedEmail, err := authpkg.EncryptEmail(normalizedEmail, encKey)
	if err != nil {
		h.Log().Error("failed to encrypt email", "error", err)
		http.Error(w, "Failed to encrypt email", http.StatusInternalServerError)
		return
	}

	passwordSalt := authpkg.GeneratePasswordSalt()
	passwordHash := authpkg.HashPassword([]byte(password), passwordSalt)

	user := &User{
		ID:           uuid.New(),
		EmailCT:      encryptedEmail.Ciphertext,
		EmailIV:      encryptedEmail.IV,
		EmailTag:     encryptedEmail.Tag,
		EmailLookup:  lookupHash,
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
		Status:       authpkg.UserStatusActive,
		CreatedAt:    time.Now(),
		CreatedBy:    "system",
		UpdatedAt:    time.Now(),
		UpdatedBy:    "system",
	}

	err = h.userRepo.Create(r.Context(), user)
	if err != nil {
		h.Log().Error("failed to create superadmin user", "error", err)
		http.Error(w, "Failed to create superadmin", http.StatusInternalServerError)
		return
	}

	bannerLines := []string{
		"═══════════════════════════════════════════════════════════",
		"  SUPERADMIN BOOTSTRAP CREDENTIALS",
		"═══════════════════════════════════════════════════════════",
		fmt.Sprintf("  Email:    %s", SuperadminEmail),
		fmt.Sprintf("  Password: %s", password),
		fmt.Sprintf("  UserID:   %s", user.ID.String()),
		"═══════════════════════════════════════════════════════════",
		"  IMPORTANT: Save these credentials securely!",
		"  TODO: Implement mandatory password change on first login",
		"═══════════════════════════════════════════════════════════",
	}

	for _, line := range bannerLines {
		h.Log().Info(line)
	}

	h.Log().Info("superadmin bootstrap credentials",
		"email", SuperadminEmail,
		"user_id", user.ID,
	)

	// TODO: Write to file (optional)
	// writeBootstrapFile(user.ID.String(), SuperadminEmail, password)

	h.Log().Info("superadmin created successfully", "id", user.ID)

	response := BootstrapResponse{
		SuperadminID: user.ID.String(),
		Email:        SuperadminEmail,
		Password:     password,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func generateSecurePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	b := make([]byte, length)

	for i := range b {
		randomByte := make([]byte, 1)
		rand.Read(randomByte)
		b[i] = charset[int(randomByte[0])%len(charset)]
	}

	return string(b)
}

func (h *SystemHandler) Log() core.Logger {
	return h.xparams.Log
}
