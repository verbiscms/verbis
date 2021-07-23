// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Find
//
// Returns a post by searching with the given ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the post was not found by the given ID.
func (s *Store) Find(id int, layout bool) (domain.PostDatum, error) {
	const op = "PostStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("id", "=", id).
		Limit(1)

	var raw []postsRaw
	err := s.DB().Select(&raw, selectStmt(q.Build()))
	if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	post := s.format(raw, layout)
	if len(post) == 0 {
		return domain.PostDatum{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No post exists with the ID: %d", id), Operation: op, Err: fmt.Errorf("no post found")}
	}

	return post[0], nil
}

// FindBySlug
//
// Returns a post by searching with the given slug.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the post was not found by the given slug.
func (s *Store) FindBySlug(slug string) (domain.PostDatum, error) {
	const op = "PostStore.FindBySlug"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("slug", "=", slug).
		Limit(1)

	var raw []postsRaw
	err := s.DB().Select(&raw, selectStmt(q.Build()))
	if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	post := s.format(raw, false)
	if len(post) == 0 {
		return domain.PostDatum{}, &errors.Error{Code: errors.NOTFOUND, Message: "No post exists with the slug: " + slug, Operation: op, Err: fmt.Errorf("no post found")}
	}

	return post[0], nil
}
