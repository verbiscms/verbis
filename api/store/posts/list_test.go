// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"github.com/verbiscms/verbis/api/test/dummy"
	"regexp"
)

var (
	ListQuery          = "SELECT * FROM `posts` ORDER BY created_at desc LIMIT 15 OFFSET 0"
	ListResourceQuery  = "SELECT * FROM `posts` WHERE `posts`.`resource` = 'news' ORDER BY created_at desc LIMIT 15 OFFSET 0"
	ListPagesQuery     = "SELECT * FROM `posts` WHERE `posts`.`resource` = '' ORDER BY created_at desc LIMIT 15 OFFSET 0"
	ListStatusQuery    = "SELECT * FROM `posts` WHERE `posts`.`status` = 'published' ORDER BY created_at desc LIMIT 15 OFFSET 0"
	CountQuery         = "SELECT COUNT(*) AS rowcount FROM (SELECT * FROM `posts` ORDER BY created_at desc"
	CountResourceQuery = "SELECT COUNT(*) AS rowcount FROM (SELECT * FROM `posts` WHERE `posts`.`resource` = 'news' ORDER BY created_at desc"
	CountPagesQuery    = "SELECT COUNT(*) AS rowcount FROM (SELECT * FROM `posts` WHERE `posts`.`resource` = '' ORDER BY created_at desc"
	CountStatusQuery   = "SELECT COUNT(*) AS rowcount FROM (SELECT * FROM `posts` WHERE `posts`.`status` = 'published' ORDER BY created_at desc"
)

func (t *PostsTestSuite) TestStore_List() {
	tt := map[string]struct {
		meta  params.Params
		mock  func(m sqlmock.Sqlmock)
		cfg   ListConfig
		total int
		want  interface{}
	}{
		"Success": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(posts[0].ID, posts[0].Slug, posts[0].Title).
					AddRow(posts[1].ID, posts[1].Slug, posts[1].Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(ListQuery))).
					WillReturnRows(rows)

				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).
					WillReturnRows(countRows)
			},
			ListConfig{},
			2,
			postData,
		},
		"Filter Error": {
			params.Params{
				Page:           api.DefaultParams.Page,
				Limit:          15,
				OrderBy:        api.DefaultParams.OrderBy,
				OrderDirection: api.DefaultParams.OrderDirection,
				Filters:        params.Filters{"wrong_column": {{Operator: "=", Value: "verbis"}}}},
			nil,
			ListConfig{},
			-1,
			"The wrong_column search query does not exist",
		},
		"Not Found": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"})
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(ListQuery))).
					WillReturnRows(rows)

				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).
					WillReturnRows(countRows)
			},
			ListConfig{},
			-1,
			"No posts available",
		},
		"Internal": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(ListQuery))).
					WillReturnError(fmt.Errorf("error"))
			},
			ListConfig{},
			-1,
			database.ErrQueryMessage,
		},
		"Count Error": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(posts[0].ID, posts[0].Slug, posts[0].Title).
					AddRow(posts[1].ID, posts[1].Slug, posts[1].Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(ListQuery))).
					WillReturnRows(rows)

				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
			ListConfig{},
			-1,
			"Error getting the total number of posts",
		},
		"Pages": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(posts[0].ID, posts[0].Slug, posts[0].Title).
					AddRow(posts[1].ID, posts[1].Slug, posts[1].Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(ListPagesQuery))).
					WillReturnRows(rows)

				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountPagesQuery)).
					WillReturnRows(countRows)
			},
			ListConfig{
				Resource: "pages",
			},
			2,
			postData,
		},
		"Resource": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(posts[0].ID, posts[0].Slug, posts[0].Title).
					AddRow(posts[1].ID, posts[1].Slug, posts[1].Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(ListResourceQuery))).
					WillReturnRows(rows)

				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountResourceQuery)).
					WillReturnRows(countRows)
			},
			ListConfig{
				Resource: "news",
			},
			2,
			postData,
		},
		"Status": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(posts[0].ID, posts[0].Slug, posts[0].Title).
					AddRow(posts[1].ID, posts[1].Slug, posts[1].Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(ListStatusQuery))).
					WillReturnRows(rows)

				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountStatusQuery)).
					WillReturnRows(countRows)
			},
			ListConfig{
				Status: "published",
			},
			2,
			postData,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, total, err := s.List(test.meta, false, test.cfg)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.total, total)
			t.RunT(test.want, got)
		})
	}
}
