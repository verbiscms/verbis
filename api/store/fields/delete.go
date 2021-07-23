// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"database/sql"
	"fmt"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Delete
//
// Deletes all fields associated with a post.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if the category was not found.
func (s *Store) Delete(postID int) error {
	const op = "FieldStore.Delete"

	q := s.Builder().
		DeleteFrom(s.Schema()+TableName).
		Where("post_id", "=", postID)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No post exists with the ID: %d", postID), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}

// deleteField
//
// Returns nil if the field was successfully deleted.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if the category was not found.
func (s *Store) deleteField(postID int, f domain.PostField) error {
	const op = "FieldStore.Delete"

	q := s.Builder().
		DeleteFrom(s.Schema()+TableName).
		Where("uuid", "=", f.UUID).
		Where("post_id", "=", postID).
		Where("field_key", "=", f.Key).
		Where("name", "=", f.Name)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: "No field exists", Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
