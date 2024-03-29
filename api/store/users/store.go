// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/verbiscms/verbis/api/common/encryption"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/database/builder"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/store/config"
)

// Repository defines methods for users
// to interact with the database.
type Repository interface {
	List(meta params.Params, role string) (domain.Users, int, error)
	Find(id int) (domain.User, error)
	FindByToken(token string) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Owner() domain.User
	Create(u domain.UserCreate) (domain.User, error)
	Update(u domain.User) (domain.User, error)
	Delete(id int) error
	CheckSession(token string) error
	ResetPassword(id int, reset domain.UserPasswordReset) error
	UpdateToken(token string) error
	Exists(id int) bool
	ExistsByEmail(email string) bool
}

// Store defines the data layer for users.
type Store struct {
	*config.Config
	// The function used for hashing passwords.
	hashPasswordFunc func(password string) (string, error)
}

const (
	// The database table name for users.
	TableName = "users"
	// The database table name for the user-roles pivot.
	PivotTableName = "user_roles"
)

var (
	// ErrUserExists is returned by validate when
	// a user already exists.
	ErrUserExists = errors.New("user already exists")
	// ErrDeleteOwner is returned by delete when
	// the owner id has been passed.
	ErrDeleteOwner = errors.New("owner cannot be deleted")
	// ErrSessionExpired is returned by check session
	// when the user session has gone passed the
	// inactive session time.
	ErrSessionExpired = errors.New("session expired")
)

// New
//
// Creates a new users store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config:           cfg,
		hashPasswordFunc: encryption.HashPassword,
	}
}

// selectStmt
//
// Helper for SELECT Statements, joining roles
// by user id
func (s *Store) selectStmt() *builder.Sqlbuilder {
	return s.Builder().
		SelectRaw(s.Schema()+"users.*, "+s.Schema()+"roles.id `roles.id`, "+s.Schema()+"roles.name `roles.name`, "+s.Schema()+"roles.description `roles.description`").
		From(s.Schema()+TableName).
		LeftJoin(s.Schema()+"user_roles", "user_roles", s.Schema()+"users.id = "+s.Schema()+"user_roles.user_id").
		LeftJoin(s.Schema()+"roles", "roles", s.Schema()+"user_roles.role_id = "+s.Schema()+"roles.id")
}
