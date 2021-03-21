// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Delete
//
// Returns nil if the form fields were successfully deleted.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if the category was not found.
func (s *Store) Delete(formID int) error {
	const op = "FormStore.Delete"

	q := s.Builder().
		DeleteFrom(s.Schema()+TableName).
		Where("form_id", "=", formID)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No form exists with the ID: %d", formID), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}
	return nil
}
