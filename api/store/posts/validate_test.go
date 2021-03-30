// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/services/theme"
	"github.com/stretchr/testify/mock"
	"regexp"
)

func (t *PostsTestSuite) TestStore_Validate() {
	tt := map[string]struct {
		input domain.PostCreate
		mock  func(m sqlmock.Sqlmock)
		theme func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			domain.PostCreate{
				Post: domain.Post{
					PageTemplate: "template",
					PageLayout:   "layout",
				},
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.Id, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindBySlugQuery))).WillReturnRows(rows)
			},
			func(m *mocks.Repository) {
				m.On("Templates", mock.Anything).Return(domain.Templates{
					domain.Template{Key: "template", Name: "Template"},
				}, nil)
				m.On("Layouts", mock.Anything).Return(domain.Layouts{
					domain.Layout{Key: "layout", Name: "Layout"},
				}, nil)
			},
			nil,
		},
		"Slug": {
			domain.PostCreate{
				Post: domain.Post{
					PageTemplate: "template",
					PageLayout:   "layout",
					Slug:         "slug",
				},
			},
			func(m sqlmock.Sqlmock) {
				q := "SELECT EXISTS (SELECT * FROM `posts` WHERE `posts`.`slug` = 'slug' AND posts.resource IS NULL)"
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(q)).
					WillReturnRows(rows)
			},
			func(m *mocks.Repository) {
				m.On("Templates", mock.Anything).Return(domain.Templates{
					domain.Template{Key: "template", Name: "Template"},
				}, nil)
				m.On("Layouts", mock.Anything).Return(domain.Layouts{
					domain.Layout{Key: "layout", Name: "Layout"},
				}, nil)
			},
			nil,
		},
		"No Page Templates": {
			domain.PostCreate{
				Post: domain.Post{
					PageTemplate: "test",
					PageLayout:   "layout",
				},
			},
			nil,
			func(m *mocks.Repository) {
				m.On("Templates", mock.Anything).Return(domain.Templates{}, nil)
				m.On("Layouts", mock.Anything).Return(domain.Layouts{
					domain.Layout{Key: "layout", Name: "Layout"},
				}, nil)
			},
			ErrNoPageTemplate.Error(),
		},
		"Template Error": {
			domain.PostCreate{},
			nil,
			func(m *mocks.Repository) {
				m.On("Templates", mock.Anything).Return(domain.Templates{}, fmt.Errorf("error"))
			},
			"error",
		},
		"No Page Layouts": {
			domain.PostCreate{
				Post: domain.Post{
					PageTemplate: "template",
				},
			},
			nil,
			func(m *mocks.Repository) {
				m.On("Templates", mock.Anything).Return(domain.Templates{
					domain.Template{Key: "template", Name: "Template"},
				}, nil)
				m.On("Layouts", mock.Anything).Return(domain.Layouts{}, nil)
			},
			ErrNoPageLayout.Error(),
		},
		"Layout Error": {
			domain.PostCreate{
				Post: domain.Post{
					PageTemplate: "template",
					PageLayout:   "layout",
				},
			},
			nil,
			func(m *mocks.Repository) {
				m.On("Templates", mock.Anything).Return(domain.Templates{
					domain.Template{Key: "template", Name: "Template"},
				}, nil)
				m.On("Layouts", mock.Anything).Return(domain.Layouts{}, fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil)
			s.Theme = &domain.ThemeConfig{}

			theme := &mocks.Repository{}
			if test.theme != nil {
				test.theme(theme)
			}
			s.ThemeService = theme

			err := s.validate(&test.input)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			fmt.Println(t.Logger)

			t.RunT(test.want, err)
		})
	}
}

func (t *PostsTestSuite) TestStore_ValidateDB() {
	tt := map[string]struct {
		input domain.PostCreate
		mock  func(m sqlmock.Sqlmock)
		want  interface{}
	}{
		"Slug": {
			domain.PostCreate{
				Post: domain.Post{
					PageTemplate: "template",
					PageLayout:   "layout",
					Slug:         "slug",
				},
			},
			func(m sqlmock.Sqlmock) {
				q := "SELECT EXISTS (SELECT * FROM `posts` WHERE `posts`.`slug` = 'slug' AND posts.resource IS NULL)"
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(q)).
					WillReturnRows(rows)
			},
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil)
			s.Theme = &domain.ThemeConfig{}

			theme := &mocks.Repository{}
			theme.On("Templates", mock.Anything).Return(domain.Templates{
				domain.Template{Key: "template", Name: "Template"},
			}, nil)
			theme.On("Layouts", mock.Anything).Return(domain.Layouts{
				domain.Layout{Key: "layout", Name: "Layout"},
			}, nil)

			s.ThemeService = theme

			err := s.validate(&test.input)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			fmt.Println(t.Logger)

			t.RunT(test.want, err)
		})
	}
}
