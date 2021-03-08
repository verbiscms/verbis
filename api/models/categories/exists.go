// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Exists
//
// Determines if a category exists by the given ID.
//
// Returns errors.INTERNAL if there was an error executing the query.
func (s *Store) Exists(id int) (bool, error) {
	const op = "CategoryStore.Exists"

	q := s.Builder().Select("id").From(TableName).Where("id", "=", id).Exists()

	var exists bool
	err := s.DB.QueryRow(q).Scan(&exists)
	if err != nil {
		return false, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return exists, nil
}

// ExistsByName
//
// Determines if a category exists by the given name.
//
// Returns errors.INTERNAL if there was an error executing the query.
func (s *Store) ExistsByName(name string) (bool, error) {
	const op = "CategoryStore.Exists"

	q := s.Builder().Select("id").From(TableName).Where("name", "=", name).Exists()

	var exists bool
	err := s.DB.QueryRow(q).Scan(&exists)
	if err != nil {
		return false, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return exists, nil
}

// ExistsBySlug
//
// Determines if a category exists by the given slug.
//
// Returns errors.INTERNAL if there was an error executing the query.
func (s *Store) ExistsBySlug(slug string) (bool, error) {
	const op = "CategoryStore.Exists"

	q := s.Builder().Select("id").From(TableName).Where("slug", "=", slug).Exists()

	var exists bool
	err := s.DB.QueryRow(q).Scan(&exists)
	if err != nil {
		return false, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return exists, nil
}
