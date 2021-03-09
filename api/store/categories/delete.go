// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Delete
//
// Returns nil if the category was successfully deleted.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if the category was not found.
func (s *Store) Delete(id int) error {
	const op = "CategoryStore.Delete"

	q := s.Builder().
		DeleteFrom(s.Schema()+TableName).
		Where("id", "=", id)

	_, err := s.DB().Exec(q.Build(), id)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error executing sql query", Operation: op, Err: err}
	}

	err = s.DeleteFromPivot(id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFromPivot
//
// Returns nil if the category was successfully deleted.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if the category was not found.
func (s *Store) DeleteFromPivot(id int) error {
	const op = "CategoryStore.DeleteFromPivot"

	q := s.Builder().
		DeleteFrom(s.Schema()+PivotTableName).
		Where("id", "=", id)

	_, err := s.DB().Exec(q.Build(), id)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error executing sql query", Operation: op, Err: err}
	}

	return nil
}
