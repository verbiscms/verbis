// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
)

// Exists
//
// Returns a bool indicating if the redirect exists by ID.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) Exists(id int) bool {
	const op = "RedirectStore.Exists"

	q := s.Builder().Select("id").From(TableName).Where("id", "=", id).Exists()

	var exists bool
	err := s.DB.QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}

// ExistsByFrom
//
// Returns a bool indicating if the redirect exists by from path.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) ExistsByFrom(from string) bool {
	const op = "RedirectStore.Exists"

	q := s.Builder().Select("id").From(TableName).Where("from_path", "=", from).Exists()

	var exists bool
	err := s.DB.QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}
