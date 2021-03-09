// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Find
//
// Returns a roles by searching with the given ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the role was not found by the given ID.
func (s *Store) Find(id int) (domain.Role, error) {
	const op = "RoleStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("id", "=", id).
		Limit(1)

	var r domain.Role
	err := s.DB().Get(&r, q.Build(), id)
	if err == sql.ErrNoRows {
		return domain.Role{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No role exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.Role{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return r, nil
}
