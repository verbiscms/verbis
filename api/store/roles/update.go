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

// Update
//
// Returns an updated role.
// Returns errors.CONFLICT if the validation failed.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not obtain the newly created ID.
func (s *Store) Update(r domain.Role) (domain.Role, error) {
	const op = "RoleStore.Update"

	err := s.validate(r)
	if err != nil {
		return domain.Role{}, err
	}

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("name", r.Name).
		Column("description", r.Description).
		Where("id", "=", r.Id)

	_, err = s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return domain.Role{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating role with the name: " + r.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Role{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return r, nil
}
