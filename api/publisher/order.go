// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
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


	post, err := r.Store.Posts.GetBySlug(last)
	if err != nil {
		top, found := r.categories(last)
		if found {
			return &post, top, nil
		}
		return nil, TypeOfPage{
			"error",
			url,
		}, fmt.Errorf("404")
	}

	// Check if the segment is an archive, or the last part
	// of the url matches a resource in the theme config.
	_, ok := r.Theme.Resources[last]
	if bool(post.IsArchive) || ok {
		return &post, TypeOfPage{
			PageType: "archive",
			Data:     *post.Resource,
		}, nil
	}

	// Single with categories
	if post.HasResource() {
		return &post, TypeOfPage{
			PageType: "single",
			Data:     *post.Resource,
		}, nil
	}

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


// category parent

func (r *publish) categories(url string) (TypeOfPage, bool){
	category, err := r.Store.Categories.GetBySlug(url)
	if err == nil {
		return TypeOfPage{}, true
	}
	return TypeOfPage{
		PageType: "category",
		Data:     category.Name,
	}, false
}

