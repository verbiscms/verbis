// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"golang.org/x/crypto/bcrypt"
)

// Login
//
// Authenticate compares the email & password for a match in the DB.
// Returns errors.NOTFOUND if the user is not found.
func (s *Store) Login(email, password string) (domain.User, error) {
	const op = "AuthStore.Login"

	user, err := s.userStore.FindByEmail(email)
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: ErrLoginMsg, Operation: op, Err: err}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: ErrLoginMsg, Operation: op, Err: err}
	}

	err = s.userStore.UpdateToken(user.Token)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
