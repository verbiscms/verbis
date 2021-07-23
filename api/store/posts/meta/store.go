// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
)

// Repository defines methods for meta options
// to interact with the database.
type Repository interface {
	Insert(id int, p domain.PostOptions) error
	Delete(postID int) error
	Exists(id int) bool
}

// Store defines the data layer for meta.
type Store struct {
	*config.Config
}

const (
	// The database table name for meta options.
	TableName = "post_options"
)

// New
//
// Creates a new meta store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
