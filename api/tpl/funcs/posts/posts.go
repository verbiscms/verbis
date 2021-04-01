// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	store "github.com/ainsleyclark/verbis/api/store/posts"
	"github.com/ainsleyclark/verbis/api/tpl/params"
	"github.com/spf13/cast"
)

const (
	// The default order by field for the list function.
	OrderBy = "created_at"
	// The default order direction field for the list function.
	OrderDirection = "desc"
)

// Find
//
// Obtains the post by ID and returns a domain.PostDatum type
// or nil if not found.
//
// Example: {{ post 123 }}
func (ns *Namespace) Find(id interface{}) interface{} {
	i, err := cast.ToIntE(id)
	if err != nil || id == nil {
		return nil
	}

	post, err := ns.deps.Store.Posts.Find(i, false)
	if err != nil {
		return nil
	}

	return post.Tpl()
}

// Posts defines the struct for returning
// posts and pagination back to the
// template.
type Posts struct {
	Posts      []domain.PostTemplate
	Pagination *pagination.Pagination
}

// List
//
// Accepts a dict (map[string]interface{}) and returns an
// array of domain.post. It sets defaults if some of the param
// arguments are missing, and returns an error if the data
// could not be marshalled.

// Returns errors.TEMPLATE if the template post params failed to parse.
//
// Example:
// {{ $result := post (dict "limit" 10 "resource" "posts") }}
// {{ with $result.Posts }}
//     {{ range $post := . }}
//         <h2>{{ $post.Title }}</h2>
//         <a href="{{ $post.Slug }}">Read more</a>
//     {{ end }}
//     {{ else }}
//         <h4>No posts found</h4>
// {{ end }}
func (ns *Namespace) List(query params.Query) (interface{}, error) {
	p := query.Get(OrderBy, OrderDirection)

	resource := query.Default("resource", "")
	status := query.Default("status", "published")

	cfg := store.ListConfig{
		Resource: resource.(string),
		Status:   status.(string),
	}

	posts, total, err := ns.deps.Store.Posts.List(p, false, cfg)
	if errors.Code(err) == errors.NOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var tplPosts = make([]domain.PostTemplate, len(posts))
	for i, post := range posts {
		tplPosts[i] = post.Tpl()
	}

	return Posts{
		Posts:      tplPosts,
		Pagination: pagination.Get(p, total),
	}, nil
}
