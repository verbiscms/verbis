// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/test/dummy"
	"regexp"
)

var (
	ListQuery          = "SELECT * FROM `categories` ORDER BY \"created_at\" desc LIMIT 15 OFFSET 0"
	CountQuery         = "SELECT COUNT(*) AS rowcount FROM (SELECT * FROM `categories` ORDER BY \"created_at\" desc"
	ListResourceQuery  = "SELECT * FROM `categories` WHERE `categories`.`resource` = 'news' ORDER BY \"created_at\" desc LIMIT 15 OFFSET 0"
	CountResourceQuery = "SELECT COUNT(*) AS rowcount FROM (SELECT * FROM `categories` WHERE `categories`.`resource` = 'news' ORDER BY \"created_at\" desc) AS rowdata"
)

func (t *CategoriesTestSuite) TestStore_List() {
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
				rows := sqlmock.NewRows([]string{"id", "slug", "name"}).
					AddRow(categories[0].Id, categories[0].Slug, categories[0].Name).
					AddRow(categories[1].Id, categories[1].Slug, categories[1].Name)
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnRows(rows)
				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).WillReturnRows(countRows)
			},
			ListConfig{},
			2,
			categories,
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
		"No Rows": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnError(sql.ErrNoRows)
			},
			ListConfig{},
			-1,
			"No categories available",
		},
		"Internal": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnError(fmt.Errorf("error"))
			},
			ListConfig{},
			-1,
			database.ErrQueryMessage,
		},
		"Count Error": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "name"}).
					AddRow(categories[0].Id, categories[0].Slug, categories[0].Name).
					AddRow(categories[1].Id, categories[1].Slug, categories[1].Name)
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnRows(rows)
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).WillReturnError(fmt.Errorf("error"))
			},
			ListConfig{},
			-1,
			"Error getting the total number of categories",
		},
		"Resource": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "name"}).
					AddRow(categories[0].Id, categories[0].Slug, categories[0].Name).
					AddRow(categories[1].Id, categories[1].Slug, categories[1].Name)
				m.ExpectQuery(regexp.QuoteMeta(ListResourceQuery)).
					WillReturnRows(rows)

				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountResourceQuery)).
					WillReturnRows(countRows)
			},
			ListConfig{
				Resource: "news",
			},
			2,
			categories,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, total, err := s.List(test.meta, test.cfg)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.total, total)
			t.RunT(test.want, got)
		})
	}
}
