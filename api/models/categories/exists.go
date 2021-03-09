// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
)

// Exists
//
// Returns a bool indicating if the category exists by ID.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) Exists(id int) bool {
	const op = "CategoryStore.Exists"

	q := s.Builder().Select("id").From(TableName).Where("id", "=", id).Exists()

	var exists bool
	err := s.DB.QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}

// ExistsByName
//
// Returns a bool indicating if the category exists by name.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) ExistsByName(name string) bool {
	const op = "CategoryStore.Exists"

	q := s.Builder().Select("id").From(TableName).Where("name", "=", name).Exists()

	var exists bool
	err := s.DB.QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}

// ExistsBySlug
//
// Returns a bool indicating if the category exists by slug.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) ExistsBySlug(slug string) bool {
	const op = "CategoryStore.Exists"

	q := s.Builder().Select("id").From(TableName).Where("slug", "=", slug).Exists()

	var exists bool
	err := s.DB.QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}
