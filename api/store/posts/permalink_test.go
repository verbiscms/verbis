// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/store/options"
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
		opts  func(repository *mocks.Repository)
		cfg   domain.ThemeConfig
		mock  func(m sqlmock.Sqlmock)
		want  string
	}{
		"Homepage": {
			domain.PostDatum{
				Post: domain.Post{Id: 1},
			},
			func(m *mocks.Repository) {
				m.On("Struct").Return(domain.Options{Homepage: 1})
			},
			domain.ThemeConfig{},
			nil,
			"/",
		},
		"Page": {
			domain.PostDatum{
				Post: domain.Post{Slug: "page"},
			},
			func(m *mocks.Repository) {
				m.On("Struct").Return(domain.Options{})
			},
			domain.ThemeConfig{},
			nil,
			"/page",
		},
		"Page With Slash": {
			domain.PostDatum{
				Post: domain.Post{Slug: "page"},
			},
			func(m *mocks.Repository) {
				m.On("Struct").Return(domain.Options{SeoEnforceSlash: true})
			},
			domain.ThemeConfig{},
			nil,
			"/page/",
		},
		"Resource": {
			domain.PostDatum{
				Post: domain.Post{Slug: "article", Resource: "news"},
			},
			func(m *mocks.Repository) {
				m.On("Struct").Return(domain.Options{})
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
			"/news/article",
		},
		"Resource With Slash": {
			domain.PostDatum{
				Post: domain.Post{Slug: "article", Resource: "news"},
			},
			func(m *mocks.Repository) {
				m.On("Struct").Return(domain.Options{SeoEnforceSlash: true})
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
			func(m *mocks.Repository) {
				m.On("Struct").Return(domain.Options{})
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
			"/news/cat/article",
		},
		"Category With Slash": {
			domain.PostDatum{
				Post:     domain.Post{Slug: "article", Resource: "news"},
				Category: &category,
			},
			func(m *mocks.Repository) {
				m.On("Struct").Return(domain.Options{SeoEnforceSlash: true})
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
			func(m *mocks.Repository) {
				m.On("Struct").Return(domain.Options{})
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
			"/news/parent/cat/article",
		},
		"Category Parent Error": {
			domain.PostDatum{
				Post:     domain.Post{Slug: "article", Resource: "news"},
				Category: &categoryChild,
			},
			func(m *mocks.Repository) {
				m.On("Struct").Return(domain.Options{})
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
			func(m *mocks.Repository) {
				m.On("Struct").Return(domain.Options{SeoEnforceSlash: true})
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
			opts := &mocks.Repository{}
			test.opts(opts)
			s.options = opts
			config.Set(test.cfg)
			got := s.permalink(&test.input)
			t.Equal(test.want, got)
		})
	}
}
