// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
)

func (t *PostCategoriesTestSuite) TestStore_Insert() {
	tt := map[string]struct {
		input *int
		want  interface{}
		mock  func(m sqlmock.Sqlmock)
	}{
		"Update": {
			&post.Category.Id,
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"Update Error": {
			&post.Category.Id,
			"error",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Delete": {
			nil,
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"Delete Error": {
			nil,
			"error",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Create": {
			&post.Category.Id,
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"Create Error": {
			&post.Category.Id,
			"error",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Nil Category": {
			nil,
			"error",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Insert(post.Id, test.input)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
