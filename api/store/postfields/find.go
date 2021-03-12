// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postfields

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Find
//
// Returns a post field by searching with the given post ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the category was not found by the given ID.
func (s *Store) Find(postID int) (domain.PostFields, error) {
	const op = "PostFieldStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("post_id", "=", postID).
		Limit(1)

	var fields domain.PostFields
	err := s.DB().Select(&fields, q.Build())
	if err == sql.ErrNoRows {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No fields exists with the post ID: %d", postID), Operation: op, Err: err}
	} else if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return fields, nil
}

// FindByPostAndKey
//
// Returns a post field by searching with the given post ID & key.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the post field was not found by the given ID.
func (s *Store) FindByPostAndKey(postID int, key string) (domain.PostFields, error) {
	const op = "PostFieldStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("post_id", "=", postID).
		Where("key", "=", key).
		Limit(1)

	var fields domain.PostFields
	err := s.DB().Select(&fields, q.Build())
	if err == sql.ErrNoRows {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No fields exists with the post ID and key: %d, %s", postID, key), Operation: op, Err: err}
	} else if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return fields, nil
}
