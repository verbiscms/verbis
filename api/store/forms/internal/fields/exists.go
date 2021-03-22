// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
)

// Exists
//
// Returns a bool indicating if the form field exists by key.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) Exists(formID int, f domain.FormField) bool {
	const op = "FieldStore.Exists"

	q := s.Builder().
		Select("id").
		From(s.Schema()+TableName).
		Where("key", "=", f.Key).
		Where("form_id", "=", formID).
		Exists()

	var exists bool
	err := s.DB().QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}
