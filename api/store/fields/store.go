// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	location "github.com/ainsleyclark/verbis/api/services/fields/converter"
	"github.com/ainsleyclark/verbis/api/store"
)

// Repository defines methods for fields
// to interact with the local FS.
type Repository interface {
	Layout(post domain.PostDatum) domain.FieldGroups
	// TODO: Create, Update & save to storage
}

// Store defines the data layer for fields.
type Store struct {
	*store.Config
	finder location.Finder
}

// New
//
// Creates a new categories store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config: cfg,
		finder: location.NewLocation(cfg.Paths.Storage),
	}
}
