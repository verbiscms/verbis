// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/verbiscms/verbis/api/common/encryption"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/store/users"
)

// Repository defines methods for the auth layer
// to interact with the database.
type Repository interface {
	Login(email, password string) (domain.User, error)
	Logout(token string) (int, error)
	ResetPassword(token, password string) error
	SendResetPassword(email string) (domain.UserPart, string, error)
	VerifyPasswordToken(token string) (domain.PasswordReset, error)
	CleanPasswordResets() error
}

// Store defines the data layer for auth.
type Store struct {
	*config.Config
	// Common util functions from the user repo.
	userStore users.Repository
	// The function used for hashing passwords.
	hashPasswordFunc func(password string) (string, error)
	// The function used for generating tokens.
	generateTokeFunc func(name, email string) string
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
func New(cfg *config.Config) *Store {
	return &Store{
		Config:           cfg,
		userStore:        users.New(cfg),
		hashPasswordFunc: encryption.HashPassword,
		generateTokeFunc: encryption.GenerateUserToken,
	}
}
