// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import "github.com/ainsleyclark/verbis/api/domain"

// postType
//
//
func (s *Store) postType(post *domain.PostDatum) domain.PostType {
	if s.Options.Homepage == post.Id {
		return domain.PostType{
			PageType: domain.HomeType,
		}
	}

	resource, ok := s.Theme.Resources[post.Slug]
	if bool(post.IsArchive) || ok {
		return domain.PostType{
			PageType: domain.ArchiveType,
			Data:     resource,
		}
	}

	// Single with resource
	if post.HasResource() {
		// TODO this should be the resource
		return domain.PostType{
			PageType: domain.SingleType,
			Data:     *post.Resource,
		}
	}

	// Check if the slug is one assigned in categories.
	cat, err := s.categories.FindBySlug(post.Slug)
	if err == nil {
		return domain.PostType{
			PageType: domain.CategoryType,
			Data:     cat,
		}
	}

	return domain.PostType{
		PageType: domain.PageType,
	}
}
