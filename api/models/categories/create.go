// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Create
//
// Create a new category
// Returns errors.CONFLICT if the the category (name) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(c domain.Category) (domain.Category, error) {
	const op = "CategoryRepository.Create"

	//if s.ExistsByName(c.Name) {
	//	return domain.Category{}, &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the post, the name %v, already exists", c.Name), Operation: op, Err: fmt.Errorf("name already exists")}
	//}

	q, err := s.Builder.BuildInsert(TableName, c, "")

	fmt.Println(err)
	fmt.Println(q)

	e, err := s.DB.Exec(q)
	if err != nil {
		fmt.Println(err)
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the category with the name: %v", c.Name), Operation: op, Err: err}
	}

	id, err := e.LastInsertId()
	if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created category ID with the name: %v", c.Name), Operation: op, Err: err}
	}
	c.Id = int(id)
	//
	//nc, err := s.Find(id)
	//if err != nil {
	//	return domain.Category{}, err
	//}

	return c, nil
}
