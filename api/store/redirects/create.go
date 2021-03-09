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

// Create
//
// Returns a new redirect upon creation.
// Returns errors.CONFLICT if the the redirect (from_path) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(r domain.Redirect) (domain.Redirect, error) {
	const op = "RedirectStore.Create"

	err := s.validate(r)
	if err != nil {
		return domain.Redirect{}, err
	}

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("from_path", r.From).
		Column("to_path", r.To).
		Column("code", r.Code).
		Column("updated_at", "NOW()").
		Column("created_at", "NOW()")

	result, err := s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating redirect with the from path: " + r.From, Operation: op, Err: err}
	} else if err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created redirect ID", Operation: op, Err: err}
	}
	r.Id = int(id)

	return r, nil
}
