// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/common/params"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// ListConfig defines the configuration for obtaining
// posts for Selects. Posts can be filtered by
// resource and status.
type ListConfig struct {
	Resource string
	Status   string
}

// List
//
// Returns a slice of posts with the total amount.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no posts available.
func (s *Store) List(meta params.Params, layout bool, cfg ListConfig) (domain.PostData, int, error) {
	const op = "PostStore.List"

	q := s.Builder().
		From(s.Schema() + TableName)

	// Apply filters.
	err := database.FilterRows(s.Driver, q, meta.Filters, TableName)
	if err != nil {
		return nil, -1, err
	}

	// Get by resource.
	if cfg.Resource != "all" && cfg.Resource != "" {
		if cfg.Resource == "pages" {
			q.Where(s.Schema()+TableName+".resource", "=", "")
		} else {
			q.Where(s.Schema()+TableName+".resource", "=", cfg.Resource)
		}
	}

	// Get status.
	if cfg.Status != "" {
		q.Where(s.Schema()+TableName+".status", "=", cfg.Status)
	}

	// Apply order.
	if meta.OrderBy != "" {
		q.OrderBy(meta.OrderBy, meta.OrderDirection)
	}
	countQ := q.Count()

	// Apply pagination.
	if !meta.LimitAll {
		q.Limit(meta.Limit).Offset((meta.Page - 1) * meta.Limit)
	}

	built := selectStmt(q.Build())

	if meta.OrderBy != "" {
		built += fmt.Sprintf(" ORDER BY posts.%s %s", meta.OrderBy, meta.OrderDirection)
	}

	// Select posts raw.
	var raw []postsRaw
	err = s.DB().Select(&raw, built)
	if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// Count the total number of posts.
	var total int
	err = s.DB().QueryRow(countQ).Scan(&total)
	if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the total number of posts", Operation: op, Err: err}
	}

	// Return not found error if no posts are available
	posts := s.format(raw, layout)
	if len(posts) == 0 {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No posts available", Operation: op}
	}

	return posts, total, nil
}
