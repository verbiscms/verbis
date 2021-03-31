// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	fields "github.com/ainsleyclark/verbis/api/mocks/store/fields"
	categories "github.com/ainsleyclark/verbis/api/mocks/store/posts/categories"
	meta "github.com/ainsleyclark/verbis/api/mocks/store/posts/meta"
	"github.com/ainsleyclark/verbis/api/test"
	"regexp"
)

var (
	UpdateQuery = "UPDATE `posts` SET `uuid` = ?, `slug` = '/post', `title` = 'post', `status` = '', `resource` = NULL, `page_template` = 'template', `layout` = 'layout', `codeinjection_head` = NULL, `codeinjection_foot` = NULL, `user_id` = 1, `published_at` = NULL, `updated_at` = NOW()"
)

func (t *PostsTestSuite) TestStore_Update() {
	category := 1
	repoSuccess := func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
		f.On("Insert", postCreate.Id, postCreate.Fields).Return(nil)
		m.On("Insert", postCreate.Id, domain.PostOptions{}).Return(nil)
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
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(post.Id), 1))

				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.Id, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).WillReturnRows(rows)
			},
			postDatum,
		},
		"Validation Failed": {
			domain.PostCreate{},
			repoSuccess,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(post.Id), 1))

				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.Id, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).WillReturnRows(rows)
			},
			"Validation failed, no page template exists",
		},
		"No Rows": {
			postCreate,
			repoSuccess,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(sql.ErrNoRows)
			},
			"Error updating post with the title: post",
		},
		"Internal Error": {
			postCreate,
			repoSuccess,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
			database.ErrQueryMessage,
		},
		"Meta Error": {
			postCreate,
			func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
				f.On("Insert", postCreate.Id, postCreate.Fields).Return(nil)
				m.On("Insert", postCreate.Id, domain.PostOptions{}).Return(fmt.Errorf("error"))
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(post.Id), 1))

				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.Id, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).WillReturnRows(rows)
			},
			"error",
		},
		"Fields Error": {
			postCreate,
			func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
				f.On("Insert", postCreate.Id, postCreate.Fields).Return(fmt.Errorf("error"))
				m.On("Insert", postCreate.Id, domain.PostOptions{}).Return(nil)
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(post.Id), 1))

				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.Id, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).WillReturnRows(rows)
			},
			"error",
		},
		"Category Error": {
			domain.PostCreate{
				Post: domain.Post{
					Id:           1,
					Title:        "post",
					Slug:         "/post",
					PageTemplate: "template",
					PageLayout:   "layout",
				},
				Category: &category,
				Fields:   domain.PostFields{},
			},
			func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
				c.On("Create", postCreate.Id, category).Return(fmt.Errorf("error"))
				f.On("Insert", postCreate.Id, postCreate.Fields).Return(nil)
				m.On("Insert", postCreate.Id, domain.PostOptions{}).Return(nil)
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(post.Id), 1))

				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.Id, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).WillReturnRows(rows)
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupMock(test.mock, test.repo)
			post, err := s.Update(test.input)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(post, test.want, 6)
		})
	}
}
