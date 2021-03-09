// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ainsleyclark/verbis/api/store"
)

// Repository defines methods for options
// to interact with the database.
type Repository interface {
}

// Store defines the data layer for options.
type Store struct {
	*store.Config
}

const (
	// The database table name for options.
	TableName = "options"
)

// New
//
// Creates a new options store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
