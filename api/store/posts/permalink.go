// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"strings"
)

// permalink
//
//
func (s *Store) permalink(post *domain.PostDatum) string {
	permaLink := ""

	postResource := post.Resource
	//hiddenCategory := true

	if post.HasResource() {
		resource, ok := s.Theme.Resources[*postResource]
		if ok {
			// TODO: This should be in domain.
			permaLink += "/" + strings.ReplaceAll(resource.Slug, "/", "")
			//hiddenCategory = resource.HideCategorySlug
		}
	}

	var catSlugs []string

	//if post.HasCategory() && !hiddenCategory {
	//	catSlugs = append(catSlugs, post.Category.Slug)
	//	parentID := post.Category.ParentId
	//
	//	for {
	//		if !post.Category.HasParent() {
	//			break
	//		}
	//		parentCategory, err := s.categories.Find(*parentID)
	//		if err != nil {
	//			break
	//		}
	//		catSlugs = append(catSlugs, parentCategory.Slug)
	//		parentID = parentCategory.ParentId
	//	}
	//}

	for i := len(catSlugs) - 1; i >= 0; i-- {
		permaLink += "/" + catSlugs[i]
	}

	isHome := post.IsHomepage(s.Options.Homepage)
	if !isHome {
		permaLink += "/" + post.Slug
	}

	if s.Options.SeoEnforceSlash && !isHome {
		permaLink += "/"
	}

	return permaLink
}
