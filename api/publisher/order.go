// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	"strings"
)

// Check if there is a resource
//  Check if there is  a category
// If there

//

type TypeOfPage struct {
	PageType string
	Data     string
}

func (r *publish) resolve(url string) (*domain.PostDatum, TypeOfPage, error) {

	// Split the url segments.
	urlTrimmed := strings.TrimSuffix(url, "/")
	urlSplit := strings.Split(urlTrimmed, "/")
	last := urlSplit[len(urlSplit)-1]

	// news/my-blog-post

	// my-blog-post

	post, err := r.Store.Posts.GetBySlug(last)
	if err != nil {


		// go cherck category

		// fails 404 handle 404

		logger.Debug(err)
	}

	// Check if the segment is an archive, or the last part
	// of the url matches a resource in the theme config.

	// news
	// insights

	_, ok := r.Theme.Resources[last]
	if bool(post.IsArchive) || ok {
		return &post, TypeOfPage{
			PageType: "archive",
			Data:     *post.Resource,
		}, nil
	}

	// news

	// Single with categories
	if post.HasResource() {
		return &post, TypeOfPage{
			PageType: "single",
			Data:     *post.Resource,
		}, nil
	}

	// news/my-category


	// Check for category archives
	category, err := r.Store.Categories.GetBySlug(last)
	if err == nil {
		return &post, TypeOfPage{
			PageType: "category",
			Data:     category.Name,
		}, nil
	}

	// category parent

	// Check if the post has no resources or categories
	// It must be a normal page.
	if !post.HasResource() {
		return &post, TypeOfPage{
			PageType: "page",
		}, nil
	}

	// No post found, it must be a 404.
	return nil, TypeOfPage{
		"error",
		url,
	}, fmt.Errorf("404")
}
