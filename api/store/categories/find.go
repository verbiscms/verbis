// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Find
//
// Returns a category by searching with the given ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the category was not found by the given Id.
func (s *Store) Find(id int) (domain.Category, error) {
	const op = "CategoryStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("id", "=", id).
		Limit(1)

	var category domain.Category
	err := s.DB().Get(&category, q.Build())
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return category, nil
}

// FindByPost
//
// Returns a category by searching with the given post ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the category was not found by the given post Id.
func (s *Store) FindByPost(id int) (domain.Category, error) {
	const op = "CategoryStore.FindByPost"

	q := s.Builder().
		From(s.Schema()+"post_categories").
		LeftJoin(s.Schema()+"categories", "c", s.Schema()+"post_categories.post_id = c.id").
		Select("c.*").
		Where(s.Schema()+"post_categories.post_id", "=", id)

	var category domain.Category
	err := s.DB().Get(&category, q.Build(), id)
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the post ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return category, nil
}

// FindBySlug
//
// Returns a category by searching with the given slug.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the category was not found by the given slug.
func (s *Store) FindBySlug(slug string) (domain.Category, error) {
	const op = "CategoryStore.FindBySlug"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("slug", "=", slug).
		Limit(1)

	var category domain.Category
	err := s.DB().Get(&category, q.Build())
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "No category exists with the slug: " + slug, Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return category, nil
}

// FindByName
//
// Returns a category by searching with the given name.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the category was not found by the given slug.
func (s *Store) FindByName(name string) (domain.Category, error) {
	const op = "CategoryStore.FindByName"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("name", "=", name).
		Limit(1)

	var category domain.Category
	err := s.DB().Get(&category, q.Build())
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "No category exists with the name: " + name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return category, nil
}

// FindParent
//
// Returns a category by searching with the given parent ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the category was not found by the given slug.
func (s *Store) FindParent(id int) (domain.Category, error) {
	const op = "CategoryRepository.GetByParent"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("parent_id", "=", id).
		Limit(1)

	var category domain.Category
	err := s.DB().Get(&category, q.Build())
	if err == sql.ErrNoRows {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the parent ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return category, nil
}
