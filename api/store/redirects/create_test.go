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
	CreateQuery = "INSERT INTO `redirects` (`from_path`, `to_path`, `code`, `updated_at`, `created_at`) VALUES ('/from', '/to', 301, NOW(), NOW())"
)

func (t *RedirectsTestSuite) TestStore_Create() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			redirect,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(redirect.ID), 1))
			},
		},
		"Validation Failed": {
			"Validation failed, the redirect from path already exists",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsByFromQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"Error creating redirect with the from path",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Last Insert ID Error": {
			"Error getting the newly created redirect ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("err")))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			cat, err := s.Create(redirect)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(cat, test.want, 2)
		})
	}
}
