// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	categories "github.com/ainsleyclark/verbis/api/mocks/store/posts/categories"
	fields "github.com/ainsleyclark/verbis/api/mocks/store/posts/fields"
	meta "github.com/ainsleyclark/verbis/api/mocks/store/posts/meta"
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
func TestPosts(t *testing.T) {
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

func (t *PostsTestSuite) SetupMock(mf func(m sqlmock.Sqlmock), mfm func(c *categories.Repository, f *fields.Repository, m *meta.Repository)) *Store {
	s := t.Setup(mf)
	c := &categories.Repository{}
	f := &fields.Repository{}
	m := &meta.Repository{}
	if mfm != nil {
		mfm(c, f, m)
	}
	s.categories = c
	s.fields = f
	s.meta = m
	return s
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
	postDatum = domain.PostDatum{
		Post: domain.Post{
			Id:    1,
			Slug:  "/post",
			Title: "post",
		},
		Fields: make(domain.PostFields, 0),
	}
	// The default posts used for testing.
	posts = domain.PostData{
		{
			Post: domain.Post{
				Id:    1,
				Slug:  "/post",
				Title: "post",
			},
		},
		{
			Post: domain.Post{
				Id:    2,
				Slug:  "/post1",
				Title: "post1",
			},
		},
	}
	postData = domain.PostData{
		{
			Post: domain.Post{
				Id:    1,
				Slug:  "/post",
				Title: "post",
			},
			Fields: make(domain.PostFields, 0),
		},
		{
			Post: domain.Post{
				Id:    2,
				Slug:  "/post1",
				Title: "post1",
			},
			Fields: make(domain.PostFields, 0),
		},
	}
)
