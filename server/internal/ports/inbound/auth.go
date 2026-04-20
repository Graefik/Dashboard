package inbound

import "context"

// AuthUseCase est exposé aux adapters driving (HTTP, CLI…).
type AuthUseCase interface {
	Login(ctx context.Context, username, password string) (LoginResult, error)
}

// LoginResult est le résultat d'un login réussi.
// Représentation purement métier : pas de token format ni de serialisation.
type LoginResult struct {
	Token     string
	Username  string
	ExpiresAt int64 // unix seconds
}
