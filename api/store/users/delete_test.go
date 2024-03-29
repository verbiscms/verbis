// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"regexp"
)

var (
	DeleteQuery = "DELETE FROM `users` WHERE `id` = '" + userID + "'"
)

func (t *UsersTestSuite) TestStore_Delete() {
	tt := map[string]struct {
		id   int
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			user.ID,
			user,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))

				m.ExpectExec(regexp.QuoteMeta(DeletePivotQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"Owner": {
			domain.OwnerRoleID,
			"The owner of the site cannot be deleted",
			nil,
		},
		"No Rows": {
			user.ID,
			"No user exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			user.ID,
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Error Pivot": {
			user.ID,
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))

				m.ExpectExec(regexp.QuoteMeta(DeletePivotQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Delete(test.id)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
