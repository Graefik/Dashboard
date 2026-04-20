package auth

import (
	"context"
	"errors"

	"back-dashboard/internal/domain/user"
	"back-dashboard/internal/ports/inbound"
	"back-dashboard/internal/ports/outbound"
)

type LoginService struct {
	users  outbound.UserRepository
	hasher outbound.PasswordHasher
	tokens outbound.TokenIssuer
}

func NewLoginService(
	users outbound.UserRepository,
	hasher outbound.PasswordHasher,
	tokens outbound.TokenIssuer,
) *LoginService {
	return &LoginService{users: users, hasher: hasher, tokens: tokens}
}

// Login vérifie les credentials et émet un token.
// En cas d'échec (user absent ou mauvais mdp), retourne toujours
// ErrInvalidCredentials : jamais indiquer lequel des deux est faux.
func (s *LoginService) Login(ctx context.Context, username, password string) (inbound.LoginResult, error) {
	u, err := s.users.FindByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return inbound.LoginResult{}, user.ErrInvalidCredentials
		}
		return inbound.LoginResult{}, err
	}

	if !s.hasher.Verify(password, u.PasswordHash) {
		return inbound.LoginResult{}, user.ErrInvalidCredentials
	}

	token, expiresAt, err := s.tokens.Issue(u.ID, u.Username)
	if err != nil {
		return inbound.LoginResult{}, err
	}

	return inbound.LoginResult{
		Token:     token,
		Username:  u.Username,
		ExpiresAt: expiresAt.Unix(),
	}, nil
}
