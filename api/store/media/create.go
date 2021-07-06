// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Create
//
// Returns a new media item upon creation.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(m domain.Media) (domain.Media, error) {
	const op = "MediaStore.Create"

	q := s.Builder().Insert(s.Schema()+TableName).
		Column("uuid", m.UUID.String()).
		Column("title", "").
		Column("alt", "").
		Column("description", "").
		Column("sizes", "?").
		Column("user_id", m.UserId).
		Column("updated_at", "NOW()").
		Column("created_at", "NOW()")

	result, err := s.DB().Exec(q.Build(), m.Sizes)
	if err == sql.ErrNoRows {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating category with the name: " + m.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created category ID", Operation: op, Err: err}
	}
	m.Id = int(id)

	sf, err := s.storage.Create(m.File)
	if err != nil {
		return domain.Media{}, err
	}
	m.File = sf

	return m, nil
}
