// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postfields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/store"
)

// Repository defines methods for post fields
// to interact with the database.
type Repository interface {
	Find(postID int) (domain.PostFields, error)
	FindByPostAndKey(postID int, key string) (domain.PostFields, error)
	Insert(postID int, fields domain.PostFields) error
	Delete(postID int) error
	Exists(field domain.PostField) bool
}

// Store defines the data layer for post fields.
type Store struct {
	*store.Config
}

const (
	// The database table name for post fields.
	TableName = "post_fields"
)

var (
	// ErrFieldExists is returned by validate when
	// a post field already exists.
	ErrFieldExists = errors.New("post field already exists")
)

// New
//
// Creates a new post fields store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
