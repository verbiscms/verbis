// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
)

// List
//
// Returns a slice of posts with the total amount.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no categories available.
func (s *Store) List(meta params.Params, layout bool, resource, status string) (domain.PostData, int, error) {
	const op = "CategoryStore.List"

	q := s.Builder().
		From(s.Schema() + TableName)

	// Apply filters.
	err := database.FilterRows(s.Driver, meta.Filters, TableName)
	if err != nil {
		return nil, -1, err
	}

	// Get by resource.
	if resource != "all" && resource != "" {
		if resource == "pages" {
			q.Where(s.Schema()+TableName, "=", "NULL")
		} else {
			q.Where(s.Schema()+TableName, "=", resource)
		}
	}

	// Get status.
	if status != "" {
		q.Where("status", "=", status)
	}

	// Apply order.
	q.OrderBy(meta.OrderBy, meta.OrderDirection)
	countQ := q.Count()

	// Apply pagination.
	if !meta.LimitAll {
		q.Limit(meta.Limit).Offset((meta.Page - 1) * meta.Limit)
	}

	// Select categories.
	var raw []postsRaw
	err = s.DB().Select(&raw, q.Build())
	if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// Count the total number of categories.
	var total int
	err = s.DB().QueryRow(countQ).Scan(&total)
	if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the total number of categories", Operation: op, Err: err}
	}

	// Return not found error if no posts are available
	posts := s.format(raw, layout)
	if len(posts) == 0 {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No posts available", Operation: op}
	}

	return posts, total, nil
}
