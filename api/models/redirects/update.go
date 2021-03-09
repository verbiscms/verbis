// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
)

// Update
//
// Returns an updated redirect.
// Returns errors.CONFLICT if the validation failed.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not obtain the newly created ID.
func (s *Store) Update(r domain.Redirect) (domain.Redirect, error) {
	const op = "RedirectStore.Update"

	err := s.validate(r)
	if err != nil {
		return domain.Redirect{}, err
	}

	q := s.Builder().Update(TableName).
		Column("from_path", r.From).
		Column("to_path", r.To).
		Column("code", r.Code).
		Column("updated_at", "NOW()").
		Where("id", "=", r.Id)

	_, err = s.DB.Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating redirect with the from path: " + r.From, Operation: op, Err: err}
	} else if err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return r, nil
}
