// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Exists
//
// Determines if a redirect exists by the given ID.
//
// Returns errors.INTERNAL if there was an error executing the query.
func (s *Store) Exists(id int) (bool, error) {
	const op = "RedirectStore.Exists"

	q := s.Builder().Select("id").From(TableName).Where("id", "=", id).Exists()

	var exists bool
	err := s.DB.QueryRow(q).Scan(&exists)
	if err != nil {
		return false, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return exists, nil
}

// ExistsByFrom
//
// Determines if a redirect exists by the given from path.
//
// Returns errors.INTERNAL if there was an error executing the query.
func (s *Store) ExistsByFrom(from string) (bool, error) {
	const op = "RedirectStore.Exists"

	q := s.Builder().Select("id").From(TableName).Where("from_path", "=", from).Exists()

	var exists bool
	err := s.DB.QueryRow(q).Scan(&exists)
	if err != nil {
		return false, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return exists, nil
}
