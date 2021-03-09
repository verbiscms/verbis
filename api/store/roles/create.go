// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
)

// Create
//
// Returns a new role upon creation.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.CONFLICT if the the role ID already exists.
func (s *Store) Create(r domain.Role) (domain.Role, error) {
	const op = "RoleStore.Create"

	err := s.validate(r)
	if err != nil {
		return domain.Role{}, err
	}

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("id", r.Id).
		Column("name", r.Name).
		Column("description", r.Description)

	_, err = s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.Role{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating role with the name: " + r.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Role{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return r, nil
}
