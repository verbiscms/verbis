// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"database/sql"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Update
//
// Returns an updated user.
// Returns errors.CONFLICT if the validation failed.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not obtain the newly created ID.
func (s *Store) Update(u domain.User) (domain.User, error) {
	const op = "UserStore.Create"

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("first_name", u.FirstName).
		Column("last_name", u.LastName).
		Column("email", u.Email).
		Column("website", u.Website).
		Column("facebook", u.FirstName).
		Column("twitter", u.Twitter).
		Column("linked_in", u.Linkedin).
		Column("instagram", u.Instagram).
		Column("biography", u.Biography).
		Column("profile_picture_id", u.ProfilePictureID).
		Column("updated_at", "NOW()").
		Where("id", "=", u.Id)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating user with the name: " + u.FirstName, Operation: op, Err: err}
	} else if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// Update the pivot table
	err = s.updateUserRoles(u.Id, u.Role.Id)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}
