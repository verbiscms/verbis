// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/store/users"
)

// ResetPassword
//
// Obtains the password reset information from the table and
// creates a new hash, it then updates the user table
// with the new details and removes the temporary
// entry in the PasswordTableName table.
//
// Returns errors.NOTFOUND if the user was not found by the given token.
// Returns errors.INTERNAL if the SQL query was invalid, unable to
// create a new password or delete from the password resets table.
func (s *Store) ResetPassword(email, password string) error {
	const op = "AuthStore.ResetPassword"

	hash, err := s.hashPasswordFunc(password)
	if err != nil {
		return err
	}

	// Update the users table.
	q := s.Builder().
		Update(s.Schema()+users.TableName).
		Column("password", "?").
		Where("email", "=", email)

	_, err = s.DB().Exec(q.Build(), hash)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating users table with the new password", Operation: op, Err: err}
	}

	return nil
}
