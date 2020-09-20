package domain

import (
	"github.com/google/uuid"
	"time"
)

// Return error is email is not verified, don't send the email verified at back to vue.
type User struct {
	Id					int				`db:"id" json:"id"`
	UUID 				uuid.UUID		`db:"uuid" json:"uuid"`
	FirstName			string 			`db:"first_name" json:"first_name" binding:"required,max=150,alpha"`
	LastName			string 			`db:"last_name" json:"last_name" binding:"required,max=150,alpha"`
	Email				string			`db:"email" json:"email" binding:"required,email,max=255,alpha"`
	// TODO: Need to fix update here
	//Password			string			`db:"password" json:"password,omitempty" binding:"required,min=8,max=60,alphanum"`
	Password			string			`db:"password" json:"password,omitempty" binding:"-"`
	Website				*string			`db:"website" json:"website"`
	Facebook			*string			`db:"facebook" json:"facebook"`
	Twitter				*string			`db:"twitter" json:"twitter"`
	Linkedin			*string			`db:"linked_in" json:"linked_in"`
	Instagram			*string			`db:"instagram" json:"instagram"`
	Token				string			`db:"token" json:"token,omitempty"`
	Role				UserRole 		`db:"roles" json:"role"`
	EmailVerifiedAt		*time.Time		`db:"email_verified_at" json:"email_verified_at"`
	CreatedAt			time.Time		`db:"created_at" json:"created_at"`
	UpdatedAt			time.Time		`db:"updated_at" json:"updated_at"`
}

type UserSession struct {
	Id 					int				`db:"id"`
	UserId				int				`db:"user_id"`
	SessionKey			string			`db:"session_key"`
	LoginTime			time.Time		`db:"login_time"`
	LastSeenTime 		time.Time		`db:"last_seen_time"`
}

type UserRole struct {
	Id 					int				`db:"id" json:"id" binding:"required,numeric"`
	Name				string 			`db:"name" json:"name"`
	Description			string 			`db:"description" json:"description"`
}

