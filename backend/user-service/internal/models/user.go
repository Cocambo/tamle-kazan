package models

import (
	"time"
)

type User struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	FirstName         string    `gorm:"size:100;not null" json:"first_name"`
	LastName          string    `gorm:"size:100;not null" json:"last_name"`
	Email             string    `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password          string    `gorm:"size:255;not null" json:"-"`
	Role              string    `gorm:"size:20;default:user" json:"role"`
	IsEmailConfirmed  bool      `gorm:"default:false" json:"is_email_confirmed"`
	EmailTokenHash    string    `gorm:"size:64" json:"email_token_hash"`
	TokenExpiresAt    time.Time `json:"token_expires_at"`
	LastConfirmSentAt time.Time `json:"last_confirm_sent_at"`
	RefreshTokenHash  string    `gorm:"size:64" json:"-"`
	RefreshExpiresAt  time.Time `json:"-"`
	// В случае если не нужно хранить эти поля в JSON ответах, можно раскомментировать следующие строки:
	// EmailTokenHash    string    `gorm:"size:64" json:"-"`
	// TokenExpiresAt    time.Time `json:"-"`
	// LastConfirmSentAt time.Time `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
