// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	themeConfig "github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/domain"
	theme "github.com/verbiscms/verbis/api/mocks/services/theme"
	fields "github.com/verbiscms/verbis/api/mocks/store/fields"
	mocks "github.com/verbiscms/verbis/api/mocks/store/options"
	categories "github.com/verbiscms/verbis/api/mocks/store/posts/categories"
	meta "github.com/verbiscms/verbis/api/mocks/store/posts/meta"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
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

	th := &theme.Repository{}
	th.On("Templates", mock.Anything).Return(domain.Templates{
		domain.Template{Key: "template", Name: "template"},
	}, nil)
	th.On("Layouts", mock.Anything).Return(domain.Layouts{
		domain.Layout{Key: "layout", Name: "Layout"},
	}, nil)

	themeConfig.Set(domain.ThemeConfig{})

	opts := &mocks.Repository{}
	opts.On("Struct").Return(domain.Options{})
	opts.On("GetTheme").Return("theme", nil)

	s := New(&config.Config{
		Driver:       t.Driver,
		ThemeService: th,
		Owner: &domain.User{
			UserPart: domain.UserPart{
				Id: 1,
			},
		},
	})

	s.options = opts

	return s
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
	// The default field groups used for testing.
	layout domain.FieldGroups
	// The default post used for testing.
	post = domain.Post{
		Id:    1,
		Slug:  "slug",
		Title: "post",
	}
	// The default post create used for testing.
	postCreate = domain.PostCreate{
		Post: domain.Post{
			Id:           1,
			Title:        "post",
			Slug:         "slug",
			PageTemplate: "template",
			PageLayout:   "layout",
		},
		Fields: make(domain.PostFields, 0),
	}
	// The default post datum used for testing.
	postDatum = domain.PostDatum{
		Post: domain.Post{
			Id:        1,
			Slug:      "slug",
			Title:     "post",
			Permalink: "/slug",
		},
		Fields: make(domain.PostFields, 0),
		Layout: make(domain.FieldGroups, 0),
	}
	// The default post datum with layout used
	// for testing.
	postDatumLayout = domain.PostDatum{
		Post: domain.Post{
			Id:        1,
			Slug:      "slug",
			Title:     "post",
			Permalink: "/slug",
		},
		Fields: make(domain.PostFields, 0),
		Layout: layout,
	}
	// The default posts used for testing.
	posts = domain.PostData{
		{
			Post: domain.Post{
				Id:        1,
				Slug:      "slug",
				Title:     "post",
				Permalink: "/slug",
			},
		},
		{
			Post: domain.Post{
				Id:        2,
				Slug:      "slug1",
				Title:     "post1",
				Permalink: "/slug",
			},
		},
	}
	postData = domain.PostData{
		{
			Post: domain.Post{
				Id:        1,
				Slug:      "slug",
				Title:     "post",
				Permalink: "/slug",
			},
			Fields: make(domain.PostFields, 0),
		},
		{
			Post: domain.Post{
				Id:        2,
				Slug:      "slug1",
				Title:     "post1",
				Permalink: "/slug1",
			},
			Fields: make(domain.PostFields, 0),
		},
	}
)
