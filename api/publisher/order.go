// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"strings"
)

type TypeOfPage struct {
	PageType string
	Data     string
}

const (
	Page     = "page"
	Single   = "single"
	Archive  = "archive"
	Category = "category"
)

// resolve
//
//
func (r *publish) resolve(url string) (*domain.PostDatum, *TypeOfPage, error) {
	// Split the url segments.
	urlTrimmed := strings.TrimSuffix(url, "/")
	urlSplit := strings.Split(urlTrimmed, "/")
	last := urlSplit[len(urlSplit)-1]
	homepage := r.Deps.Options.Homepage

	if last == "" {
		post, err := r.Store.Posts.GetByID(homepage, false)
		if err != nil {
			return nil, nil, err
		}
		return &post, &TypeOfPage{
			PageType: "home",
		}, nil
	}

	post, err := r.Store.Posts.GetBySlug(last)

	isURLValid := r.testCheck(last, &post, urlTrimmed)
	if !isURLValid {
		return nil, nil, fmt.Errorf("not found")
	}

	if err != nil {
		trimmedPost, pErr := r.Store.Posts.GetBySlug(urlTrimmed)
		if pErr != nil {
			return nil, nil, err
		}
		// Check if its the homepage, return 404 if it is.
		if trimmedPost.IsHomepage(homepage) {
			return nil, nil, fmt.Errorf("post is the homepage")
		}
		return &trimmedPost, &TypeOfPage{
			PageType: Page,
		}, nil
	}

	// Check if the segment is an archive, or the last part
	// of the url matches a resource in the theme config.
	resource, ok := r.Config.Resources[last]
	if bool(post.IsArchive) || ok {
		return &post, &TypeOfPage{
			PageType: Archive,
			Data:     resource.Name,
		}, nil
	}

	// Single with resource
	if post.HasResource() {
		return &post, &TypeOfPage{
			PageType: Single,
			Data:     *post.Resource,
		}, nil
	}

	category, err := r.Store.Categories.GetBySlug(url)
	if err == nil {
		return &post, &TypeOfPage{
			PageType: Category,
			Data:     category.Name,
		}, nil
	}

	// Check if the post has no resources/
	// It must be a normal page.
	return &post, &TypeOfPage{
		PageType: Page,
	}, nil
}

// check if url parts match, if it doesnt return 404.
// Pass back resource category

// DONT ALLOW FORWARD SLASHES IN VUE
func (r *publish) testCheck(last string, post *domain.PostDatum, full string) bool {
	checkUrl := ""

	postResource := post.Resource
	hiddenCategory := true

	if postResource != nil {
		resource, ok := r.Config.Resources[*postResource]
		if ok {
			checkUrl += resource.Slug
			hiddenCategory = resource.HideCategorySlug
		}
	}

	var catSlugs []string

	if post.Category != nil && !hiddenCategory {
		catSlugs = append(catSlugs, post.Category.Slug)
		parent := post.Category.ParentId

		if parent != nil {

			for {
				parentCategory, err := r.Store.Categories.GetParent(*parent)
				if err != nil {
					break
				}
				catSlugs = append(catSlugs, parentCategory.Slug)
				parent = &parentCategory.Id
			}
		}
	}

	for i := len(catSlugs) - 1; i >= 0; i-- {
		checkUrl += "/" + catSlugs[i]
	}

	checkUrl += "/" + last

	if full == checkUrl {
		return true
	}

	return false
}
