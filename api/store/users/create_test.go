// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/test"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO `users` (`uuid`, `first_name`, `last_name`, `email`, `password`, `website`, `facebook`, `twitter`, `linked_in`, `instagram`, `biography`, `profile_picture_id`, `token`, `updated_at`, `created_at`) VALUES (?, 'Verbis', 'CMS', 'verbis@verbiscms.com', ?, '', 'Verbis', '', '', '', '', NULL, ?, NOW(), NOW())"
)

func (t *UsersTestSuite) TestStore_Create() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			userCreate.User,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}, test.DBAnyString{}, test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(user.Id), 1))

				m.ExpectExec(regexp.QuoteMeta(CreatePivotQuery)).
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
			"Error creating user with the name",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}, test.DBAnyString{}, test.DBAnyString{}).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}, test.DBAnyString{}, test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Last Insert ID Error": {
			"Error getting the newly created user ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}, test.DBAnyString{}, test.DBAnyString{}).
					WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("err")))
			},
		},
		"Error Pivot": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}, test.DBAnyString{}, test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(user.Id), 1))

				m.ExpectExec(regexp.QuoteMeta(CreatePivotQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			cat, err := s.Create(userCreate)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(cat, test.want, 3)
		})
	}
}

func (t *UsersTestSuite) TestStore_Create_FailedHash() {
	s := t.Setup(nil)
	s.hashPasswordFunc = func(password string) (string, error) {
		return "", fmt.Errorf("error")
	}
	_, err := s.Create(userCreate)
	want := "error"
	t.Equal(want, err.Error())
}
