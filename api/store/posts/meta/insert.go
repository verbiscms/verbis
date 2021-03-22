// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Insert
//
// Checks to see if the post options record exists
// before updating or creating the new record.
func (s *Store) Insert(id int, p domain.PostOptions) error {
	if s.Exists(id) {
		err := s.update(id, p)
		if err != nil {
			return err
		}
	} else {
		err := s.create(id, p)
		if err != nil {
			return err
		}
	}
	return nil
}

// create
//
// Returns nil if the meta was successfully created.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) create(id int, p domain.PostOptions) error {
	const op = "MetaStore.Create"

	// No support for marshalling json for builder currently.
	q := "INSERT INTO " + s.Schema() + "post_options (post_id, seo, meta) VALUES (?, ?, ?)"

	_, err := s.DB().Exec(q, id, p.Seo, p.Meta)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error creating meta with the post ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}

// update
//
// Returns nil if the meta was successfully updated.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) update(id int, p domain.PostOptions) error {
	const op = "MetaStore.Update"

	// No support for marshalling json for builder currently.
	q := "UPDATE " + s.Schema() + "post_options SET seo = ?, meta = ? WHERE post_id = ?"

	_, err := s.DB().Exec(q, p.Seo, p.Meta, id)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error updating meta with the post ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
