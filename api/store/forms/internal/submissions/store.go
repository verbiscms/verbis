// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package submissions

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store"
)

// Repository defines methods for form submissions
// to interact with the database.
type Repository interface {
	Find(formID int) (domain.FormFields, error)
	Insert(formID int, f domain.FormField) error
	Delete(formID int) error
	Exists(formID int, f domain.FormField) bool
}

// Store defines the data layer for form
// submissions.
type Store struct {
	*store.Config
}

const (
	// The database table name for form submissions.
	TableName = "form_submissions"
)

// New
//
// Creates a new meta store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
