// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
)

// Create
//
// Create a new category
// Returns errors.CONFLICT if the the category (name) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Update(c domain.Category) (domain.Category, error) {
	const op = "CategoryStore.Create"

	//exists, err := s.ExistsByName(c.Name)
	//if exists || err != nil {
	//	return domain.Category{}, &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the post, the name %v, already exists", c.Name), Operation: op, Err: fmt.Errorf("name already exists")}
	//}

	q := s.Builder().Update(TableName).
		Column("uuid", "?").
		Column("slug", c.Slug).
		Column("name", c.Name).
		Column("description", c.Description).
		Column("parent_id", c.ParentId).
		Column("resource", c.Resource).
		Column("archive_id", c.ArchiveId).
		Column("updated_at", "NOW()").
		Column("created_at", "NOW()").
		Where("id", "=", c.Id).
		Build()

	result, err := s.DB.Exec(q, uuid.New().String())
	if err != nil {
		fmt.Println(err)
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating category with the name: " + c.Name, Operation: op, Err: err}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created category ID", Operation: op, Err: err}
	}
	c.Id = int(id)

	return c, nil
}
