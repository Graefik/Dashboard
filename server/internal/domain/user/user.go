package user

import (
	"strings"
	"time"
)

// User est l'entité métier. Zéro dépendance technique (pas de gorm, pas de json).
type User struct {
	ID           uint
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}

// NewUser valide les invariants et construit une entité.
// Le hash doit être calculé par l'adapter security avant d'arriver ici.
func NewUser(username, passwordHash string) (*User, error) {
	u := strings.TrimSpace(username)
	if u == "" {
		return nil, ErrInvalidUsername
	}
	if len(u) > 100 {
		return nil, ErrInvalidUsername
	}
	if passwordHash == "" {
		return nil, ErrInvalidPasswordHash
	}
	return &User{
		Username:     u,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UTC(),
	}, nil
}
