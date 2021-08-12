// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"database/sql"
	"fmt"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/store/posts/meta"
)

// Lock
//
// Returns nil if the page was locked successfully.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if the post meta was not found.
func (s *Store) Lock(postID int, token string) error {
	const op = "PostStore.Lock"

	q := s.Builder().
		Update(s.Schema()+meta.TableName).
		Column("edit_lock", token).
		Where("post_id", "=", postID)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No post exists with the ID: %d", postID), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}

// Unlock
//
// Returns nil if the page was unlocked successfully.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if the post meta was not found.
func (s *Store) Unlock(postID int) error {
	const op = "PostStore.Lock"

	q := s.Builder().
		Update(s.Schema()+meta.TableName).
		Column("edit_lock", "").
		Where("post_id", "=", postID)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No post exists with the ID: %d", postID), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
