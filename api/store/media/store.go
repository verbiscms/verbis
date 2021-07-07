// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/store/config"
	"github.com/ainsleyclark/verbis/api/store/media/sizes"
)

// Repository defines methods for media items
// to interact with the database.
type Repository interface {
	List(meta params.Params) (domain.MediaItems, int, error)
	Find(id int) (domain.Media, error)
	FindByName(name string) (domain.Media, error)
	FindByURL(url string) (domain.Media, string, error)
	Create(m domain.Media) (domain.Media, error)
	Update(m domain.Media) (domain.Media, error)
	Delete(id int) error
}

// Store defines the data layer for media.
type Store struct {
	*config.Config
	sizes sizes.Repository
}

const (
	// TableName is the database table name for media.
	TableName = "media"
	// TableSizesName is the database table name for media
	// sizes.
	TableSizesName = "media_sizes"
)

// New
//
// Creates a new media store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
		sizes:  sizes.New(cfg),
	}
}
