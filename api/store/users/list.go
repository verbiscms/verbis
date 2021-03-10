// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"strings"
)

// List
//
// Returns a slice of users with the total amount.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no categories available.
func (s *Store) List(meta params.Params) (domain.Users, int, error) {
	const op = "UserStore.List"

	q := s.SelectStmt()

	// Check if there is a role filter, for example
	// roles.name and reorder meta.Filters
	// TODO: This should be separate (in FilterRows function)
	table := TableName
	for k, v := range meta.Filters {
		if strings.Contains(k, "roles") {
			arr := strings.Split(k, ".")
			if len(arr) > 1 {
				meta.Filters[arr[1]] = v
				delete(meta.Filters, k)
				table = "roles"
			}
		}
	}

	// Apply filters
	err := database.FilterRows(s.Driver, meta.Filters, table)
	if err != nil {
		return nil, -1, err
	}

	// Apply order
	q.OrderBy(meta.OrderBy, meta.OrderDirection)
	countQ := q.Count()

	// Apply pagination
	if !meta.LimitAll {
		q.Limit(meta.Limit).Offset((meta.Page - 1) * meta.Limit)
	}

	// Select users
	var users domain.Users
	err = s.DB().Select(&users, q.Build())
	if err == sql.ErrNoRows {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No users available", Operation: op, Err: err}
	} else if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// Count the total number of users
	var total int
	err = s.DB().QueryRow(countQ).Scan(&total)
	if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the total number of users", Operation: op, Err: err}
	}

	return users, total, nil
}
