// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserPart
	Password      string     `db:"password" json:"password,omitempty" binding:""`
	Token         string     `db:"token" json:"token,omitempty"`
	TokenLastUsed *time.Time `db:"token_last_used" json:"token_last_used,omitempty"`
}

// UserPart defines the User with non-sensitive information
type UserPart struct {
	Id               int        `db:"id" json:"id"`
	UUID             uuid.UUID  `db:"uuid" json:"uuid"`
	FirstName        string     `db:"first_name" json:"first_name" binding:"required,max=150,alpha"`
	LastName         string     `db:"last_name" json:"last_name" binding:"required,max=150,alpha"`
	Email            string     `db:"email" json:"email" binding:"required,email,max=255"`
	Website          *string    `db:"website" json:"website,omitempty" binding:"omitempty,url"`
	Facebook         *string    `db:"facebook" json:"facebook"`
	Twitter          *string    `db:"twitter" json:"twitter"`
	Linkedin         *string    `db:"linked_in" json:"linked_in"`
	Instagram        *string    `db:"instagram" json:"instagram"`
	Biography        *string    `db:"biography" json:"biography"`
	Role             UserRole   `db:"roles" json:"role"`
	ProfilePictureID *int       `db:"profile_picture_id" json:"profile_picture_id"`
	EmailVerifiedAt  *time.Time `db:"email_verified_at" json:"email_verified_at"`
	CreatedAt        time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at" json:"updated_at"`
}

type Users []User

type UsersParts []UserPart

func (u Users) HideCredentials() UsersParts {
	var p UsersParts
	for _, v := range u {
		p = append(p, v.HideCredentials())
	}
	return p
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

func (u *User) HidePassword() {
	u.Password = ""
}

func (u *User) HideCredentials() UserPart {
	return UserPart{
		Id:               u.Id,
		UUID:             u.UUID,
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		Email:            u.Email,
		Website:          u.Website,
		Facebook:         u.Facebook,
		Twitter:          u.Twitter,
		Linkedin:         u.Linkedin,
		Instagram:        u.Instagram,
		Biography:        u.Biography,
		Role:             u.Role,
		ProfilePictureID: u.ProfilePictureID,
		EmailVerifiedAt:  u.EmailVerifiedAt,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}
