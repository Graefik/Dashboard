package auth

import (
	"context"
	"fmt"
	"log"
	"strings"

	"back-dashboard/internal/domain/user"
	"back-dashboard/internal/ports/outbound"
)

type Bootstrap struct {
	users    outbound.UserRepository
	hasher   outbound.PasswordHasher
	passGen  outbound.PasswordGenerator
	username string
	length   int
}

func NewBootstrap(
	users outbound.UserRepository,
	hasher outbound.PasswordHasher,
	passGen outbound.PasswordGenerator,
	username string,
	length int,
) *Bootstrap {
	if username == "" {
		username = "admin"
	}
	if length <= 0 {
		length = 32
	}
	return &Bootstrap{
		users:    users,
		hasher:   hasher,
		passGen:  passGen,
		username: username,
		length:   length,
	}
}

// EnsureAdmin crée l'utilisateur admin au premier boot s'il n'existe pas,
// génère un mot de passe aléatoire, le hashe et le log EN CLAIR une seule fois
// dans stdout (style Jenkins).
func (b *Bootstrap) EnsureAdmin(ctx context.Context) error {
	count, err := b.users.CountAll(ctx)
	if err != nil {
		return fmt.Errorf("count users: %w", err)
	}
	if count > 0 {
		return nil
	}

	plain, err := b.passGen.Generate(b.length)
	if err != nil {
		return fmt.Errorf("generate password: %w", err)
	}

	hash, err := b.hasher.Hash(plain)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	u, err := user.NewUser(b.username, hash)
	if err != nil {
		return fmt.Errorf("build user: %w", err)
	}

	if err := b.users.Create(ctx, u); err != nil {
		return fmt.Errorf("persist admin: %w", err)
	}

	printInitialPassword(b.username, plain)
	return nil
}

func printInitialPassword(username, plain string) {
	banner := strings.Repeat("*", 72)
	log.Printf("\n\n%s\n%s\n%s\n\nGraefik — initial admin credentials\n\n  username: %s\n  password: %s\n\nCe mot de passe est affiché UNE SEULE FOIS.\nNotez-le maintenant, il ne sera jamais ré-affiché.\n\n%s\n%s\n%s\n\n",
		banner, banner, banner,
		username, plain,
		banner, banner, banner,
	)
}
