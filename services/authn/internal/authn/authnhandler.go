package authn

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	authpkg "github.com/pulap/pulap/pkg/lib/auth"
	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/services/authn/internal/config"
)

const AuthMaxBodyBytes = 1 << 20

// SignUpRequest represents the signup payload
type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignInRequest represents the signin payload
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents successful authentication response
type AuthResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token,omitempty"`
}

// NewAuthHandler creates a new AuthHandler for authentication operations.
func NewAuthHandler(repo UserRepo, xparams config.XParams) *AuthHandler {
	return &AuthHandler{
		repo:    repo,
		xparams: xparams,
	}
}

type AuthHandler struct {
	repo    UserRepo
	xparams config.XParams
}

func (h *AuthHandler) RegisterRoutes(r chi.Router) {
	r.Route("/authn", func(r chi.Router) {
		r.Post("/signup", h.SignUp)
		r.Post("/signin", h.SignIn)
		r.Post("/signout", h.SignOut)
	})
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	req, ok := h.decodeSignUpPayload(w, r, log)
	if !ok {
		return
	}

	// Validate signup request
	validationErrors := ValidateSignUpRequest(req.Email, req.Password)
	if len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Create domain user using pure functions
	normalizedEmail := authpkg.NormalizeEmail(req.Email)

	// Use encryption keys from config
	encryptionKey := []byte(h.xparams.Cfg.Auth.EncryptionKey)
	signingKey := []byte(h.xparams.Cfg.Auth.SigningKey)

	// Debug: Check key lengths
	log.Info("encryption key configured", "length", len(encryptionKey))
	log.Info("signing key configured", "length", len(signingKey))

	// Encrypt email for storage
	encryptedEmail, err := authpkg.EncryptEmail(normalizedEmail, encryptionKey)
	if err != nil {
		log.Error("error encrypting email", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not create account")
		return
	}

	emailLookup := authpkg.ComputeLookupHash(normalizedEmail, signingKey)

	// Check if user already exists
	existingUser, err := h.repo.GetByEmailLookup(ctx, emailLookup)
	if err != nil {
		log.Error("error checking existing user", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not create account")
		return
	}
	if existingUser != nil {
		log.Debug("user already exists")
		core.RespondError(w, http.StatusConflict, "User already exists")
		return
	}

	// Generate salt and hash password
	salt := authpkg.GeneratePasswordSalt()
	passwordHash := authpkg.HashPassword([]byte(req.Password), salt)

	// TODO: Encrypt email (needs AES-GCM implementation in authpkg)
	// For now, store plaintext (will be encrypted once crypto functions are complete)

	// Create service user
	user := NewUser()
	user.EmailCT = encryptedEmail.Ciphertext
	user.EmailIV = encryptedEmail.IV
	user.EmailTag = encryptedEmail.Tag
	user.EmailLookup = emailLookup
	user.PasswordHash = passwordHash
	user.PasswordSalt = salt
	user.BeforeCreate()

	if err := h.repo.Create(ctx, user); err != nil {
		log.Error("cannot create user", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Could not create account")
		return
	}

	// Return success (no token in signup, user needs to signin)
	w.WriteHeader(http.StatusCreated)
	core.RespondSuccess(w, AuthResponse{User: user})
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)
	ctx := r.Context()

	req, ok := h.decodeSignInPayload(w, r, log)
	if !ok {
		return
	}

	// Validate signin request
	validationErrors := ValidateSignInRequest(req.Email, req.Password)
	if len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Normalize email and compute lookup hash
	normalizedEmail := authpkg.NormalizeEmail(req.Email)
	signingKey := []byte(h.xparams.Cfg.Auth.SigningKey)
	emailLookup := authpkg.ComputeLookupHash(normalizedEmail, signingKey)

	// Find user by email lookup
	user, err := h.repo.GetByEmailLookup(ctx, emailLookup)
	if err != nil {
		log.Error("error finding user", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Authentication failed")
		return
	}
	if user == nil {
		log.Debug("user not found")
		core.RespondError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Verify password using pure function
	if !authpkg.VerifyPasswordHash([]byte(req.Password), user.PasswordHash, user.PasswordSalt) {
		log.Debug("invalid password")
		core.RespondError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Check user status
	if user.Status != authpkg.UserStatusActive {
		log.Debug("user not active", "status", user.Status)
		core.RespondError(w, http.StatusForbidden, "Account is not active")
		return
	}

	// Generate session token
	token, err := h.generateSessionToken(user.ID.String())
	if err != nil {
		log.Error("error generating session token", "error", err)
		core.RespondError(w, http.StatusInternalServerError, "Authentication failed")
		return
	}

	core.RespondSuccess(w, AuthResponse{User: user, Token: token})
}

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	log := h.logForRequest(r)

	// TODO: Invalidate session token

	log.Debug("user signed out")
	w.WriteHeader(http.StatusNoContent)
}

// Helper methods

func (h *AuthHandler) logForRequest(r *http.Request) core.Logger {
	return h.xparams.Log.With(
		"request_id", core.RequestIDFrom(r.Context()),
		"method", r.Method,
		"path", r.URL.Path,
	)
}

func (h *AuthHandler) decodeSignUpPayload(w http.ResponseWriter, r *http.Request, log core.Logger) (SignUpRequest, bool) {
	var req SignUpRequest

	r.Body = http.MaxBytesReader(w, r.Body, AuthMaxBodyBytes)
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Debug("cannot read request body", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Could not read request body")
		return req, false
	}

	if len(strings.TrimSpace(string(body))) == 0 {
		log.Debug("empty request body")
		core.RespondError(w, http.StatusBadRequest, "Request body is empty")
		return req, false
	}

	if err := json.Unmarshal(body, &req); err != nil {
		log.Debug("cannot decode JSON", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Could not parse JSON")
		return req, false
	}

	return req, true
}

// generateSessionToken creates a session token for the user
func (h *AuthHandler) generateSessionToken(userID string) (string, error) {
	// Parse session TTL
	ttl, err := time.ParseDuration(h.xparams.Cfg.Auth.SessionTTL)
	if err != nil {
		return "", fmt.Errorf("invalid session TTL: %w", err)
	}

	// Get or generate Ed25519 private key
	privateKey, err := h.getTokenPrivateKey()
	if err != nil {
		return "", fmt.Errorf("could not get private key: %w", err)
	}

	// Generate session ID
	sessionID := uuid.New().String()

	// Generate token
	return authpkg.GenerateSessionToken(userID, sessionID, privateKey, ttl)
}

// getTokenPrivateKey gets or generates the Ed25519 private key for tokens
func (h *AuthHandler) getTokenPrivateKey() (ed25519.PrivateKey, error) {
	// Try to get from config first
	if h.xparams.Cfg.Auth.TokenPrivateKey != "" {
		keyBytes, err := base64.StdEncoding.DecodeString(h.xparams.Cfg.Auth.TokenPrivateKey)
		if err != nil {
			return nil, fmt.Errorf("error decode private key: %w", err)
		}
		return ed25519.PrivateKey(keyBytes), nil
	}

	// Generate new key pair if not configured
	// In production, this should be persistent
	_, privateKey, err := authpkg.GenerateKeyPair()
	if err != nil {
		return nil, fmt.Errorf("error generate key pair: %w", err)
	}

	return privateKey, nil
}

func (h *AuthHandler) decodeSignInPayload(w http.ResponseWriter, r *http.Request, log core.Logger) (SignInRequest, bool) {
	var req SignInRequest

	r.Body = http.MaxBytesReader(w, r.Body, AuthMaxBodyBytes)
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Debug("cannot read request body", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Could not read request body")
		return req, false
	}

	if len(strings.TrimSpace(string(body))) == 0 {
		log.Debug("empty request body")
		core.RespondError(w, http.StatusBadRequest, "Request body is empty")
		return req, false
	}

	if err := json.Unmarshal(body, &req); err != nil {
		log.Debug("cannot decode JSON", "error", err)
		core.RespondError(w, http.StatusBadRequest, "Could not parse JSON")
		return req, false
	}

	return req, true
}
