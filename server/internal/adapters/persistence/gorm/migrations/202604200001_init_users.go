package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m202604200001InitUsers() *gormigrate.Migration {
	type User struct {
		ID           uint      `gorm:"primaryKey"`
		Username     string    `gorm:"uniqueIndex;size:100;not null"`
		PasswordHash string    `gorm:"size:255;not null;column:password_hash"`
		CreatedAt    time.Time `gorm:"not null"`
	}

	return &gormigrate.Migration{
		ID: "202604200001_init_users",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
