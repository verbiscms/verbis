// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
)

// Owner
//
// Returns a the owner of the installation.
// Logs errors.INTERNAL if there was an error executing the query.
// Logs errors.NOTFOUND if the category was not found by the given ID.
func (s *Store) Owner() domain.User {
	const op = "userStore.Owner"

	q := s.selectStmt().
		Where(s.Schema()+"roles.id", "=", domain.OwnerRoleID).
		Limit(1)

	var user domain.User
	err := s.DB().Get(&user, q.Build())
	if err == sql.ErrNoRows {
		logger.WithError(&errors.Error{Code: errors.NOTFOUND, Message: "No owner exists", Operation: op, Err: err}).Panic()
	} else if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Panic()
	}

	return user
}
