// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
)

func (t *OptionsTestSuite) TestStore_Insert() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Update": {
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"option_name"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"Update Error": {
			"error",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"option_name"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Create": {
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"option_name"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"Create Error": {
			"error",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"option_name"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Insert(options)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
