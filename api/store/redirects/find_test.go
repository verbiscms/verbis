// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"regexp"
)

var (
	FindQuery       = "SELECT * FROM `redirects` WHERE `id` = '" + redirectID + "' LIMIT 1"
	FindByFromQuery = "SELECT * FROM `redirects` WHERE `from_path` = '" + redirect.From + "' LIMIT 1"
)

func (t *RedirectsTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			redirect,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "from_path", "to_path", "code"}).
					AddRow(redirect.Id, redirect.From, redirect.To, redirect.Code)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No redirect exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(redirect.Id)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}

func (t *RedirectsTestSuite) TestStore_FindByFrom() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			redirect,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "from_path", "to_path", "code"}).
					AddRow(redirect.Id, redirect.From, redirect.To, redirect.Code)
				m.ExpectQuery(regexp.QuoteMeta(FindByFromQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No redirect exists with the from path",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByFromQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByFromQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.FindByFrom(redirect.From)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
