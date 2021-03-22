// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
)

// checkOwner
//
// checkOwner Checks if the author is set or if the author does not exist.
// Returns the owner ID under circumstances.
func (s *Store) checkOwner(id int) int {
	if id == 0 || !s.users.Exists(id) {
		return s.Owner.Id
	}
	return id
}

func (s *Store) validate(p domain.PostCreate) error {
	const op = "PostStore.Validate"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where(s.Schema()+TableName+".slug", "=", p.Slug)

	if p.Category != nil {
		q.LeftJoin(s.Schema()+"post_categories", "pc", s.Schema()+"posts.id = "+s.Schema()+"pc.post_id").
			LeftJoin(s.Schema()+"categories", "c", "pc.category_id = c.id").
			Where(s.Schema()+"c.id", "=", p.Category)
	}

	if p.Resource == nil {
		q.WhereRaw(s.Schema() + TableName + ".resource IS NULL")
	} else {
		q.Where(s.Schema()+TableName+".resource", "=", p.Resource)
	}

	var exists bool
	err := s.DB().QueryRow(q.Exists()).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}
