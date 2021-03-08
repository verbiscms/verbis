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
// Deletes a category from categories and post categories table.
//
// Returns errors.NOTFOUND if the category was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) Delete(id int) error {
	const op = "CategoryStore.Delete"

	q := s.Builder().DeleteFrom(TableName).WhereRaw("`id` = ?")
	_, err := s.DB.Exec(q.Build(), id)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("No category exists with the ID: %v", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error executing sql query", Operation: op, Err: err}
	}

	err = s.DeleteFromPivot(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// DeleteFromPivot
//
//
func (s *Store) DeleteFromPivot(id int) error {
	const op = "CategoryStore.DeleteFromPivot"

	q := s.Builder().DeleteFrom(PivotTableName).WhereRaw("`category_id` = ?")
	_, err := s.DB.Exec(q.Build(), id)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("No category exists with the ID: %v", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error executing sql query", Operation: op, Err: err}
	}

	return nil
}
