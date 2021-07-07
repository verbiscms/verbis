// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sizes

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store/config"
)

// Repository defines methods for media sizes
// to interact with the database.
type Repository interface {
	Find(mediaId int) (domain.MediaSizes, error)
	Create(mediaId int, sizes domain.MediaSizes) (domain.MediaSizes, error)
	Delete(mediaId int) error
}

// Store defines the data layer for media sizes.
type Store struct {
	*config.Config
}

const (
	// TableName is the database table name for media sizes.
	TableName = "media_sizes"
)

// New
//
// Creates a new media sizes store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
