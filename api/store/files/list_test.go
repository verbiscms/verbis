// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"database/sql"
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
	ListQuery  = "SELECT * FROM `files` ORDER BY created_at desc LIMIT 15 OFFSET 0"
	CountQuery = "SELECT COUNT(*) AS rowcount FROM (SELECT * FROM `files` ORDER BY created_at desc"
)

func (t *FilesTestSuite) TestStore_List() {
	tt := map[string]struct {
		meta  params.Params
		mock  func(m sqlmock.Sqlmock)
		total int
		want  interface{}
	}{
		"Success": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "url", "name", "bucket_id", "provider"}).
					AddRow(files[0].ID, files[0].URL, files[0].Name, files[0].BucketID, files[0].Provider).
					AddRow(files[1].ID, files[1].URL, files[1].Name, files[1].BucketID, files[1].Provider)
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnRows(rows)
				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).WillReturnRows(countRows)
			},
			2,
			files,
		},
		"Filter Error": {
			params.Params{
				Page:           api.DefaultParams.Page,
				Limit:          15,
				OrderBy:        api.DefaultParams.OrderBy,
				OrderDirection: api.DefaultParams.OrderDirection,
				Filters:        params.Filters{"wrong_column": {{Operator: "=", Value: "verbis"}}}},
			nil,
			-1,
			"The wrong_column search query does not exist",
		},
		"No Rows": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).
					WillReturnError(sql.ErrNoRows)
			},
			-1,
			"No files available",
		},
		"Internal": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
			-1,
			database.ErrQueryMessage,
		},
		"Count Error": {
			dummy.DefaultParams,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "url", "name", "bucket_id", "provider"}).
					AddRow(files[0].ID, files[0].URL, files[0].Name, files[0].BucketID, files[0].Provider).
					AddRow(files[1].ID, files[1].URL, files[1].Name, files[1].BucketID, files[1].Provider)
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).
					WillReturnRows(rows)
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
			-1,
			"Error getting the total number of files",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, total, err := s.List(test.meta)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.total, total)
			t.RunT(test.want, got)
		})
	}
}
