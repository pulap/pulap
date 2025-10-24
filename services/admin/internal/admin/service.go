package admin

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/services/admin/internal/config"
)

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*User, error)
	ListUsers(ctx context.Context) ([]*User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, req *UpdateUserRequest) (*User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	ListUsersByStatus(ctx context.Context, status string) ([]*User, error)

	CreateRole(ctx context.Context, req *CreateRoleRequest) (*Role, error)
	GetRole(ctx context.Context, id uuid.UUID) (*Role, error)
	ListRoles(ctx context.Context) ([]*Role, error)
	UpdateRole(ctx context.Context, id uuid.UUID, req *UpdateRoleRequest) (*Role, error)
	DeleteRole(ctx context.Context, id uuid.UUID) error

	CreateGrant(ctx context.Context, req *CreateGrantRequest) (*Grant, error)
	GetGrant(ctx context.Context, id uuid.UUID) (*Grant, error)
	ListGrants(ctx context.Context) ([]*Grant, error)
	DeleteGrant(ctx context.Context, id uuid.UUID) error
}

type defaultService struct {
	repos   Repos
	xparams config.XParams
}

type Repos struct {
	UserRepo  UserRepo
	RoleRepo  RoleRepo
	GrantRepo GrantRepo
}

//authzHelper := auth.NewAuthzHelper(authzHTTPClient, 5*time.Minute)

func NewDefaultService(repos Repos, xparams config.XParams) *defaultService {
	return &defaultService{
		repos:   repos,
		xparams: xparams,
	}
}

func (s *defaultService) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {

	return s.repos.UserRepo.Create(ctx, req)
}

func (s *defaultService) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
	return s.repos.UserRepo.Get(ctx, id)
}

func (s *defaultService) ListUsers(ctx context.Context) ([]*User, error) {
	return s.repos.UserRepo.List(ctx)
}

func (s *defaultService) UpdateUser(ctx context.Context, id uuid.UUID, req *UpdateUserRequest) (*User, error) {
	return s.repos.UserRepo.Update(ctx, id, req)
}

func (s *defaultService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repos.UserRepo.Delete(ctx, id)
}

func (s *defaultService) ListUsersByStatus(ctx context.Context, status string) ([]*User, error) {
	return s.repos.UserRepo.ListByStatus(ctx, status)
}

func (s *defaultService) CreateRole(ctx context.Context, req *CreateRoleRequest) (*Role, error) {
	return s.repos.RoleRepo.Create(ctx, req)
}

func (s *defaultService) GetRole(ctx context.Context, id uuid.UUID) (*Role, error) {
	return s.repos.RoleRepo.Get(ctx, id)
}

func (s *defaultService) ListRoles(ctx context.Context) ([]*Role, error) {
	return s.repos.RoleRepo.List(ctx)
}

func (s *defaultService) UpdateRole(ctx context.Context, id uuid.UUID, req *UpdateRoleRequest) (*Role, error) {
	return s.repos.RoleRepo.Update(ctx, id, req)
}

func (s *defaultService) DeleteRole(ctx context.Context, id uuid.UUID) error {
	return s.repos.RoleRepo.Delete(ctx, id)
}

func (s *defaultService) CreateGrant(ctx context.Context, req *CreateGrantRequest) (*Grant, error) {
	// Validar usuario
	user, err := s.repos.UserRepo.Get(ctx, req.UserID)
	if err != nil || user == nil {
		return nil, fmt.Errorf("user not found")
	}
	// Validar rol
	if strings.EqualFold(req.GrantType, "role") {
		roleID, err := uuid.Parse(req.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid role id")
		}

		role, err := s.repos.RoleRepo.Get(ctx, roleID)
		if err != nil || role == nil {
			return nil, fmt.Errorf("role not found")
		}
	}
	return s.repos.GrantRepo.Create(ctx, req)
}

func (s *defaultService) GetGrant(ctx context.Context, id uuid.UUID) (*Grant, error) {
	return s.repos.GrantRepo.Get(ctx, id)
}

func (s *defaultService) ListGrants(ctx context.Context) ([]*Grant, error) {
	return s.repos.GrantRepo.List(ctx)
}

func (s *defaultService) DeleteGrant(ctx context.Context, id uuid.UUID) error {
	return s.repos.GrantRepo.Delete(ctx, id)
}

func (s *defaultService) Log() core.Logger {
	return s.xparams.Log()
}

func (s *defaultService) Cfg() *config.Config {
	return s.xparams.Cfg()
}

func (s *defaultService) Trace() core.Tracer {
	return s.xparams.Tracer()
}
