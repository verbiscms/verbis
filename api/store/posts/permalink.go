// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	//"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/store/categories"
	//	"github.com/verbiscms/verbis/api/store/categories"
)

// permalink
//
// Returns the posts absolute URI with resources and
// categories. Forward slashes are added to the
// permalink if they are enabled within the
// options.
func (s *Store) permalink(post *domain.PostDatum) string {
	permaLink := ""

	opts := s.options.Struct()

	cfg, err := s.Theme.Get(opts.ActiveTheme)
	if err != nil {
		logger.WithError(err).Panic()
	}

	postResource := post.Resource
	hiddenCategory := true

	if post.IsHomepage(opts.Homepage) {
		return "/"
	}

	if post.HasResource() {
		resource, ok := cfg.Resources[postResource]
		if ok {
			permaLink += "/" + resource.Slug
			hiddenCategory = resource.HideCategorySlug
		}
	}

	var catSlugs []string
	if post.HasCategory() && !hiddenCategory {
		catSlugs = append(catSlugs, post.Category.Slug)
		parentID := post.Category.ParentId

		for {
			if parentID == nil {
				break
			}

			q := s.Builder().
				From(s.Schema()+categories.TableName).
				Where("id", "=", *parentID).
				Limit(1)

			var category domain.Category
			err := s.DB().Get(&category, q.Build())
			if err != nil {
				break
			}

			catSlugs = append(catSlugs, category.Slug)
			parentID = category.ParentId
		}
	}

	for i := len(catSlugs) - 1; i >= 0; i-- {
		permaLink += "/" + catSlugs[i]
	}

	isHome := post.IsHomepage(opts.Homepage)
	if !isHome || permaLink == "" {
		permaLink += "/" + post.Slug
	}

	if opts.SeoEnforceSlash && !isHome {
		permaLink += "/"
	}

	return permaLink
}
