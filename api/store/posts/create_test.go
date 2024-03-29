// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	fields "github.com/verbiscms/verbis/api/mocks/store/fields"
	categories "github.com/verbiscms/verbis/api/mocks/store/posts/categories"
	meta "github.com/verbiscms/verbis/api/mocks/store/posts/meta"
	"github.com/verbiscms/verbis/api/test"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO `posts` (`uuid`, `slug`, `title`, `status`, `resource`, `page_template`, `layout`, `codeinjection_head`, `codeinjection_foot`, `user_id`, `published_at`, `updated_at`, `created_at`) VALUES (?, 'slug', 'post', 'draft', '', 'template', 'layout', '', '', 1, NULL, NOW(), NOW())"
)

func (t *PostsTestSuite) TestStore_Create() {
	category := 1

	repoSuccess := func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
		var cat *int
		c.On("Insert", postCreate.ID, cat).Return(nil)
		f.On("Insert", postCreate.ID, postCreate.Fields).Return(nil)
		m.On("Insert", postCreate.ID, domain.PostOptions{}).Return(nil)
	}

	tt := map[string]struct {
		input domain.PostCreate
		repo  func(c *categories.Repository, f *fields.Repository, m *meta.Repository)
		mock  func(m sqlmock.Sqlmock)
		want  interface{}
	}{
		"Success": {
			postCreate,
			repoSuccess,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(post.ID), 1))

				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.ID, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).WillReturnRows(rows)
			},
			postDatum,
		},
		"Validation Failed": {
			domain.PostCreate{},
			repoSuccess,
			func(m sqlmock.Sqlmock) {
				q := "SELECT EXISTS (SELECT `posts`.`id` FROM `posts` WHERE `posts`.`slug` = '' AND `posts`.`resource` = '')"
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(q)).
					WillReturnRows(rows)
			},
			"Validation failed, the slug already exists",
		},
		"No Rows": {
			postCreate,
			repoSuccess,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(sql.ErrNoRows)
			},
			"Error creating post with the title",
		},
		"Internal Error": {
			postCreate,
			repoSuccess,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
			database.ErrQueryMessage,
		},
		"Last Insert ID Error": {
			postCreate,
			repoSuccess,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("err")))
			},
			"Error getting the newly created post ID",
		},
		"Meta Error": {
			postCreate,
			func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
				f.On("Insert", postCreate.ID, postCreate.Fields).Return(nil)
				m.On("Insert", postCreate.ID, domain.PostOptions{}).Return(fmt.Errorf("error"))
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(post.ID), 1))

				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.ID, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).
					WillReturnRows(rows)
			},
			"error",
		},
		"Fields Error": {
			postCreate,
			func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
				f.On("Insert", postCreate.ID, postCreate.Fields).Return(fmt.Errorf("error"))
				m.On("Insert", postCreate.ID, domain.PostOptions{}).Return(nil)
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(post.ID), 1))

				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.ID, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).
					WillReturnRows(rows)
			},
			"error",
		},
		"Category Error": {
			domain.PostCreate{
				Post: domain.Post{
					ID:           1,
					Title:        "post",
					Slug:         "slug",
					PageTemplate: "template",
					PageLayout:   "layout",
				},
				Category: &category,
				Fields:   domain.PostFields{},
			},
			func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
				c.On("Insert", postCreate.ID, &category).Return(fmt.Errorf("error"))
				f.On("Insert", postCreate.ID, postCreate.Fields).Return(nil)
				m.On("Insert", postCreate.ID, domain.PostOptions{}).Return(nil)
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(post.ID), 1))

				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.ID, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).
					WillReturnRows(rows)
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupMock(test.mock, test.repo)
			post, err := s.Create(test.input)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(post, test.want, 6)
		})
	}
}
