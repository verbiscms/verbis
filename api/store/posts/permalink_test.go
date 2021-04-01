// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"regexp"
)

var (
	FindCategoryQuery = "SELECT * FROM `categories` WHERE `id` = '2' LIMIT 1"
	parent            = 2
	category          = domain.Category{
		Id:   1,
		Slug: "cat",
		Name: "Category",
	}
	categoryChild = domain.Category{
		Id:       1,
		Slug:     "cat",
		Name:     "Category",
		ParentId: &parent,
	}
	categoryParent = domain.Category{
		Id:   1,
		Slug: "parent",
		Name: "Category",
	}
)

func (t *PostsTestSuite) TestStore_Permalink() {
	tt := map[string]struct {
		input domain.PostDatum
		opts  domain.Options
		cfg   domain.ThemeConfig
		mock  func(m sqlmock.Sqlmock)
		want  string
	}{
		"Page": {
			domain.PostDatum{
				Post: domain.Post{Slug: "page"},
			},
			domain.Options{},
			domain.ThemeConfig{},
			nil,
			"/page",
		},
		"Page With Slash": {
			domain.PostDatum{
				Post: domain.Post{Slug: "page"},
			},
			domain.Options{
				SeoEnforceSlash: true,
			},
			domain.ThemeConfig{},
			nil,
			"/page/",
		},
		"Resource": {
			domain.PostDatum{
				Post: domain.Post{Slug: "article", Resource: "news"},
			},
			domain.Options{},
			domain.ThemeConfig{
				Resources: domain.Resources{
					"news": {
						Name: "News",
						Slug: "news",
					},
				},
			},
			nil,
			"/news/article",
		},
		"Resource With Slash": {
			domain.PostDatum{
				Post: domain.Post{Slug: "article", Resource: "news"},
			},
			domain.Options{
				SeoEnforceSlash: true,
			},
			domain.ThemeConfig{
				Resources: domain.Resources{
					"news": {
						Name: "News",
						Slug: "news",
					},
				},
			},
			nil,
			"/news/article/",
		},
		"Category": {
			domain.PostDatum{
				Post:     domain.Post{Slug: "article", Resource: "news"},
				Category: &category,
			},
			domain.Options{},
			domain.ThemeConfig{
				Resources: domain.Resources{
					"news": {
						Name:             "News",
						Slug:             "news",
						HideCategorySlug: false,
					},
				},
			},
			nil,
			"/news/cat/article",
		},
		"Category With Slash": {
			domain.PostDatum{
				Post:     domain.Post{Slug: "article", Resource: "news"},
				Category: &category,
			},
			domain.Options{
				SeoEnforceSlash: true,
			},
			domain.ThemeConfig{
				Resources: domain.Resources{
					"news": {
						Name:             "News",
						Slug:             "news",
						HideCategorySlug: false,
					},
				},
			},
			nil,
			"/news/cat/article/",
		},
		"Category Parent": {
			domain.PostDatum{
				Post:     domain.Post{Slug: "article", Resource: "news"},
				Category: &categoryChild,
			},
			domain.Options{},
			domain.ThemeConfig{
				Resources: domain.Resources{
					"news": {
						Name:             "News",
						Slug:             "news",
						HideCategorySlug: false,
					},
				},
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "name"}).
					AddRow(categoryParent.Id, categoryParent.Slug, categoryParent.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindCategoryQuery)).WillReturnRows(rows)
			},
			"/news/parent/cat/article",
		},
		"Category Parent Error": {
			domain.PostDatum{
				Post:     domain.Post{Slug: "article", Resource: "news"},
				Category: &categoryChild,
			},
			domain.Options{},
			domain.ThemeConfig{
				Resources: domain.Resources{
					"news": {
						Name:             "News",
						Slug:             "news",
						HideCategorySlug: false,
					},
				},
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
			"/news/cat/article",
		},
		"Category Parent With Slash": {
			domain.PostDatum{
				Post:     domain.Post{Slug: "article", Resource: "news"},
				Category: &categoryChild,
			},
			domain.Options{
				SeoEnforceSlash: true,
			},
			domain.ThemeConfig{
				Resources: domain.Resources{
					"news": {
						Name:             "News",
						Slug:             "news",
						HideCategorySlug: false,
					},
				},
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "name"}).
					AddRow(categoryParent.Id, categoryParent.Slug, categoryParent.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindCategoryQuery)).WillReturnRows(rows)
			},
			"/news/parent/cat/article/",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			s.Options = &test.opts
			s.Theme = &test.cfg

			got := s.permalink(&test.input)
			t.Equal(test.want, got)
		})
	}
}
