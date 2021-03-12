// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/store"
)

// Repository defines methods for fields
// to interact with the local FS.
type Repository interface {
}

// Store defines the data layer for fields.
type Store struct {
	*store.Config
}

var (
	// ErrFieldGroupExists is returned by validate when
	// a field group already exists.
	ErrFieldGroupExists = errors.New("field group already exists")
)

// New
//
// Creates a new categories store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
