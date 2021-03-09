// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/store"
)

// Repository defines methods for users
// to interact with the database.
type Repository interface {
	List(meta params.Params) (domain.Users, int, error) // done
	Find(id int) (domain.User, error)                   //done
	FindByToken(token string) (domain.User, error)      //done
	FindByEmail(email string) (domain.User, error)      //done
	Owner() domain.User
	Create(u domain.UserCreate) (domain.User, error)
	Update(u domain.User) (domain.User, error)
	Delete(id int) error
	CheckSession(token string) error
	ResetPassword(id int, reset domain.UserPasswordReset) error
	CheckToken(token string) (domain.User, error)
	Exists(id int) bool              //done
	ExistsByEmail(email string) bool //done
}

// Store defines the data layer for users.
type Store struct {
	*store.Config
}

const (
	// The database table name for users.
	TableName = "users"
)

var (
	// ErrUserExists is returned by validate when
	// a user already exists.
	ErrUserExists = errors.New("user already exists")
)

// New
//
// Creates a new users store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
