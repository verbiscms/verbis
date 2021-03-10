// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/google/uuid"
)

// Create
//
// Returns a new category upon creation.
// Returns errors.CONFLICT if the the category (name) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(u domain.UserCreate) (domain.User, error) {
	const op = "UserStore.Create"

	err := s.validate(u.User)
	if err != nil {
		return domain.User{}, err
	}

	password, err := s.hashPasswordFunc(u.Password)
	if err != nil {
		return domain.User{}, err
	}

	token := encryption.GenerateUserToken(u.FirstName+u.LastName, u.Email)

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("uuid", "?").
		Column("first_name", u.FirstName).
		Column("last_name", u.LastName).
		Column("email", u.Email).
		Column("password", "?").
		Column("website", u.Website).
		Column("facebook", u.FirstName).
		Column("twitter", u.Twitter).
		Column("linked_in", u.Linkedin).
		Column("instagram", u.Instagram).
		Column("biography", u.Biography).
		Column("profile_picture_id", u.ProfilePictureID).
		Column("token", "?").
		Column("updated_at", "NOW()").
		Column("created_at", "NOW()")

	result, err := s.DB().Exec(q.Build(), uuid.New().String(), password, token)
	if err == sql.ErrNoRows {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating user with the name: " + u.FirstName, Operation: op, Err: err}
	} else if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created user ID", Operation: op, Err: err}
	}
	u.User.Id = int(id)

	// Insert into the pivot table
	err = s.createUserRoles(int(id), u.Role.Id)
	if err != nil {
		return domain.User{}, err
	}

	return u.User, nil
}
