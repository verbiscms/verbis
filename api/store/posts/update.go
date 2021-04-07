// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Update
//
// Returns an updated post.
// Returns errors.CONFLICT if the validation failed.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not obtain the newly created ID.
func (s *Store) Update(p domain.PostCreate) (domain.PostDatum, error) {
	const op = "PostStore.Create"

	oldPost, err := s.Find(p.Id, false)
	if err != nil {
		return domain.PostDatum{}, err
	}

	var checkSlug bool
	if oldPost.Slug != p.Slug {
		checkSlug = true
	}

	err = s.validate(&p, checkSlug)
	if err != nil {
		return domain.PostDatum{}, err
	}

	q := s.Builder().
		Update(s.Schema()+TableName).
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
		Where("id", "=", p.Id)

	_, err = s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating post with the title: " + p.Title, Operation: op, Err: err}
	} else if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// Update the post meta.
	err = s.meta.Insert(p.Id, p.SeoMeta)
	if err != nil {
		return domain.PostDatum{}, err
	}

	// Update the post fields.
	err = s.fields.Insert(p.Id, p.Fields)
	if err != nil {
		return domain.PostDatum{}, err
	}

	// Update the post categories
	err = s.categories.Insert(p.Id, p.Category)
	if err != nil {
		return domain.PostDatum{}, err
	}

	return s.Find(p.Id, true)
}
