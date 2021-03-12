// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// PostsTestSuite defines the helper used for post
// testing.
type PostsTestSuite struct {
	test.DBSuite
}

// TestPosts
//
// Assert testing has begun.
func TestCategories(t *testing.T) {
	suite.Run(t, &PostsTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock posts database
// for testing.
func (t *PostsTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&store.Config{
		Driver: t.Driver,
	})
}

const (
	// The default POST ID used for testing.
	postID = "1"
)

var (
	// The default post used for testing.
	post = domain.Post{
		Id:    1,
		Slug:  "/post",
		Title: "post",
	}
	// The default post create used for testing.
	postCreate = domain.PostCreate{
		Post: domain.Post{
			Id:    1,
			Title: "post",
			Slug:  "/post",
		},
	}
	postData = domain.PostDatum{
		Post: domain.Post{
			Id:    1,
			Slug:  "/post",
			Title: "post",
		},
	}
	// The default post with wrong validation used for testing.
	postBadValidation = domain.PostCreate{
		Post: domain.Post{
			Id:    1,
			Title: "post",
		},
	}
	// The default posts used for testing.
	posts = domain.PostData{
		{
			Post: domain.Post{
				Id:    123,
				Slug:  "/post",
				Title: "post",
			},
		},
		{
			Post: domain.Post{
				Id:    124,
				Slug:  "/post1",
				Title: "post1",
			},
		},
	}
)
