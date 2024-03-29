// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"database/sql"
	"fmt"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Find
//
// Returns a redirect by searching with the given ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the redirect was not found by the given ID.
func (s *Store) Find(id int) (domain.Redirect, error) {
	const op = "RedirectStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("id", "=", id).
		Limit(1)

	var redirect domain.Redirect
	err := s.DB().Get(&redirect, q.Build())
	if err == sql.ErrNoRows {
		return domain.Redirect{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No redirect exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return redirect, nil
}

// FindByFrom
//
// Returns a redirect by searching with the given from path.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the redirect was not found by the given ID.
func (s *Store) FindByFrom(from string) (domain.Redirect, error) {
	const op = "RedirectStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("from_path", "=", from).
		Limit(1)

	var redirect domain.Redirect
	err := s.DB().Get(&redirect, q.Build())
	if err == sql.ErrNoRows {
		return domain.Redirect{}, &errors.Error{Code: errors.NOTFOUND, Message: "No redirect exists with the from path: " + from, Operation: op, Err: err}
	} else if err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return redirect, nil
}
