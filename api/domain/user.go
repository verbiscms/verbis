package domain

import (
	"github.com/google/uuid"
	"time"
)

// Return error is email is not verified, don't send the email verified at back to vue.
type User struct {
	Id               int        `db:"id" json:"id"`
	UUID             uuid.UUID  `db:"uuid" json:"uuid"`
	FirstName        string     `db:"first_name" json:"first_name" binding:"required,max=150,alpha"`
	LastName         string     `db:"last_name" json:"last_name" binding:"required,max=150,alpha"`
	Email            string     `db:"email" json:"email" binding:"required,email,max=255"`
	Password         string     `db:"password" json:"password,omitempty" binding:""`
	Website          *string    `db:"website" json:"website,omitempty" binding:"omitempty,url"`
	Facebook         *string    `db:"facebook" json:"facebook"`
	Twitter          *string    `db:"twitter" json:"twitter"`
	Linkedin         *string    `db:"linked_in" json:"linked_in"`
	Instagram        *string    `db:"instagram" json:"instagram"`
	Biography        *string    `db:"biography" json:"biography"`
	ProfilePictureID *int       `db:"profile_picture_id" json:"profile_picture_id"`
	Token            string     `db:"token" json:"token,omitempty"`
	TokenLastUsed    *time.Time `db:"token_last_used" json:"token_last_used,omitempty"`
	Role             UserRole   `db:"roles" json:"role"`
	EmailVerifiedAt  *time.Time `db:"email_verified_at" json:"email_verified_at"`
	CreatedAt        time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at" json:"updated_at"`
}

type Users []User

func (u Users) HideCredentials() {
	for _, v := range u {
		v.Password = ""
		v.Token = ""
	}
}

type UserCreate struct {
	User
	Password        string `db:"password" json:"password,omitempty" binding:"required,min=8,max=60"`
	ConfirmPassword string `json:"confirm_password,omitempty" binding:"required,eqfield=Password,required"`
}

type UserPasswordReset struct {
	DBPassword      string `json:"-" binding:""`
	CurrentPassword string `json:"current_password" binding:"required,password"`
	NewPassword     string `json:"new_password" binding:"required,min=8,max=60"`
	ConfirmPassword string `json:"confirm_password" binding:"eqfield=NewPassword,required"`
}

type UserRole struct {
	Id          int    `db:"id" json:"id" binding:"required,numeric"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

func (u *User) HideCredentials() {
	u.Password = ""
	u.Token = ""
}

func (u *User) HidePassword() {
	u.Password = ""
}

