package outbound

import (
	"context"

	"back-dashboard/internal/domain/user"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*user.User, error)
	Create(ctx context.Context, u *user.User) error
	CountAll(ctx context.Context) (int64, error)
}
