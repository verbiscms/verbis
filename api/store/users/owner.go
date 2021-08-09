// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
)

// Owner
//
// Returns a the owner of the installation.
// Logs errors.INTERNAL if there was an error executing the query.
// Logs errors.NOTFOUND if the category was not found by the given ID.
func (s *Store) Owner() domain.User {
	const op = "UserStore.Owner"

	q := s.selectStmt().
		Where(s.Schema()+"roles.id", "=", domain.OwnerRoleID).
		Limit(1)

	var user domain.User
	err := s.DB().Get(&user, q.Build())
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.NOTFOUND, Message: "No owner exists", Operation: op, Err: err}).Fatal()
	}

	return user
}
