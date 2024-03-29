// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"database/sql"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Update
//
// Returns an updated redirect.
// Returns errors.CONFLICT if the validation failed.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not obtain the newly created ID.
func (s *Store) Update(r domain.Redirect) (domain.Redirect, error) {
	const op = "RedirectStore.Update"

	oldRedirect, err := s.Find(r.ID)
	if err != nil {
		return domain.Redirect{}, err
	}

	if oldRedirect.From != r.From {
		err := s.validate(r)
		if err != nil {
			return domain.Redirect{}, err
		}
	}

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("from_path", r.From).
		Column("to_path", r.To).
		Column("code", r.Code).
		Column("updated_at", "NOW()").
		Where("id", "=", r.ID)

	_, err = s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating redirect with the from path: " + r.From, Operation: op, Err: err}
	} else if err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return r, nil
}
