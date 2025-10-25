package authz

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	authpkg "github.com/pulap/pulap/pkg/lib/auth"
	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/services/authz/internal/config"
)

// BootstrapService handles the coordination of system bootstrap
type BootstrapService struct {
	roleRepo   RoleRepo
	grantRepo  GrantRepo
	httpClient *http.Client
	xparams    config.XParams
}

// BootstrapStatusResponse matches AuthN response
type BootstrapStatusResponse struct {
	NeedsBootstrap bool   `json:"needs_bootstrap"`
	SuperadminID   string `json:"superadmin_id,omitempty"`
}

// BootstrapResponse matches AuthN response
type BootstrapResponse struct {
	SuperadminID string `json:"superadmin_id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

func NewBootstrapService(roleRepo RoleRepo, grantRepo GrantRepo, xparams config.XParams) *BootstrapService {
	return &BootstrapService{
		roleRepo:  roleRepo,
		grantRepo: grantRepo,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		xparams: xparams,
	}
}

// Bootstrap orchestrates the complete bootstrap process
func (s *BootstrapService) Bootstrap(ctx context.Context) error {
	s.Log().Info("Starting bootstrap process...")

	status, err := s.getBootstrapStatus(ctx)
	if err != nil {
		return fmt.Errorf("failed to get bootstrap status: %w", err)
	}

	var superadminID string

	if status.NeedsBootstrap {
		s.Log().Info("System needs bootstrap, triggering AuthN bootstrap...")

		response, err := s.triggerBootstrap(ctx)
		if err != nil {
			return fmt.Errorf("failed to trigger bootstrap: %w", err)
		}

		superadminID = response.SuperadminID
		s.Log().Info("Bootstrap triggered successfully",
			"superadmin_id", response.SuperadminID,
			"email", response.Email,
			"password", response.Password) // log credentials for initial setup
	} else {
		s.Log().Info("System already bootstrapped", "superadmin_id", status.SuperadminID)
		superadminID = status.SuperadminID
	}

	if err := s.bootstrapRolesAndGrants(ctx, superadminID); err != nil {
		return fmt.Errorf("failed to bootstrap roles and grants: %w", err)
	}

	s.Log().Info("Bootstrap process completed successfully")
	return nil
}

// getBootstrapStatus calls AuthN to check bootstrap status
func (s *BootstrapService) getBootstrapStatus(ctx context.Context) (*BootstrapStatusResponse, error) {
	authNURL := s.Cfg().Auth.AuthNURL
	url := authNURL + "/system/bootstrap-status"
	s.Log().Info("AuthN URL: " + url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bootstrap status request failed: %d", resp.StatusCode)
	}

	// Parse wrapped response
	var wrapped struct {
		Data BootstrapStatusResponse `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapped); err != nil {
		return nil, err
	}

	return &wrapped.Data, nil
}

// triggerBootstrap calls AuthN to create superadmin
func (s *BootstrapService) triggerBootstrap(ctx context.Context) (*BootstrapResponse, error) {
	authNURL := s.Cfg().Auth.AuthNURL
	url := authNURL + "/system/bootstrap"

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bootstrap request failed: %d", resp.StatusCode)
	}

	var response BootstrapResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// bootstrapRolesAndGrants seeds roles and creates superadmin grant
func (s *BootstrapService) bootstrapRolesAndGrants(ctx context.Context, superadminID string) error {
	// Step 1: Seed base roles (idempotent)
	if err := s.seedRoles(ctx); err != nil {
		return fmt.Errorf("failed to seed roles: %w", err)
	}

	// Step 2: Ensure superadmin grant exists (idempotent)
	if err := s.ensureSuperadminGrant(ctx, superadminID); err != nil {
		return fmt.Errorf("failed to ensure superadmin grant: %w", err)
	}

	return nil
}

// seedRoles creates base system roles
func (s *BootstrapService) seedRoles(ctx context.Context) error {
	baseRoles := []struct {
		Name        string
		Description string
		Permissions []string
	}{
		{
			Name:        "superadmin",
			Description: "System superadmin with all permissions",
			Permissions: []string{
				"*:*", // Wildcard - can do everything
			},
		},
		{
			Name:        "admin",
			Description: "System administrator",
			Permissions: []string{
				"users:create",
				"users:read",
				"users:update",
				"users:delete",
				"users:list",
				"roles:create",
				"roles:read",
				"roles:update",
				"roles:delete",
				"roles:list",
				"grants:create",
				"grants:read",
				"grants:delete",
				"grants:list",
			},
		},
		{
			Name:        "user",
			Description: "Regular user",
			Permissions: []string{
				"users:read",   // Can read own profile
				"users:update", // Can update own profile
			},
		},
	}

	for _, roleData := range baseRoles {
		// Check if role exists (idempotent)
		existing, err := s.roleRepo.GetByName(ctx, roleData.Name)
		if err == nil && existing != nil {
			s.Log().Info("Role already exists, skipping", "name", roleData.Name)
			continue
		}

		role := &Role{
			ID:          uuid.New(),
			Name:        roleData.Name,
			Permissions: roleData.Permissions,
			Status:      authpkg.UserStatusActive,
			CreatedAt:   time.Now(),
			CreatedBy:   "system",
			UpdatedAt:   time.Now(),
			UpdatedBy:   "system",
		}

		if err := s.roleRepo.Create(ctx, role); err != nil {
			return fmt.Errorf("failed to create role %s: %w", roleData.Name, err)
		}

		s.Log().Info("Role created successfully", "name", roleData.Name, "id", role.ID)
	}

	return nil
}

// ensureSuperadminGrant creates grant for superadmin if it doesn't exist
func (s *BootstrapService) ensureSuperadminGrant(ctx context.Context, superadminID string) error {
	// Get superadmin role
	role, err := s.roleRepo.GetByName(ctx, "superadmin")
	if err != nil {
		return fmt.Errorf("superadmin role not found: %w", err)
	}

	// Parse superadmin UUID
	userID, err := uuid.Parse(superadminID)
	if err != nil {
		return fmt.Errorf("invalid superadmin ID: %w", err)
	}

	// Check if grant already exists (idempotent)
	grants, err := s.grantRepo.ListByUserID(ctx, userID)
	if err != nil {
		s.Log().Error("Failed to check existing grants, proceeding anyway", "error", err)
	} else {
		for _, g := range grants {
			if g.GrantType == GrantTypeRole && g.Value == role.ID.String() {
				s.Log().Info("Superadmin grant already exists", "grant_id", g.ID)
				return nil
			}
		}
	}

	// Create grant
	grant := &Grant{
		ID:        uuid.New(),
		UserID:    userID,
		GrantType: GrantTypeRole,
		Value:     role.ID.String(),
		Scope:     Scope{Type: "global", ID: ""},
		ExpiresAt: nil,
		Status:    authpkg.UserStatusActive,
		CreatedAt: time.Now(),
		CreatedBy: "system",
		UpdatedAt: time.Now(),
		UpdatedBy: "system",
	}

	if err := s.grantRepo.Create(ctx, grant); err != nil {
		return fmt.Errorf("failed to create superadmin grant: %w", err)
	}

	s.Log().Info("Superadmin grant created successfully",
		"grant_id", grant.ID,
		"user_id", userID,
		"role_id", role.ID)

	return nil
}

func (s *BootstrapService) Log() core.Logger {
	return s.xparams.Log()
}

func (s *BootstrapService) Cfg() *config.Config {
	return s.xparams.Cfg()
}

func (s *BootstrapService) Trace() core.Tracer {
	return s.xparams.Tracer()
}
