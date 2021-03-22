// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Delete
//
// Returns nil if the fork was successfully deleted.
// Returns errors.NOTFOUND if the form was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) Delete(id int) error {
	const op = "FormStore.Delete"

	q := s.Builder().
		DeleteFrom(s.Schema()+TableName).
		Where("id", "=", id)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No form exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	err = s.fields.Delete(id)
	if err != nil {
		return err
	}

	err = s.submissions.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
