// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
)

// Create
//
// Returns a new post upon creation.
// Returns errors.CONFLICT if the the category (name) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(p domain.PostCreate) (domain.PostDatum, error) {
	const op = "PostsStore.Create"

	//err := s.validate(p)
	//if err != nil {
	//	return domain.PostDatum{}, err
	//}

	if p.Status == "" {
		p.Status = "draft"
	}

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("uuid", "?").
		Column("slug", p.Slug).
		Column("title", p.Title).
		Column("status", p.Status).
		Column("resource", p.Resource).
		Column("page_template", p.PageTemplate).
		Column("layout", p.PageLayout).
		Column("codeinjection_head", p.CodeInjectionHead).
		Column("codeinjection_foot", p.CodeInjectionFoot).
		Column("user_id", s.checkOwner(p.UserId)).
		Column("published_at", p.PublishedAt).
		Column("updated_at", "NOW()").
		Column("created_at", "NOW()")

	result, err := s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating post with the title: " + p.Title, Operation: op, Err: err}
	} else if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created category ID", Operation: op, Err: err}
	}
	p.Id = int(id)

	// Create the post meta.
	err = s.meta.Insert(int(id), p.SeoMeta)
	if err != nil {
		return domain.PostDatum{}, err
	}

	// Create the post fields.
	err = s.fields.Insert(int(id), p.Fields)
	if err != nil {
		return domain.PostDatum{}, err
	}

	// Create the post categories
	if p.Category != nil {
		err = s.categories.Create(int(id), *p.Category)
		if err != nil {
			return domain.PostDatum{}, err
		}
	}

	// TODO!
	return s.Find(p.Id, true)
}
