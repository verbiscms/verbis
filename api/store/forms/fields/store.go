// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
)

// Repository defines methods for form fields
// to interact with the database.
type Repository interface {
	Find(formID int) (domain.FormFields, error)
	Insert(formID int, f domain.FormField) error
	Delete(formID int) error
	Exists(formID int, f domain.FormField) bool
}

// Store defines the data layer for form
// fields.
type Store struct {
	*config.Config
}

const (
	// The database table name for form fields.
	TableName = "form_fields"
)

// New
//
// Creates a new meta store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
