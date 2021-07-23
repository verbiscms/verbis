// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seeds

import (
	"github.com/verbiscms/verbis/api/domain"
)

// runPosts will insert all demo psots for the user.
//
//nolint
func (s *Seeder) runPosts() error {

	p := domain.PostCreate{
		Post: domain.Post{
			Slug:         "/",
			Title:        "Welcome to Verbis",
			Status:       "published",
			PageTemplate: "",
			PageLayout:   "",
			UserId:       0,
		},
		Author:   0,
		Category: nil,
	}

	_, err := s.models.Posts.Create(p)
	if err != nil {
		return err
	}

	return nil
}
