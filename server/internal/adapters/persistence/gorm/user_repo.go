package gormdb

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"back-dashboard/internal/adapters/persistence/gorm/models"
	"back-dashboard/internal/domain/user"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	var m models.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	return toDomain(&m), nil
}

func (r *UserRepo) Create(ctx context.Context, u *user.User) error {
	m := fromDomain(u)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	u.ID = m.ID
	return nil
}

func (r *UserRepo) CountAll(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func toDomain(m *models.User) *user.User {
	return &user.User{
		ID:           m.ID,
		Username:     m.Username,
		PasswordHash: m.PasswordHash,
		CreatedAt:    m.CreatedAt,
	}
}

func fromDomain(u *user.User) *models.User {
	return &models.User{
		ID:           u.ID,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
	}
}
