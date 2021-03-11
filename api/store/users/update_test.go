// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

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
	UpdateQuery = "UPDATE `users` SET `uuid` = ?, `first_name` = 'Verbis', `last_name` = 'CMS', `email` = 'verbis@verbiscms.com', `password` = ?, `website` = NULL, `facebook` = 'Verbis', `twitter` = NULL, `linked_in` = NULL, `instagram` = NULL, `biography` = NULL, `profile_picture_id` = NULL, `token` = ?, `updated_at` = NOW() WHERE `id` = '1'"
)

func (t *UsersTestSuite) TestStore_Update() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			user,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(user.Id), 1))

				m.ExpectExec(regexp.QuoteMeta(UpdatePivotQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"Validation Failed": {
			"Validation failed, choose another email address",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsByEmailQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"Error updating user with the name",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Error Pivot": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(user.Id), 1))

				m.ExpectExec(regexp.QuoteMeta(UpdatePivotQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			cat, err := s.Update(user)
			if err != nil {
				t.Contains(errors.Message(err), test.want, err)
				return
			}
			t.RunT(cat, test.want, 3)
		})
	}
}
