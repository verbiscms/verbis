// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Insert
//
// Checks to see if the form field record exists
// before updating or creating the new record.
func (s *Store) Insert(formID int, f domain.FormField) error {
	if s.Exists(formID, f) {
		err := s.update(formID, f)
		if err != nil {
			return err
		}
	} else {
		err := s.create(formID, f)
		if err != nil {
			return err
		}
	}
	return nil
}

// create
//
// Returns nil if the field was successfully created.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) create(formID int, f domain.FormField) error {
	const op = "FieldStore.Create"

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("uuid", "?").
		Column("form_id", formID).
		Column("key", f.Key).
		Column("label", f.Label).
		Column("type", f.Type).
		Column("validation", f.Validation).
		Column("required", f.Required).
		Column("options", "?")

	_, err := s.DB().Exec(q.Build(), uuid.New().String(), f.Options)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error creating form field with the key:" + f.Key, Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}

// update
//
// Returns nil if the field was successfully updated.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) update(formID int, f domain.FormField) error {
	const op = "FieldStore.Update"

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("form_id", formID).
		Column("key", f.Key).
		Column("label", f.Label).
		Column("type", f.Type).
		Column("validation", f.Validation).
		Column("required", f.Required).
		Column("options", "?")

	_, err := s.DB().Exec(q.Build(), f.Options)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating form field with the key:" + f.Key, Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
