// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postcategories

import (
	"github.com/ainsleyclark/verbis/api/store"
)

// Repository defines methods for post categories
// to interact with the database.
type Repository interface {
	Create(postID int, catID int) error
	Update(postID int, catID int) error
	Delete(postID int) error
}

// Store defines the data layer for post fields.
type Store struct {
	*store.Config
}

const (
	// The database table name for post categories.
	TableName = "post_categories"
)

// New
//
// Creates a new post fields store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
