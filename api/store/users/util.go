// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// validate
//
// Returns nil if the validation passed on update/create.
// Returns errors.CONFLICT if the email already exists in the store.
func (s *Store) validate(c domain.User) error {
	const op = "userStore.Validate"

	exists := s.ExistsByEmail(c.Email)
	if exists {
		return &errors.Error{Code: errors.CONFLICT, Message: "Validation failed, choose another email address", Operation: op, Err: ErrUserExists}
	}

	return nil
}
