// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postcategories

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Update
//
// Returns nil if there was no error updating.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not obtain the newly created ID.
func (s *Store) Update(postID int, catID int) error {
	const op = "PostCategoriesStore.Create"

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("category_id", catID).
		Where("post_id", "=", postID)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error updating post category with the post ID: %d", postID), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
