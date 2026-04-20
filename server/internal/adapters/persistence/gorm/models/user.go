package models

import "time"

// User est le MODÈLE PERSISTANCE. Distinct de domain/user.User.
// Porte les tags gorm et la forme SQL.
type User struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"uniqueIndex;size:100;not null;column:username"`
	PasswordHash string    `gorm:"size:255;not null;column:password_hash"`
	CreatedAt    time.Time `gorm:"not null;column:created_at"`
}

func (User) TableName() string { return "users" }
