// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

func (t *PostsTestSuite) TestStore_Validate() {
	t.T().Skip("Skipping test validation")
	//category := 1
	//
	//tt := map[string]struct {
	//	input     domain.PostCreate
	//	mock      func(m sqlmock.Sqlmock)
	//	theme     func(m *mocks.Repository)
	//	resources domain.Resources
	//	want      interface{}
	//}{
	//	"Success": {
	//		domain.PostCreate{
	//			Post: domain.Post{
	//				PageTemplate: "template",
	//				PageLayout:   "layout",
	//				Slug:         "slug",
	//			},
	//		},
	//		func(m sqlmock.Sqlmock) {
	//			q := "SELECT EXISTS (SELECT `posts`.`id` FROM `posts` WHERE `posts`.`slug` = 'slug' AND `posts`.`resource` = '')"
	//			rows := sqlmock.NewRows([]string{"id"}).
	//				AddRow(false)
	//			m.ExpectQuery(regexp.QuoteMeta(q)).
	//				WillReturnRows(rows)
	//		},
	//		func(m *mocks.Repository) {
	//			m.On("Templates", mock.Anything).Return(domain.Templates{
	//				domain.Template{Key: "template", Name: "template"},
	//			}, nil)
	//			m.On("Layouts", mock.Anything).Return(domain.Layouts{
	//				domain.Layout{Key: "layout", Name: "Layout"},
	//			}, nil)
	//		},
	//		nil,
	//		nil,
	//	},
	//	"Hidden": {
	//		domain.PostCreate{
	//			Post: domain.Post{
	//				Resource: "news",
	//			},
	//		},
	//		nil,
	//		nil,
	//		domain.Resources{"news": domain.Resource{
	//			Hidden: true,
	//		}},
	//		nil,
	//	},
	//	"No Resource": {
	//		domain.PostCreate{
	//			Post: domain.Post{
	//				PageTemplate: "template",
	//				PageLayout:   "layout",
	//				Slug:         "slug",
	//			},
	//		},
	//		func(m sqlmock.Sqlmock) {
	//			q := "SELECT EXISTS (SELECT `posts`.`id` FROM `posts` WHERE `posts`.`slug` = 'slug' AND `posts`.`resource` = '')"
	//			rows := sqlmock.NewRows([]string{"id"}).
	//				AddRow(true)
	//			m.ExpectQuery(regexp.QuoteMeta(q)).
	//				WillReturnRows(rows)
	//		},
	//		func(m *mocks.Repository) {
	//			m.On("Templates", mock.Anything).Return(domain.Templates{
	//				domain.Template{Key: "template", Name: "template"},
	//			}, nil)
	//			m.On("Layouts", mock.Anything).Return(domain.Layouts{
	//				domain.Layout{Key: "layout", Name: "Layout"},
	//			}, nil)
	//		},
	//		nil,
	//		ErrPostsExists.Error(),
	//	},
	//	"With Resource": {
	//		domain.PostCreate{
	//			Post: domain.Post{
	//				PageTemplate: "template",
	//				PageLayout:   "layout",
	//				Slug:         "slug",
	//				Resource:     "resource",
	//			},
	//		},
	//		func(m sqlmock.Sqlmock) {
	//			q := "SELECT EXISTS (SELECT `posts`.`id` FROM `posts` WHERE `posts`.`slug` = 'slug' AND `posts`.`resource` = 'resource')"
	//			rows := sqlmock.NewRows([]string{"id"}).
	//				AddRow(true)
	//			m.ExpectQuery(regexp.QuoteMeta(q)).
	//				WillReturnRows(rows)
	//		},
	//		func(m *mocks.Repository) {
	//			m.On("Templates", mock.Anything).Return(domain.Templates{
	//				domain.Template{Key: "template", Name: "template"},
	//			}, nil)
	//			m.On("Layouts", mock.Anything).Return(domain.Layouts{
	//				domain.Layout{Key: "layout", Name: "Layout"},
	//			}, nil)
	//		},
	//		nil,
	//		ErrPostsExists.Error(),
	//	},
	//	"With Category": {
	//		domain.PostCreate{
	//			Post: domain.Post{
	//				PageTemplate: "template",
	//				PageLayout:   "layout",
	//				Slug:         "slug",
	//			},
	//			Category: &category,
	//		},
	//		func(m sqlmock.Sqlmock) {
	//			q := "SELECT EXISTS (SELECT `posts`.`id` FROM `posts` LEFT JOIN `post_categories` AS `pc` ON `posts`.`id` = `pc`.`post_id` LEFT JOIN `categories` AS `c` ON `pc`.`category_id` = `c`.`id` WHERE `posts`.`slug` = 'slug' AND `c`.`id` = '1' AND `posts`.`resource` = '')"
	//			rows := sqlmock.NewRows([]string{"id"}).
	//				AddRow(true)
	//			m.ExpectQuery(regexp.QuoteMeta(q)).
	//				WillReturnRows(rows)
	//		},
	//		func(m *mocks.Repository) {
	//			m.On("Templates", mock.Anything).Return(domain.Templates{
	//				domain.Template{Key: "template", Name: "template"},
	//			}, nil)
	//			m.On("Layouts", mock.Anything).Return(domain.Layouts{
	//				domain.Layout{Key: "layout", Name: "Layout"},
	//			}, nil)
	//		},
	//		nil,
	//		ErrPostsExists.Error(),
	//	},
	//	"No Page Templates": {
	//		domain.PostCreate{
	//			Post: domain.Post{
	//				PageTemplate: "test",
	//				PageLayout:   "layout",
	//			},
	//		},
	//		nil,
	//		func(m *mocks.Repository) {
	//			m.On("Templates", mock.Anything).Return(domain.Templates{}, nil)
	//			m.On("Layouts", mock.Anything).Return(domain.Layouts{
	//				domain.Layout{Key: "layout", Name: "Layout"},
	//			}, nil)
	//		},
	//		nil,
	//		ErrNoPageTemplate.Error(),
	//	},
	//	"Template Error": {
	//		domain.PostCreate{},
	//		nil,
	//		func(m *mocks.Repository) {
	//			m.On("Templates", mock.Anything).Return(domain.Templates{}, fmt.Errorf("error"))
	//		},
	//		nil,
	//		"error",
	//	},
	//	"No Page Layouts": {
	//		domain.PostCreate{
	//			Post: domain.Post{
	//				PageTemplate: "template",
	//			},
	//		},
	//		nil,
	//		func(m *mocks.Repository) {
	//			m.On("Templates", mock.Anything).Return(domain.Templates{
	//				domain.Template{Key: "template", Name: "template"},
	//			}, nil)
	//			m.On("Layouts", mock.Anything).Return(domain.Layouts{}, nil)
	//		},
	//		nil,
	//		ErrNoPageLayout.Error(),
	//	},
	//	"Layout Error": {
	//		domain.PostCreate{
	//			Post: domain.Post{
	//				PageTemplate: "template",
	//				PageLayout:   "layout",
	//			},
	//		},
	//		nil,
	//		func(m *mocks.Repository) {
	//			m.On("Templates", mock.Anything).Return(domain.Templates{
	//				domain.Template{Key: "template", Name: "template"},
	//			}, nil)
	//			m.On("Layouts", mock.Anything).Return(domain.Layouts{}, fmt.Errorf("error"))
	//		},
	//		nil,
	//		"error",
	//	},
	//}
	//
	//for name, test := range tt {
	//	t.Run(name, func() {
	//		s := t.Setup(test.mock)
	//		config.Set(domain.ThemeConfig{Resources: test.resources})
	//
	//		theme := &mocks.Repository{}
	//		if test.theme != nil {
	//			test.theme(theme)
	//		}
	//		s.ThemeService = theme
	//
	//		err := s.validate(&test.input, true)
	//		if err != nil {
	//			t.Contains(err.Error(), test.want)
	//			return
	//		}
	//
	//		t.RunT(test.want, err)
	//	})
	//}
}
