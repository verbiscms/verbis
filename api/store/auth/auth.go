// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/store/users"
)

// Repository defines methods for auth
// to interact with the database.
type Repository interface {
	Login(email, password string) (domain.User, error)
	Logout(token string) (int, error)
	ResetPassword(token, password string) error
	SendResetPassword(email string) error
	VerifyPasswordToken(token string) (domain.PasswordReset, error)
	CleanPasswordResets() error
}

// Store defines the data layer for auth.
type Store struct {
	*store.Config
	// Common util functions from the user repo.
	UserStore users.Repository
	// The function used for hashing passwords.
	hashPasswordFunc func(password string) (string, error)
}

const (
	// The database table name for password resets.
	PasswordTableName = "password_resets"
	// ErrLoginMsg is returned by login when
	// authentication failed.
	ErrLoginMsg = "These credentials don't match our records."
)

// New
//
// Creates a new auth store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config:           cfg,
		UserStore:        users.New(cfg),
		hashPasswordFunc: encryption.HashPassword,
	}
}
