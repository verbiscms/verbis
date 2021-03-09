// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
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
		Column("uuid", "?").
		Column("slug", c.Slug).
		Column("name", c.Name).
		Column("description", c.Description).
		Column("parent_id", c.ParentId).
		Column("resource", c.Resource).
		Column("archive_id", c.ArchiveId).
		Column("updated_at", "NOW()").
		Where("id", "=", c.Id)

	_, err = s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating category with the name: " + c.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return c, nil
}
