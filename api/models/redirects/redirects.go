// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/jmoiron/sqlx"
)

// Repository defines methods for redirects
// to interact with the database.
type Repository interface {
	List(meta params.Params) (domain.Categories, int, error)
	Find(id int) (domain.Redirect, error)
	Create(redirect domain.Redirect) (domain.Redirect, error)
	Update(redirect domain.Redirect) (domain.Redirect, error)
	Delete(id int) error
	Exists(id int) bool
	ExistsByFromPath(from string) bool
}

// Store defines the data layer for Redirects.
type Store struct {
	*database.Model
}

const (
	// The redirects database table name.
	TableName = "redirects"
)

// New
//
// Creates a new Redirects store.
func New(db *sqlx.DB) *Store {
	return &Store{
		Model: database.NewModel(db),
	}
}
