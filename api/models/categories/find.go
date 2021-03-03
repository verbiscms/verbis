// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Find
//
// Get the category by Id
// Returns errors.NOTFOUND if the category was not found by the given Id.
func (s *Store) Find(id int64) (domain.Category, error) {
	const op = "CategoryRepository.Find"

	q := s.Builder.Select("*").From(TableName).WhereRaw("`id` = ?").Limit(1)

	var category domain.Category
	err := s.DB.Get(&category, q.Build(), id)
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Error executing sql query", Operation: op, Err: err}
	}

	return category, nil
}

// Get the category by post
// Returns errors.NOTFOUND if the category was not found by the given Post Id.
//func (s *Store) FindByPost(id int64) (domain.Category, error) {
//const op = "CategoryRepository.GetByPost"
//nolint
//q := s.Builder.Select("*").From(TableName).WhereRaw("post_id = ?").Limit(1) //nolint
//
//var category domain.Category
//err := s.DB.Get(&category, q.Build(), id)
//if err == sql.ErrNoRows {
//	return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the post ID: %d", id), Operation: op, Err: err}
//} else if err != nil {
//	return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Error executing sql query", Operation: op, Err: err}
//}
//
////return category, nil
//
//var c domain.Category
//if err := s.DB.Get(&c, "SELECT * FROM categories c WHERE EXISTS (SELECT post_id FROM post_categories p WHERE p.post_id = ? AND c.id = p.category_id) LIMIT 1", postId); err != nil {
//	return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get category with the post ID: %d", postId), Operation: op, Err: err}
//}
//return &c, nil
//}

// FindBySlug
//
// Find a category by the given slug.
// Returns errors.NOTFOUND if the category was not found by the given slug.
func (s *Store) FindBySlug(slug string) (domain.Category, error) {
	const op = "CategoryRepository.FindBySlug"

	q := s.Builder.Select("*").From(TableName).WhereRaw("`slug` = ?").Limit(1)

	var category domain.Category
	err := s.DB.Get(&category, q.Build(), slug)
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "No category exists with the slug: " + slug, Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Error executing sql query", Operation: op, Err: err}
	}

	return category, nil
}

// FindByName
//
// Find a category by the given name.
// Returns errors.NOTFOUND if the category was not found by the given slug.
func (s *Store) FindByName(name string) (domain.Category, error) {
	const op = "CategoryRepository.FindByName"

	q := s.Builder.Select("*").From(TableName).WhereRaw("`name` = ?").Limit(1)

	var category domain.Category
	err := s.DB.Get(&category, q.Build(), name)
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "No category exists with the name: " + name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Error executing sql query", Operation: op, Err: err}
	}

	return category, nil
}
