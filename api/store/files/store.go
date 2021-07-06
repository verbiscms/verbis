// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/store/config"
)

// Repository defines methods for storage items
// to interact with the database.
type Repository interface {
	List(meta params.Params) (domain.Files, int, error)
	Find(id int) (domain.File, error)
	FindByURL(url string) (domain.File, error)
	Create(f domain.File) (domain.File, error)
	Delete(id int) error
}

// Store defines the data layer for roles.
type Store struct {
	*config.Config
}

const (
	// TableName defines the database table name for
	// storage items.
	TableName = "files"
)

// New
//
// Creates a new storage store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
