// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/store/posts"
	"github.com/verbiscms/verbis/api/store"
	"testing"
)

func TestCrumbs_Length(t *testing.T) {
	tt := map[string]struct {
		input Items
		want  int
	}{
		"One": {
			Items{{Link: "link-one"}},
			1,
		},
		"Two": {
			Items{{Link: "link-one"}, {Link: "link-two"}},
			2,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Length()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestItems_Reverse(t *testing.T) {
	tt := map[string]struct {
		input Items
		want  Items
	}{
		"Two": {
			Items{{Link: "link-one"}, {Link: "link-two"}},
			Items{{Link: "link-two"}, {Link: "link-one"}},
		},
		"Three": {
			Items{{Link: "link-one"}, {Link: "link-two"}, {Link: "link-three"}},
			Items{{Link: "link-three"}, {Link: "link-two"}, {Link: "link-one"}},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Reverse()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetBreadcrumbs(t *testing.T) {
	opts := domain.Options{
		BreadcrumbsEnable:       true,
		BreadcrumbsTitle:        "Title",
		BreadcrumbsSeparator:    "|",
		BreadcrumbsHomepageText: "Home",
		BreadcrumbsHideHomePage: false,
		Homepage:                1,
		SiteURL:                 "http://verbiscms.com",
	}

	tt := map[string]struct {
		input domain.PostDatum
		opts  domain.Options
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Home": {
			domain.PostDatum{
				Post: domain.Post{ID: 1},
			},
			opts,
			nil,
			Breadcrumbs{
				Enabled:   true,
				Title:     "Title",
				Separator: "|",
				Items: Items{
					{Link: "http://verbiscms.com", Text: "Home", Position: 1, Found: true, Active: true},
				},
			},
		},
		"Hide Home": {
			domain.PostDatum{
				Post: domain.Post{ID: 1},
			},
			domain.Options{
				BreadcrumbsEnable:       true,
				BreadcrumbsHideHomePage: true,
				Homepage:                1,
				SiteURL:                 "http://verbiscms.com",
			},
			nil,
			Breadcrumbs{
				Enabled: true,
			},
		},
		"Depth of One": {
			domain.PostDatum{
				Post: domain.Post{Permalink: "/news"},
			},
			opts,
			func(m *mocks.Repository) {
				m.On("FindBySlug", "news").
					Return(domain.PostDatum{Post: domain.Post{Permalink: "/news", Title: "News"}}, nil).
					Once()
			},
			Breadcrumbs{
				Enabled:   true,
				Title:     "Title",
				Separator: "|",
				Items: Items{
					{Link: "http://verbiscms.com", Text: "Home", Position: 1, Found: true, Active: false},
					{Link: "http://verbiscms.com/news", Text: "News", Position: 2, Found: true, Active: true},
				},
			},
		},
		"Depth of Two": {
			domain.PostDatum{
				Post: domain.Post{Permalink: "/news/technology"},
			},
			opts,
			func(m *mocks.Repository) {
				m.On("FindBySlug", "news").
					Return(domain.PostDatum{Post: domain.Post{Permalink: "/news", Title: "News"}}, nil).
					Once()
				m.On("FindBySlug", "technology").
					Return(domain.PostDatum{Post: domain.Post{Permalink: "/news/technology", Title: "Technology"}}, nil).
					Once()
			},
			Breadcrumbs{
				Enabled:   true,
				Title:     "Title",
				Separator: "|",
				Items: Items{
					{Link: "http://verbiscms.com", Text: "Home", Position: 1, Found: true, Active: false},
					{Link: "http://verbiscms.com/news", Text: "News", Position: 2, Found: true, Active: false},
					{Link: "http://verbiscms.com/news/technology", Text: "Technology", Position: 3, Found: true, Active: true},
				},
			},
		},
		"Depth of Three": {
			domain.PostDatum{
				Post: domain.Post{Permalink: "/news/technology/websites"},
			},
			opts,
			func(m *mocks.Repository) {
				m.On("FindBySlug", "news").
					Return(domain.PostDatum{Post: domain.Post{Permalink: "/news", Title: "News"}}, nil).
					Once()
				m.On("FindBySlug", "technology").
					Return(domain.PostDatum{Post: domain.Post{Permalink: "/news/technology", Title: "Technology"}}, nil).
					Once()
				m.On("FindBySlug", "websites").
					Return(domain.PostDatum{Post: domain.Post{Permalink: "/news/technology/websites", Title: "Websites"}}, nil).
					Once()
			},
			Breadcrumbs{
				Enabled:   true,
				Title:     "Title",
				Separator: "|",
				Items: Items{
					{Link: "http://verbiscms.com", Text: "Home", Position: 1, Found: true, Active: false},
					{Link: "http://verbiscms.com/news", Text: "News", Position: 2, Found: true, Active: false},
					{Link: "http://verbiscms.com/news/technology", Text: "Technology", Position: 3, Found: true, Active: false},
					{Link: "http://verbiscms.com/news/technology/websites", Text: "Websites", Position: 4, Found: true, Active: true},
				},
			},
		},
		"Disabled": {
			domain.PostDatum{},
			domain.Options{
				BreadcrumbsEnable: false,
			},
			nil,
			Breadcrumbs{
				Enabled: false,
			},
		},
		"Enforce Slash": {
			domain.PostDatum{
				Post: domain.Post{Permalink: "/news"},
			},
			domain.Options{
				BreadcrumbsEnable:       true,
				BreadcrumbsHomepageText: "Home",
				SeoEnforceSlash:         true,
				SiteURL:                 "http://verbiscms.com",
			},
			func(m *mocks.Repository) {
				m.On("FindBySlug", "news").
					Return(domain.PostDatum{Post: domain.Post{Permalink: "/news/", Title: "News"}}, nil).
					Once()
			},
			Breadcrumbs{
				Enabled: true,
				Items: Items{
					{Link: "http://verbiscms.com", Text: "Home", Position: 1, Found: true, Active: false},
					{Link: "http://verbiscms.com/news/", Text: "News", Position: 2, Found: true, Active: true},
				},
			},
		},
		"Enforce Slash Error": {
			domain.PostDatum{
				Post: domain.Post{Permalink: "/news"},
			},
			domain.Options{
				BreadcrumbsEnable:       true,
				BreadcrumbsHomepageText: "Home",
				SeoEnforceSlash:         true,
				SiteURL:                 "http://verbiscms.com",
			},
			func(m *mocks.Repository) {
				m.On("FindBySlug", "news").
					Return(domain.PostDatum{Post: domain.Post{}}, fmt.Errorf("error"))
			},
			Breadcrumbs{
				Enabled: true,
				Items: Items{
					{Link: "http://verbiscms.com", Text: "Home", Position: 1, Found: true, Active: false},
					{Link: "http://verbiscms.com/news/", Text: "News", Position: 2, Found: false, Active: true},
				},
			},
		},
		"Error": {
			domain.PostDatum{
				Post: domain.Post{Permalink: "/news/technology/"},
			},
			opts,
			func(m *mocks.Repository) {
				m.On("FindBySlug", "news").
					Return(domain.PostDatum{Post: domain.Post{Permalink: "/news", Title: "News"}}, nil).
					Once()
				m.On("FindBySlug", "technology").
					Return(domain.PostDatum{}, fmt.Errorf("error")).
					Once()
			},
			Breadcrumbs{
				Enabled:   true,
				Title:     "Title",
				Separator: "|",
				Items: Items{
					{Link: "http://verbiscms.com", Text: "Home", Position: 1, Found: true, Active: false},
					{Link: "http://verbiscms.com/news", Text: "News", Position: 2, Found: true, Active: false},
					{Link: "http://verbiscms.com/news/technology", Text: "Technology", Position: 3, Found: false, Active: true},
				},
			},
		},
		"Trailing Slash": {
			domain.PostDatum{
				Post: domain.Post{Permalink: "/news/"},
			},
			opts,
			func(m *mocks.Repository) {
				m.On("FindBySlug", "news").
					Return(domain.PostDatum{Post: domain.Post{Permalink: "/news", Title: "News"}}, nil).
					Once()
			},
			Breadcrumbs{
				Enabled:   true,
				Title:     "Title",
				Separator: "|",
				Items: Items{
					{Link: "http://verbiscms.com", Text: "Home", Position: 1, Found: true, Active: false},
					{Link: "http://verbiscms.com/news", Text: "News", Position: 2, Found: true, Active: true},
				},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			s := &mocks.Repository{}

			if test.mock != nil {
				test.mock(s)
			}

			got := GetBreadcrumbs(&test.input, &deps.Deps{
				Options: &test.opts,
				Store:   &store.Repository{Posts: s},
			})

			assert.Equal(t, test.want, got)
		})
	}
}

func TestSplitTitle(t *testing.T) {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Simple": {
			"post",
			"Post",
		},
		"Hyphens": {
			"post-title-test",
			"Post Title Test",
		},
		"Illegal Characters": {
			"%$£_post-ti%^tle-t&*%est",
			"Post Title Test",
		},
		"Trailing Hyphens": {
			"-post-title-test-",
			"Post Title Test",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := cleanTitle(test.input)
			assert.Equal(t, test.want, got)
		})
	}

	t.Run("Bad Regex", func(t *testing.T) {
		rg := titleRegex
		defer func() {
			titleRegex = rg
		}()
		titleRegex = "[)"
		got := cleanTitle("post")
		assert.Equal(t, "Post", got)
	})
}
