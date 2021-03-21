// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/test"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO `forms` (`uuid`, `name`, `email_send`, `email_message`, `email_subject`, `store_db`, `updated_at`, `created_at`) VALUES (?, 'Form', FALSE, '', '', FALSE, NOW(), NOW())"
)

func (t *FormsTestSuite) TestStore_Create() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			form,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(form.Id), 1))
			},
		},
		"No Rows": {
			"Error creating form with the name",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Last Insert ID Error": {
			"Error getting the newly created form ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("err")))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			cat, err := s.Create(form)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(cat, test.want)
		})
	}
}
