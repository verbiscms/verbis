// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
)

// Create
//
// Returns a new media item upon creation.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(m domain.Media) (domain.Media, error) {
	const op = "MediaStore.Create"

	q := s.Builder().Insert(s.Schema()+TableName).
		Column("uuid", "?").
		Column("url", m.Url).
		Column("file_path", m.FilePath).
		Column("file_size", m.FileSize).
		Column("file_name", m.FileName).
		Column("sizes", m.Sizes).
		Column("mime", m.Mime).
		Column("user_id", m.UserId).
		Column("updated_at", "NOW()").
		Column("created_at", "NOW()")

	result, err := s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating category with the name: " + m.FileName, Operation: op, Err: err}
	} else if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created category ID", Operation: op, Err: err}
	}
	m.Id = int(id)

	return m, nil
}
