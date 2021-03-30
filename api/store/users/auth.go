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
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"time"
)

// CheckSession
//
// Returns nil if the session is valid.
// Returns errors.NOTFOUND if the user was not found.
// Returns errors.CONFLICT if the user session expired.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) CheckSession(token string) error {
	const op = "UserStore.CheckSession"

	u, err := s.FindByToken(token)
	if err != nil {
		return err
	}

	// The user is not logged in
	if u.TokenLastUsed == nil {
		return nil
	}

	// Destroy the token and create a new one if session expired.
	inactiveFor := time.Since(*u.TokenLastUsed).Minutes()
	if int(inactiveFor) > s.Config.Theme.Admin.InactiveSessionTime {
		newToken := encryption.GenerateUserToken(u.FirstName+u.LastName, u.Email)

		q := s.Builder().
			Update(s.Schema()+TableName).
			Column("token", "?").
			Column("updated_at", "NOW()").
			Where("token", "=", token)

		_, err := s.DB().Exec(q.Build(), newToken)
		if err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error updating the user's token with the name: %v %v", u.FirstName, u.LastName), Operation: op, Err: err}
		}

		return &errors.Error{Code: errors.CONFLICT, Message: "Session expired, please login again.", Operation: op, Err: ErrSessionExpired}
	}

	err = s.UpdateToken(token)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating the user last token used column.", Operation: op, Err: err}
	}

	return nil
}

// ResetPassword
//
// Returns nil if the password as reset successfully.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.INVALID if the user could not be updated.
func (s *Store) ResetPassword(id int, reset domain.UserPasswordReset) error {
	const op = "userStore.ResetPassword"

	password, err := s.hashPasswordFunc(reset.NewPassword)
	if err != nil {
		return err
	}

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("password", "?").
		Where("id", "=", id)

	_, err = s.DB().Exec(q.Build(), password)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INVALID, Message: "Error updating user password", Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}

// UpdateToken
//
// Returns nil if the token last used column was successfully updated.
// Returns errors.INTERNAL if the sql query failed to execute or the token does not exist.
func (s *Store) UpdateToken(token string) error {
	const op = "userStore.UpdateToken"

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("token_last_used", "NOW()").
		Where("token", "=", token)

	_, err := s.DB().Exec(q.Build(), token)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating the user last token used column.", Operation: op, Err: err}
	}

	return nil
}
