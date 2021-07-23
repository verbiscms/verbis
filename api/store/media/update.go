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

// Update
//
// Returns an updated media item.
// Returns errors.CONFLICT if the validation failed.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not obtain the newly created ID.
func (s *Store) Update(m domain.Media) (domain.Media, error) {
	const op = "CategoryStore.Create"

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("title", m.Title).
		Column("alt", m.Alt).
		Column("description", m.Description).
		Column("updated_at", "NOW()").
		Where("id", "=", m.Id)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating media item with the url: " + m.File.Url, Operation: op, Err: err}
	} else if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return m, nil
}
