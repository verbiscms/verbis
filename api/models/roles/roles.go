// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/jmoiron/sqlx"
)

// RedirectRepository defines methods for Redirects
// to interact with the database.
type Repository interface {
	List() ([]domain.Role, error)
	Find(id int64) (domain.Role, error)
	Create(r *domain.Role) (domain.Role, error)
	Update(r *domain.Role) (domain.Role, error)
	Delete(id int64) error
	Exists(name string) bool
}

// RedirectStore defines the data layer for Redirects
type Store struct {
	Builder builder.Sqlbuilder
	DB      *sqlx.DB
}

const TableName = "roles"

func New(db *sqlx.DB) *Store {
	return &Store{
		Builder: builder.Sqlbuilder{
			Dialect: "mysql",
		},
		DB: db,
	}
}
