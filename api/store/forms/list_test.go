// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	fields "github.com/verbiscms/verbis/api/mocks/store/forms/fields"
	submissions "github.com/verbiscms/verbis/api/mocks/store/forms/submissions"
	"github.com/verbiscms/verbis/api/test/dummy"
	"regexp"
)

var (
	ListQuery  = "SELECT * FROM `forms` ORDER BY created_at desc LIMIT 15 OFFSET 0"
	CountQuery = "SELECT COUNT(*) AS rowcount FROM (SELECT * FROM `forms` ORDER BY created_at desc"
)

func (t *FormsTestSuite) TestStore_List() {
	tt := map[string]struct {
		meta      params.Params
		mockForms func(f *fields.Repository, s *submissions.Repository)
		mock      func(m sqlmock.Sqlmock)
		total     int
		want      interface{}
	}{
		"Success": {
			dummy.DefaultParams,
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Find", forms[0].ID).Return(forms[0].Fields, nil)
				f.On("Find", forms[1].ID).Return(forms[1].Fields, nil)
				s.On("Find", forms[0].ID).Return(forms[0].Submissions, nil)
				s.On("Find", forms[1].ID).Return(forms[1].Submissions, nil)
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(forms[0].ID, forms[0].Name).
					AddRow(forms[1].ID, forms[1].Name)
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnRows(rows)
				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).WillReturnRows(countRows)
			},
			2,
			forms,
		},
		"Filter Error": {
			params.Params{
				Page:           api.DefaultParams.Page,
				Limit:          15,
				OrderBy:        api.DefaultParams.OrderBy,
				OrderDirection: api.DefaultParams.OrderDirection,
				Filters:        params.Filters{"wrong_column": {{Operator: "=", Value: "verbis"}}}},
			nil,
			nil,
			-1,
			"The wrong_column search query does not exist",
		},
		"No Rows": {
			dummy.DefaultParams,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnError(sql.ErrNoRows)
			},
			-1,
			"No forms available",
		},
		"Internal": {
			dummy.DefaultParams,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnError(fmt.Errorf("error"))
			},
			-1,
			database.ErrQueryMessage,
		},
		"Count Error": {
			dummy.DefaultParams,
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(forms[0].ID, forms[0].Name).
					AddRow(forms[1].ID, forms[1].Name)
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnRows(rows)
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).WillReturnError(fmt.Errorf("error"))
			},
			-1,
			"Error getting the total number of forms",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock, test.mockForms)
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
