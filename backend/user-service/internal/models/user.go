package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"size:100;not null" json:"first_name"`
	LastName  string    `gorm:"size:100;not null" json:"last_name"`
	Email     string    `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Role      string    `gorm:"size:20;default:user" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
