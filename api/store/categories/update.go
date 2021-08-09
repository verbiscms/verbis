// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Update
//
// Returns an updated category.
// Returns errors.CONFLICT if the validation failed.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not obtain the newly created ID.
func (s *Store) Update(c domain.Category) (domain.Category, error) {
	const op = "CategoryStore.Create"

	err := s.validate(c)
	if err != nil {
		return domain.Category{}, err
	}

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("slug", c.Slug).
		Column("name", c.Name).
		Column("description", c.Description).
		Column("parent_id", c.ParentID).
		Column("resource", c.Resource).
		Column("archive_id", c.ArchiveID).
		Column("updated_at", "NOW()").
		Where("id", "=", c.ID)

	_, err = s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating category with the name: " + c.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return c, nil
}
