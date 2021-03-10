// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Delete
//
// Returns nil if the user was successfully deleted.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if the category was not found.
func (s *Store) Delete(id int) error {
	const op = "CategoryStore.Delete"

	if id == domain.OwnerRoleID {
		return &errors.Error{Code: errors.INTERNAL, Message: "The owner of the site cannot be deleted", Operation: op, Err: ErrDeleteOwner}
	}

	q := s.Builder().
		DeleteFrom(s.Schema()+TableName).
		Where("id", "=", id)

	_, err := s.DB().Exec(q.Build(), id)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No user exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// Delete from the pivot table
	err = s.deleteUserRoles(id)
	if err != nil {
		return err
	}

	return nil
}
