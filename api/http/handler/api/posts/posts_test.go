// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/posts"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// PostsTestSuite defines the helper used for post
// testing.
type PostsTestSuite struct {
	test.HandlerSuite
}

// TestPosts
//
// Assert testing has begun.
func TestPosts(t *testing.T) {
	suite.Run(t, &PostsTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock posts handler
// for testing.
func (t *PostsTestSuite) Setup(mf func(m *mocks.Repository)) *Posts {
	cache.Init()
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}
	return &Posts{
		Deps: &deps.Deps{
			Store: &store.Repository{
				Posts: m,
			},
		},
	}
}

var (
	// The default post used for testing.
	post = domain.Post{
		Id:           123,
		Slug:         "/post",
		Title:        "post",
		PageTemplate: "tpl",
		PageLayout:   "layout",
	}
	// The default post create used for testing.
	postCreate = domain.PostCreate{
		Post: domain.Post{
			Id:           123,
			Title:        "post",
			Slug:         "/post",
			PageTemplate: "tpl",
			PageLayout:   "layout",
		},
	}
	postData = domain.PostDatum{
		Post: domain.Post{
			Id:           123,
			Slug:         "/post",
			Title:        "post",
			PageTemplate: "tpl",
			PageLayout:   "layout",
		},
	}
	// The default post with wrong validation used for testing.
	postBadValidation = domain.PostCreate{
		Post: domain.Post{
			Id: 123,
			//	Title:        "post",
			Slug:         "/post",
			PageTemplate: "tpl",
			PageLayout:   "layout",
		},
	}
	// The default posts used for testing.
	posts = domain.PostData{
		{
			Post: domain.Post{
				Id:           123,
				Slug:         "/post",
				Title:        "post",
				PageTemplate: "tpl",
				PageLayout:   "layout",
			},
		},
		{
			Post: domain.Post{
				Id:           124,
				Slug:         "/post1",
				Title:        "post1",
				PageTemplate: "tpl",
				PageLayout:   "layout",
			},
		},
	}
	// The default params used for testing.
	defaultParams = params.Params{
		Page:           api.DefaultParams.Page,
		Limit:          15,
		OrderBy:        api.DefaultParams.OrderBy,
		OrderDirection: api.DefaultParams.OrderDirection,
		Filters:        nil,
	}
)
