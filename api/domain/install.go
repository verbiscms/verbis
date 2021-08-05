// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

type (
	// InstallPreflight defines the data used for installing the
	// Verbis database via Vue.
	InstallPreflight struct {
		DbHost     string `json:"db_host" binding:"required"`     //nolint
		DbPort     string `json:"db_port" binding:"required"`     //nolint
		DbDatabase string `json:"db_database" binding:"required"` //nolint
		DbUser     string `json:"db_user" binding:"required"`     //nolint
		DbPassword string `json:"db_password" binding:"required"` //nolint
	}
	// InstallVerbis defines the data used for installing the
	// Verbis user information and critical system info
	// via Vue.
	InstallVerbis struct {
		DbHost              string `json:"db_host" binding:"required"`     //nolint
		DbPort              string `json:"db_port" binding:"required"`     //nolint
		DbDatabase          string `json:"db_database" binding:"required"` //nolint
		DbUser              string `json:"db_user" binding:"required"`     //nolint
		DbPassword          string `json:"db_password" binding:"required"` //nolint
		SiteTitle           string `json:"site_title" binding:"required"`
		SiteUrl             string `json:"site_url" binding:"required,url"` //nolint
		UserFirstName       string `json:"user_first_name" binding:"required,max=150,alpha"`
		UserLastName        string `json:"user_last_name" binding:"required,max=150,alpha"`
		UserEmail           string `json:"user_email" binding:"required,email,max=255"`
		UserPassword        string `json:"user_password" binding:"required,min=8,max=60"`
		UserConfirmPassword string `json:"user_confirm_password,omitempty" binding:"required,eqfield=UserPassword,required"`
		Robots              bool   `json:"robots"`
	}
)

// ToUser returns user information from the
// installation struct.
func (i *InstallVerbis) ToUser() *User {
	return &User{
		UserPart: UserPart{
			FirstName: i.UserFirstName,
			LastName:  i.UserLastName,
			Email:     i.UserEmail,
		},
		Password: i.UserPassword,
	}
}
