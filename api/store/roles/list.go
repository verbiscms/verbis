// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"database/sql"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// List
//
// Returns a slice of categories with the total amount.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no categories available.
func (s *Store) List() (domain.Roles, error) {
	const op = "RolesStore.List"

	q := s.Builder().
		From(s.Schema()+TableName).
		OrderBy("id", "desc")

	var roles domain.Roles
	err := s.DB().Select(&roles, q.Build())
	if err == sql.ErrNoRows {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No roles available", Operation: op, Err: err}
	} else if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return roles, nil
}
