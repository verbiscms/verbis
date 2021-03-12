// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
)

// Exists
//
// Returns a bool indicating if the option exists by name.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) Exists(name string) bool {
	const op = "OptionStore.Exists"

	q := s.Builder().
		Select("option_name").
		From(s.Schema()+TableName).
		Where("option_name", "=", name).
		Exists()

	var exists bool
	err := s.DB().QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}
