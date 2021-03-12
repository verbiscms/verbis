// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postmeta

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store"
)

// Repository defines methods for meta options
// to interact with the database.
type Repository interface {
	Insert(id int, p domain.PostOptions) error
	Exists(id int) bool
}

// Store defines the data layer for meta.
type Store struct {
	*store.Config
}

const (
	// The database table name for meta options.
	TableName = "post_options"
)

// New
//
// Creates a new meta store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
