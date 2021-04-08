// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/mailer/events"
	"github.com/ainsleyclark/verbis/api/store/users"
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
func (s *Store) ResetPassword(token, password string) error {
	const op = "AuthStore.ResetPassword"

	// Check if the token is valid.
	rp, err := s.VerifyPasswordToken(token)
	if err != nil {
		return err
	}

	hash, err := s.hashPasswordFunc(password)
	if err != nil {
		return err
	}

	// Update the users table.
	q := s.Builder().
		Update(s.Schema()+users.TableName).
		Column("password", "?").
		Where("email", "=", rp.Email)

	_, err = s.DB().Exec(q.Build(), hash)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating the users table with the new password", Operation: op, Err: err}
	}

	// Delete from the password resets table.
	q = s.Builder().
		DeleteFrom(s.Schema()+PasswordTableName).
		Where("token", "=", token)

	_, err = s.DB().Exec(q.Build())
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting from the password resets table", Operation: op, Err: err})
	}

	return nil
}

// SendResetPassword
//
// Obtains the user by email and generates a new email token.
// A temporary record is inserted to the password resets
// table and an email is sent to the user by the reset
// passwords event.
// passwords event.
// Returns errors.NOTFOUND if the user was not found by the given email.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) SendResetPassword(email string) error {
	const op = "AuthRepository.SendResetPassword"

	user, err := s.userStore.FindByEmail(email)
	if err != nil {
		return err
	}

	token, err := encryption.GenerateEmailToken(email)
	if err != nil {
		return err
	}

	q := s.Builder().
		Insert(s.Schema()+PasswordTableName).
		Column("email", email).
		Column("token", token).
		Column("created_at", "NOW()")

	_, err = s.DB().Exec(q.Build())
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error inserting into password resets", Operation: op, Err: err}
	}

	// TODO: Mailer! This should be an interface with sending
	// 	methods. To test.
	rp, err := events.NewResetPassword()
	if err != nil {
		return err
	}

	err = rp.Send(&user, s.Options.SiteUrl, token, s.Options.SiteTitle)
	if err != nil {
		return err
	}

	return nil
}

// VerifyPasswordToken
//
// Checks to see if the token is valid from the password resets table.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the user was not found by the given token.
func (s *Store) VerifyPasswordToken(token string) (domain.PasswordReset, error) {
	const op = "AuthStore.VerifyPasswordToken"

	q := s.Builder().
		From(s.Schema()+PasswordTableName).
		Where("token", "=", token).
		Limit(1)

	var pr domain.PasswordReset
	err := s.DB().Get(&pr, q.Build())
	if err == sql.ErrNoRows {
		return domain.PasswordReset{}, &errors.Error{Code: errors.NOTFOUND, Message: "No user exists with the token: " + token, Operation: op, Err: err}
	} else if err != nil {
		return domain.PasswordReset{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return pr, nil
}

// CleanPasswordResets
//
// Verify the token is valid from the password resets table
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) CleanPasswordResets() error {
	const op = "AuthStore.CleanPasswordResets"

	q := s.Builder().
		DeleteFrom(s.Schema() + PasswordTableName).
		WhereRaw("created_at < (NOW() - INTERVAL 2 HOUR)")

	_, err := s.DB().Exec(q.Build())
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error deleting from the reset passwords table", Operation: op, Err: err}
	}

	return nil
}
