// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/store"
)

// Repository defines methods for media items
// to interact with the database.
type Repository interface {
	List(meta params.Params) (domain.MediaItems, int, error) // done
	Find(id int) (domain.Media, error)
	FindByName(name string) (domain.Media, error)
	FindByURL(url string) (string, domain.Mime, error)
	Update(m domain.Media) (domain.Media, error)
	Delete(id int) error
	Exists(fileName string) bool // done
}

// Store defines the data layer for media.
type Store struct {
	*store.Config
}

const (
	// The database table name for media.
	TableName = "media"
)

// New
//
// Creates a new media store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
