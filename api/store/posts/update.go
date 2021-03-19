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

// Update
//
// Returns an updated post.
// Returns errors.CONFLICT if the validation failed.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not obtain the newly created ID.
func (s *Store) Update(p domain.PostCreate) (domain.PostDatum, error) {
	const op = "CategoryStore.Create"

	//err := s.validate(c)
	//if err != nil {
	//	return domain.Category{}, err
	//}

	q := s.Builder().
		Update(s.Schema()+TableName).
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
		Column("updated_at", "NOW()")

	_, err := s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating post with the title: " + p.Title, Operation: op, Err: err}
	} else if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// TODO!
	return domain.PostDatum{}, nil
}
