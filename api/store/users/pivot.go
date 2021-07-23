// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
)

// createUserRoles
//
// Returns nil if the upon inserting into the pivot table.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) createUserRoles(userID, roleID int) error {
	const op = "userStore.InsertUserRoles"

	q := s.Builder().
		Insert(s.Schema()+PivotTableName).
		Column("user_id", userID).
		Column("role_id", roleID)

	_, err := s.DB().Exec(q.Build())
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}

// updateUserRoles
//
// Returns nil if the upon updating the pivot table.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) updateUserRoles(userID, roleID int) error {
	const op = "userStore.UpdateUserRoles"

	q := s.Builder().
		Update(s.Schema()+PivotTableName).
		Column("role_id", roleID).
		Where("user_id", "=", userID)

	_, err := s.DB().Exec(q.Build())
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}

// deleteUserRoles
//
// Returns nil if the upon deleting from the pivot table.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) deleteUserRoles(userID int) error {
	const op = "userStore.UpdateUserRoles"

	q := s.Builder().
		DeleteFrom(s.Schema()+PivotTableName).
		Where("user_id", "=", userID)

	_, err := s.DB().Exec(q.Build())
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
