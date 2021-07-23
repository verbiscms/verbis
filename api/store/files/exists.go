// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
)

// Exists
//
// Returns a bool indicating if the file exists by name.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) Exists(fileName string) bool {
	const op = "FileStore.Exists"

	q := s.Builder().
		Select("id").
		From(s.Schema()+TableName).
		Where("name", "=", fileName).
		Exists()

	var exists bool
	err := s.DB().QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}
