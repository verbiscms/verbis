// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Create
//
// Returns a new category upon creation.
// Returns errors.CONFLICT if the the category (name) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(c domain.Category) (domain.Category, error) {
	const op = "CategoryStore.Create"

	err := s.validate(c)
	if err != nil {
		return domain.Category{}, err
	}

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("uuid", "?").
		Column("slug", c.Slug).
		Column("name", c.Name).
		Column("description", c.Description).
		Column("parent_id", c.ParentID).
		Column("resource", c.Resource).
		Column("archive_id", c.ArchiveID).
		Column("updated_at", "NOW()").
		Column("created_at", "NOW()")

	result, err := s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating category with the name: " + c.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created category ID", Operation: op, Err: err}
	}
	c.ID = int(id)

	return c, nil
}
