package domain

import (
	"time"
)

// Auth
type Auth struct{}

// PasswordReset defines the struct for interacting with the
// password resets table.
type PasswordReset struct {
	Id        int       `db:"id" json:"-"`
	Email     string    `db:"email" json:"email" binding:"required,email"`
	Token     string    `db:"token" json:"token" binding:"required,email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ResetPassword struct {
	NewPassword     string `json:"new_password" binding:"required,min=8,max=60"`
	ConfirmPassword string `json:"confirm_password" binding:"eqfield=NewPassword,required"`
	Token           string `db:"token" json:"token" binding:"required"`
}

type SendResetPassword struct {
	Email string `json:"email" binding:"required,email"`
}