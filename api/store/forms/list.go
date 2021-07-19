// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/common/params"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// List
//
// Returns a slice of forms with the total amount.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no forms available.
func (s *Store) List(meta params.Params) (domain.Forms, int, error) {
	const op = "FormStore.List"

	q := s.Builder().
		From(s.Schema() + TableName)

	// Apply filters.
	err := database.FilterRows(s.Driver, q, meta.Filters, TableName)
	if err != nil {
		return nil, -1, err
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

	// Select forms.
	var forms domain.Forms
	err = s.DB().Select(&forms, q.Build())
	if err == sql.ErrNoRows {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No forms available", Operation: op, Err: err}
	} else if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// Count the total number of forms.
	var total int
	err = s.DB().QueryRow(countQ).Scan(&total)
	if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the total number of forms", Operation: op, Err: err}
	}

	// Assign form fields & submissions.
	for k, v := range forms {
		fields, err := s.fields.Find(v.Id)
		if err == nil {
			forms[k].Fields = fields
		}
		submission, err := s.submissions.Find(v.Id)
		if err == nil {
			forms[k].Submissions = submission
		}
	}

	return forms, total, nil
}
