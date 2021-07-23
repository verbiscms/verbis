// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"database/sql"
	"fmt"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
)

// Delete
//
// Returns nil if the file was successfully deleted.
// Returns errors.NOTFOUND if the file was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) Delete(id int) error {
	const op = "FileStore.Delete"

	q := s.Builder().
		DeleteFrom(s.Schema()+TableName).
		Where("id", "=", id)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No file exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
