// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/domain"
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

func (s *Store) validate(p domain.PostCreate, slug string) error {

	q := s.Builder().
		Where("slug", "=", slug)

	// Needs some work
	if p.Category != nil {
		q.LeftJoin(s.Schema()+"categories", "c", s.Schema()+"post_categories.post_id = c.id").
			Where("cat", "=", p.Category)
	}

	if p.Resource != nil {
		q.Where("resource", "=", p.Resource)
	}

	return nil
}
