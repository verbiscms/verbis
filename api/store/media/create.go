// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"database/sql"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Create
//
// Returns a new media item upon creation.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(m domain.Media) (domain.Media, error) {
	const op = "MediaStore.Create"

	// Insert main table
	q := s.Builder().Insert(s.Schema()+TableName).
		Column("title", "").
		Column("alt", "").
		Column("description", "").
		Column("user_id", m.UserID).
		Column("file_id", m.FileID).
		Column("updated_at", "NOW()").
		Column("created_at", "NOW()")

	result, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating media item with the name: " + m.File.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created media item ID", Operation: op, Err: err}
	}
	m.ID = int(id)

	// Insert to sizes table
	sizes, err := s.sizes.Create(m.ID, m.Sizes)
	if err != nil {
		return domain.Media{}, err
	}
	m.Sizes = sizes

	return m, nil
}
