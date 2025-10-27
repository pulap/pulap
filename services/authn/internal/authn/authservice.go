package authn

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	authpkg "github.com/pulap/pulap/pkg/lib/auth"
	"github.com/pulap/pulap/services/authn/internal/config"
)

var (
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInactiveAccount    = errors.New("account is not active")
)

func SignUpUser(ctx context.Context, repo UserRepo, cfg *config.Config, email, password string) (*User, error) {
	if repo == nil {
		return nil, errors.New("user repository is required")
	}
	if cfg == nil {
		return nil, errors.New("configuration is required")
	}

	normalizedEmail := authpkg.NormalizeEmail(email)
	encryptionKey := []byte(cfg.Auth.EncryptionKey)
	signingKey := []byte(cfg.Auth.SigningKey)

	encryptedEmail, err := authpkg.EncryptEmail(normalizedEmail, encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("encrypt email: %w", err)
	}

	emailLookup := authpkg.ComputeLookupHash(normalizedEmail, signingKey)

	existing, err := repo.GetByEmailLookup(ctx, emailLookup)
	if err != nil {
		return nil, fmt.Errorf("lookup user: %w", err)
	}
	if existing != nil {
		return nil, ErrUserExists
	}

	salt := authpkg.GeneratePasswordSalt()
	passwordHash := authpkg.HashPassword([]byte(password), salt)

	user := NewUser()
	user.EmailCT = encryptedEmail.Ciphertext
	user.EmailIV = encryptedEmail.IV
	user.EmailTag = encryptedEmail.Tag
	user.EmailLookup = emailLookup
	user.PasswordHash = passwordHash
	user.PasswordSalt = salt
	user.BeforeCreate()

	if err := repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

func SignInUser(ctx context.Context, repo UserRepo, cfg *config.Config, email, password string) (*User, string, error) {
	if repo == nil {
		return nil, "", errors.New("user repository is required")
	}
	if cfg == nil {
		return nil, "", errors.New("configuration is required")
	}

	normalizedEmail := authpkg.NormalizeEmail(email)
	signingKey := []byte(cfg.Auth.SigningKey)
	emailLookup := authpkg.ComputeLookupHash(normalizedEmail, signingKey)

	user, err := repo.GetByEmailLookup(ctx, emailLookup)
	if err != nil {
		return nil, "", fmt.Errorf("lookup user: %w", err)
	}
	if user == nil {
		return nil, "", ErrInvalidCredentials
	}

	if !authpkg.VerifyPasswordHash([]byte(password), user.PasswordHash, user.PasswordSalt) {
		return nil, "", ErrInvalidCredentials
	}

	if user.Status != authpkg.UserStatusActive {
		return nil, "", ErrInactiveAccount
	}

	token, err := generateSessionToken(cfg, user.ID)
	if err != nil {
		return nil, "", fmt.Errorf("generate session token: %w", err)
	}

	return user, token, nil
}

func GenerateBootstrapStatus(ctx context.Context, repo UserRepo, cfg *config.Config) (*User, error) {
	if repo == nil {
		return nil, errors.New("user repository is required")
	}
	if cfg == nil {
		return nil, errors.New("configuration is required")
	}

	signingKey := []byte(cfg.Auth.SigningKey)
	normalizedEmail := authpkg.NormalizeEmail(SuperadminEmail)
	lookupHash := authpkg.ComputeLookupHash(normalizedEmail, signingKey)

	return repo.GetByEmailLookup(ctx, lookupHash)
}

func BootstrapSuperadmin(ctx context.Context, repo UserRepo, cfg *config.Config) (*User, string, error) {
	if repo == nil {
		return nil, "", errors.New("user repository is required")
	}
	if cfg == nil {
		return nil, "", errors.New("configuration is required")
	}

	signingKey := []byte(cfg.Auth.SigningKey)
	encryptionKey := []byte(cfg.Auth.EncryptionKey)

	normalizedEmail := authpkg.NormalizeEmail(SuperadminEmail)
	lookupHash := authpkg.ComputeLookupHash(normalizedEmail, signingKey)

	existing, err := repo.GetByEmailLookup(ctx, lookupHash)
	if err != nil {
		return nil, "", fmt.Errorf("lookup superadmin: %w", err)
	}
	if existing != nil {
		return existing, "", nil
	}

	encryptedEmail, err := authpkg.EncryptEmail(normalizedEmail, encryptionKey)
	if err != nil {
		return nil, "", fmt.Errorf("encrypt email: %w", err)
	}

	password := generateSecurePassword(32)
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

	if err := repo.Create(ctx, user); err != nil {
		return nil, "", fmt.Errorf("create superadmin: %w", err)
	}

	return user, password, nil
}

func generateSessionToken(cfg *config.Config, userID uuid.UUID) (string, error) {
	ttl, err := time.ParseDuration(cfg.Auth.SessionTTL)
	if err != nil {
		return "", fmt.Errorf("invalid session TTL: %w", err)
	}

	privateKey, err := tokenPrivateKey(cfg.Auth.TokenPrivateKey)
	if err != nil {
		return "", fmt.Errorf("get private key: %w", err)
	}

	sessionID := uuid.New().String()

	return authpkg.GenerateSessionToken(userID.String(), sessionID, privateKey, ttl)
}

func tokenPrivateKey(encoded string) (ed25519.PrivateKey, error) {
	if encoded != "" {
		keyBytes, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			return nil, fmt.Errorf("decode private key: %w", err)
		}
		return ed25519.PrivateKey(keyBytes), nil
	}

	_, privateKey, err := authpkg.GenerateKeyPair()
	if err != nil {
		return nil, fmt.Errorf("generate key pair: %w", err)
	}
	return privateKey, nil
}
