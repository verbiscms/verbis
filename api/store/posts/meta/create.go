// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"database/sql"
	"fmt"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// create
//
// Returns nil if the meta was successfully created.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) create(id int, p domain.PostOptions) error {
	const op = "MetaStore.Create"

	// No support for marshalling json for builder currently.
	q := "INSERT INTO " + s.Schema() + "post_options (post_id, seo, meta, edit_lock) VALUES (?, ?, ?, ?)"

	_, err := s.DB().Exec(q, id, p.Seo, p.Meta, "")
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error creating meta with the post ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
