package domain

import (
	"time"
)

type Auth struct {}

type PasswordReset struct {
	Email			string			`db:"email" json:"email" binding:"required,email"`
	Token			string			`db:"token" json:"token" binding:"required,email"`
	CreatedAt		time.Time		`db:"created_at" json:"created_at"`
}

