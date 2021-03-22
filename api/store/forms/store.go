// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/store/forms/fields"
	"github.com/ainsleyclark/verbis/api/store/forms/submissions"
	"github.com/google/uuid"
)

// Repository defines methods for forms
// to interact with the database.
type Repository interface {
	List(meta params.Params) (domain.Forms, int, error)
	Find(id int) (domain.Form, error)
	FindByUUID(uuid uuid.UUID) (domain.Form, error)
	Create(f domain.Form) (domain.Form, error)
	Update(f domain.Form) (domain.Form, error)
	Delete(id int) error
	Fields(id int) (domain.FormFields, error)
}

// Store defines the data layer for forms.
type Store struct {
	*store.Config
	fields      fields.Repository
	submissions submissions.Repository
}

const (
	// The database table name for forms.
	TableName = "forms"
	// The database table name for form fields.
	FieldsTableName = "form_fields"
)

var (
	// ErrFormExists is returned by validate when
	// a form already exists.
	ErrFormExists = errors.New("form already exists")
)

// New
//
// Creates a new form store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config:      cfg,
		fields:      fields.New(cfg),
		submissions: submissions.New(cfg),
	}
}
