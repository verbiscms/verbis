package domain

import (
	"time"
)

// Auth
type Auth struct {}

// PasswordReset defines the struct for interacting with the
// password resets table
type PasswordReset struct {
	Email			string			`db:"email" json:"email" binding:"required,email"`
	Token			string			`db:"token" json:"token" binding:"required,email"`
	CreatedAt		time.Time		`db:"created_at" json:"created_at"`
}

