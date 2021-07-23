// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/store/config"
)

// Repository defines methods for roles
// to interact with the database.
type Repository interface {
	List() (domain.Roles, error)
	Find(id int) (domain.Role, error)
	Create(r domain.Role) (domain.Role, error)
	Update(r domain.Role) (domain.Role, error)
	Exists(name string) bool
}

// Store defines the data layer for roles.
type Store struct {
	*config.Config
}

const (
	// The database table name for roles.
	TableName = "roles"
)

var (
	// ErrRoleExists is returned by validate when
	// a role name already exists.
	ErrRoleExists = errors.New("role already exists")
)

// New
//
// Creates a new roles store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
