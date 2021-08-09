// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"database/sql"
	"fmt"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Find
//
// Returns a user by searching with the given ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the category was not found by the given ID.
func (s *Store) Find(id int) (domain.User, error) {
	const op = "UserStore.Find"

	q := s.selectStmt().
		Where(s.Schema()+"users.id", "=", id).
		Limit(1)

	var user domain.User
	err := s.DB().Get(&user, q.Build())
	if err == sql.ErrNoRows {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return user, nil
}

// FindByToken
//
// Returns a user by searching with the given token.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the category was not found by the given ID.
func (s *Store) FindByToken(token string) (domain.User, error) {
	const op = "UserStore.FindByToken"

	q := s.selectStmt().
		Where(s.Schema()+"users.token", "=", token).
		Limit(1)

	var user domain.User
	err := s.DB().Get(&user, q.Build())
	if err == sql.ErrNoRows {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "No user exists with the token: " + token, Operation: op, Err: err}
	} else if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return user, nil
}

// FindByEmail
//
// Returns a user by searching with the given email.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the category was not found by the given ID.
func (s *Store) FindByEmail(email string) (domain.User, error) {
	const op = "UserStore.FindByEmail"

	q := s.selectStmt().
		Where(s.Schema()+"users.email", "=", email).
		Limit(1)

	var user domain.User
	err := s.DB().Get(&user, q.Build())
	if err == sql.ErrNoRows {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "No user exists with the email: " + email, Operation: op, Err: err}
	} else if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return user, nil
}
