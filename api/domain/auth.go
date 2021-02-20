// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
