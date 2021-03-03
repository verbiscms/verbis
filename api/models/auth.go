// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/mail/events"
	"golang.org/x/crypto/bcrypt"
)

// AuthRepository defines methods for for Users to gain
// Auth for interacting with the database.
type AuthRepository interface {
	Authenticate(email string, password string) (domain.User, error)
	Logout(token string) (int, error)
	ResetPassword(token string, password string) error
	SendResetPassword(email string) error
	VerifyEmail(md5String string) error
	VerifyPasswordToken(token string) error
	CleanPasswordResets() error
}

// AuthStore defines the data layer for Authentication
type AuthStore struct {
	*StoreConfig
	optionsRepo domain.Options
}

// newAuth - Construct
func newAuth(cfg *StoreConfig) *AuthStore {
	return &AuthStore{
		StoreConfig: cfg,
		optionsRepo: cfg.Options.GetStruct(),
	}
}

// Authenticate compares the email & password for a match in the DB.
// Returns errors.NOTFOUND if the user is not found.
func (s *AuthStore) Authenticate(email string, password string) (domain.User, error) {
	const op = "AuthRepository.Authenticate"

	var u domain.User
	if err := s.DB.Get(&u, "SELECT * FROM users WHERE email = ? LIMIT 1", email); err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "These credentials don't match our records.", Operation: op, Err: err}
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "These credentials don't match our records.", Operation: op, Err: err}
	}

	_, err = s.DB.Exec("UPDATE users SET token_last_used = NOW() WHERE token = ?", u.Token)
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "Could not update the user token last used.", Operation: op, Err: err}
	}

	return u, nil
}

// Logout checks to see if see if the the token is valid & then
// proceeds to create a new token and returns the user Id.
// Returns errors.NOTFOUND if the user was not found by the given token.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *AuthStore) Logout(token string) (int, error) {
	const op = "AuthRepository.Logout"

	var u domain.User
	if err := s.DB.Get(&u, "SELECT * FROM users WHERE token = ? LIMIT 1", token); err != nil {
		return -1, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get user with token: %v", token), Operation: op, Err: err}
	}

	newToken := encryption.GenerateUserToken(u.FirstName+u.LastName, u.Email)
	_, err := s.DB.Exec("UPDATE users SET token = ?, updated_at = NOW() WHERE token = ?", newToken, token)
	if err != nil {
		return -1, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the user's token with the name: %v", u.FirstName+" "+u.LastName), Operation: op, Err: err}
	}

	return u.Id, nil
}

// ResetPassword obtains the password reset information from the
// table and creates a new hash, it then updates the user table
// with the new details and removes the temporary entry in
// the reset_passwords table.
// Returns errors.NOTFOUND if the user was not found by the given token.
// Returns errors.INTERNAL if the SQL query was invalid, unable to
// create a new password or delete from the password resets table.
func (s *AuthStore) ResetPassword(token string, password string) error {
	const op = "AuthRepository.ResetPassword"

	var rp domain.PasswordReset
	if err := s.DB.Get(&rp, "SELECT * FROM password_resets WHERE token = ? LIMIT 1", token); err != nil {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get user with token: %v", token), Operation: op}
	}

	hashedPassword, err := encryption.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = s.DB.Exec("UPDATE users SET password = ? WHERE email = ?", hashedPassword, rp.Email)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not update the users table with the new password", Operation: op, Err: err}
	}

	if _, err := s.DB.Exec("DELETE FROM password_resets WHERE token = ?", token); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not delete from the password resets table", Operation: op, Err: err}
	}

	return nil
}

// SendResetPassword obtains the user by email and generates a new email token.
// A temporary record is inserted to the password resets table and an email
// is sent to the user by the reset passwords event.
// Returns errors.NOTFOUND if the user was not found by the given email.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *AuthStore) SendResetPassword(email string) error {
	const op = "AuthRepository.SendResetPassword"

	var u domain.User
	if err := s.DB.Get(&u, "SELECT * FROM users WHERE email = ? LIMIT 1", email); err != nil {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not find the user with the email: %s", email), Operation: op, Err: err}
	}

	token, err := encryption.GenerateEmailToken(email)
	if err != nil {
		return err
	}

	q := "INSERT INTO password_resets (email, token, created_at) VALUES (?, ?, NOW())"
	_, err = s.DB.Exec(q, email, token)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not insert into password resets", Operation: op, Err: err}
	}

	rp, err := events.NewResetPassword()
	if err != nil {
		return err
	}

	// TODO: Clean up here
	siteURL := s.optionsRepo.SiteURL + "/admin"
	if api.SuperAdmin {
		siteURL = "http://127.0.0.1:8090/admin"
	}

	err = rp.Send(&u, siteURL, token, s.optionsRepo.SiteTitle)
	if err != nil {
		return err
	}

	return nil
}

// VerifyEmail the users email address based on the encryption hash string passed
// Returns errors.NOTFOUND if the user was not found by the md5string email.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *AuthStore) VerifyEmail(md5String string) error {
	const op = "AuthRepository.VerifyEmail"

	var userVerified = struct {
		Id   int    `db:"id"` //nolint
		Hash string `db:"hash"`
	}{}

	if err := s.DB.Get(&userVerified, "SELECT id AS id, MD5(CONCAT(id, email)) AS hash FROM users WHERE MD5(CONCAT(id, email)) = ?", md5String); err != nil {
		return &errors.Error{Code: errors.NOTFOUND, Message: "Could not find the user for email verification", Operation: op, Err: err}
	}

	q := "UPDATE users SET email_verified_at = NOW() WHERE ID = ?"
	_, err := s.DB.Exec(q, userVerified.Id)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could update the user with the Id: %d", userVerified.Id), Operation: op, Err: err}
	}

	return nil
}

// VerifyPasswordToken the token is valid from the password resets table
// Returns errors.NOTFOUND if the user was not found by the given token.
func (s *AuthStore) VerifyPasswordToken(token string) error {
	const op = "AuthRepository.VerifyPasswordToken"
	var pr domain.PasswordReset
	if err := s.DB.Get(&pr, "SELECT * FROM password_resets WHERE token = ? LIMIT 1", token); err != nil {
		return &errors.Error{Code: errors.NOTFOUND, Message: "We couldn't find a email matching that token", Operation: op, Err: err}
	}
	return nil
}

// Verify the token is valid from the password resets table
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *AuthStore) CleanPasswordResets() error {
	const op = "AuthRepository.CleanPasswordResets"
	if _, err := s.DB.Exec("DELETE FROM password_resets WHERE created_at < (NOW() - INTERVAL 2 HOUR)"); err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Could not delete from the reset passwords table", Operation: op, Err: err}
	}
	return nil
}
