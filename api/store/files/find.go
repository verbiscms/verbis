// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Find
//
// Returns a file by searching with the given ID.
// Returns errors.NOTFOUND if the file was not found by the given ID.
// Returns errors.INTERNAL if there was an error executing the query.
func (s *Store) Find(id int) (domain.File, error) {
	const op = "FileStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("id", "=", id).
		Limit(1)

	var r domain.File
	err := s.DB().Get(&r, q.Build())
	if err == sql.ErrNoRows {
		return domain.File{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No file exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return r, nil
}

// FindByURL
//
// Returns a file by searching with the given URL.
// Returns errors.NOTFOUND if the file was not found by the given URL
// Returns errors.INTERNAL if there was an error executing the query.
func (s *Store) FindByURL(url string) (domain.File, error) {
	const op = "FileStore.FindByURL"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("url", "=", url).
		Limit(1)

	var r domain.File
	err := s.DB().Get(&r, q.Build())
	if err == sql.ErrNoRows {
		return domain.File{}, &errors.Error{Code: errors.NOTFOUND, Message: "No file exists with the URL: " + url, Operation: op, Err: err}
	} else if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return r, nil
}
