// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
)

// List
//
// Returns a slice of media items with the total amount.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no media items available.
func (s *Store) List(meta params.Params) (domain.MediaItems, int, error) {
	const op = "MediaStore.List"

	q := s.Builder().
		From(s.Schema() + TableName)

	// Apply filters.
	err := database.FilterRows(s.Driver, q, meta.Filters, TableName)
	if err != nil {
		return nil, -1, err
	}

	// Apply order.
	q.OrderBy(meta.OrderBy, meta.OrderDirection)
	countQ := q.Count()

	// Apply pagination.
	if !meta.LimitAll {
		q.Limit(meta.Limit).Offset((meta.Page - 1) * meta.Limit)
	}

	// Select categories.
	var media domain.MediaItems
	err = s.DB().Select(&media, q.Build())
	if err == sql.ErrNoRows {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No media items available", Operation: op, Err: err}
	} else if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// Count the total number of categories.
	var total int
	err = s.DB().QueryRow(countQ).Scan(&total)
	if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the total number of media items", Operation: op, Err: err}
	}

	return media, total, nil
}
