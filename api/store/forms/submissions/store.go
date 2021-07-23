// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package submissions

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
)

// Repository defines methods for form submissions
// to interact with the database.
type Repository interface {
	Find(formID int) (domain.FormSubmissions, error)
	Create(f domain.FormSubmission) error
	Delete(formID int) error
}

// Store defines the data layer for form
// submissions.
type Store struct {
	*config.Config
}

const (
	// The database table name for form submissions.
	TableName = "form_submissions"
)

// New
//
// Creates a new meta store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
