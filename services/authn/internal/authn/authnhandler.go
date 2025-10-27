package authn

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/pkg/lib/telemetry"
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
		tlm: telemetry.NewHTTP(
			telemetry.WithTracer(xparams.Tracer()),
			telemetry.WithMetrics(xparams.Metrics()),
		),
	}
}

type AuthHandler struct {
	repo    UserRepo
	xparams config.XParams
	tlm     *telemetry.HTTP
}

func (h *AuthHandler) RegisterRoutes(r chi.Router) {
	r.Route("/authn", func(r chi.Router) {
		r.Post("/signup", h.SignUp)
		r.Post("/signin", h.SignIn)
		r.Post("/signout", h.SignOut)
	})
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "AuthHandler.SignUp")
	defer finish()

	log := h.log(r)
	ctx := r.Context()

	req, ok := h.decodeSignUpPayload(w, r, log)
	if !ok {
		return
	}

	validationErrors := ValidateSignUpRequest(req.Email, req.Password)
	if len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	log.Info("encryption key configured", "length", len([]byte(h.cfg().Auth.EncryptionKey)))
	log.Info("signing key configured", "length", len([]byte(h.cfg().Auth.SigningKey)))

	user, err := SignUpUser(ctx, h.repo, h.cfg(), req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, ErrUserExists):
			log.Debug("user already exists")
			core.RespondError(w, http.StatusConflict, "User already exists")
		default:
			log.Error("cannot create user", "error", err)
			core.RespondError(w, http.StatusInternalServerError, "Could not create account")
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	core.RespondSuccess(w, AuthResponse{User: user})
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "AuthHandler.SignIn")
	defer finish()

	log := h.log(r)
	ctx := r.Context()

	req, ok := h.decodeSignInPayload(w, r, log)
	if !ok {
		return
	}

	validationErrors := ValidateSignInRequest(req.Email, req.Password)
	if len(validationErrors) > 0 {
		log.Debug("validation failed", "errors", validationErrors)
		core.RespondError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	user, token, err := SignInUser(ctx, h.repo, h.cfg(), req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidCredentials):
			log.Debug("invalid credentials")
			core.RespondError(w, http.StatusUnauthorized, "Invalid credentials")
		case errors.Is(err, ErrInactiveAccount):
			log.Debug("user not active")
			core.RespondError(w, http.StatusForbidden, "Account is not active")
		default:
			log.Error("error signing in", "error", err)
			core.RespondError(w, http.StatusInternalServerError, "Authentication failed")
		}
		return
	}

	core.RespondSuccess(w, AuthResponse{User: user, Token: token})
}

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	w, r, finish := h.tlm.Start(w, r, "AuthHandler.SignOut")
	defer finish()

	log := h.log(r)

	// TODO: Invalidate session token

	log.Debug("user signed out")
	w.WriteHeader(http.StatusNoContent)
}

// Helper methods
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

func (h *AuthHandler) log(req ...*http.Request) core.Logger {
	logger := h.xparams.Log()
	if len(req) > 0 && req[0] != nil {
		r := req[0]
		return logger.With(
			"request_id", core.RequestIDFrom(r.Context()),
			"method", r.Method,
			"path", r.URL.Path,
		)
	}
	return logger
}

func (h *AuthHandler) cfg() *config.Config { return h.xparams.Cfg() }

func (h *AuthHandler) trace() core.Tracer { return h.xparams.Tracer() }
