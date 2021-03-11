// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/test"
	"regexp"
	"time"
)

var (
	UpdateTokenQuery     = "UPDATE `users` SET `token` = ?, `updated_at` = NOW() WHERE `token` = 'token'"
	UpdateTokenUsedQuery = "UPDATE `users` SET `token_last_used` = NOW() WHERE `token` = 'token'"
	RestPasswordQuery    = "UPDATE `users` SET `password` = ? WHERE `id` = '1'"
)

func (t *UsersTestSuite) TestStore_CheckSession() {
	now := time.Now()
	old := time.Now().AddDate(0, -1, 0)
	u := domain.User{
		UserPart:      user.UserPart,
		TokenLastUsed: &now,
		Token:         user.Token,
	}

	tt := map[string]struct {
		session int
		want    interface{}
		mock    func(m sqlmock.Sqlmock)
	}{
		"Not Found": {
			1000,
			"No user exists with the token",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByTokenQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Nil Token Last Used": {
			1000,
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "token", "roles.name"}).
					AddRow(user.Id, user.FirstName, user.LastName, user.Email, user.Token, user.Role.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindByTokenQuery)).WillReturnRows(rows)
			},
		},
		"Update Token": {
			1,
			"Session expired, please login again",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "token_last_used", "token"}).
					AddRow(user.Id, &old, "token")
				m.ExpectQuery(regexp.QuoteMeta(FindByTokenQuery)).WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(UpdateTokenQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"Update Token Error": {
			1,
			"Error updating the user's token with the name",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "token_last_used", "token"}).
					AddRow(user.Id, &old, "token")
				m.ExpectQuery(regexp.QuoteMeta(FindByTokenQuery)).WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(UpdateTokenQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Update Session": {
			1000,
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "token_last_used", "token"}).
					AddRow(user.Id, &now, "token")
				m.ExpectQuery(regexp.QuoteMeta(FindByTokenQuery)).WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(UpdateTokenUsedQuery)).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"Update Session Error": {
			1000,
			"Error updating the user last token used column",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "token_last_used", "token"}).
					AddRow(user.Id, &now, "token")
				m.ExpectQuery(regexp.QuoteMeta(FindByTokenQuery)).WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(UpdateTokenUsedQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupSession(test.session, test.mock)
			err := s.CheckSession(u.Token)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}

func (t *UsersTestSuite) TestStore_ResetPassword() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			userCreate.User,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(RestPasswordQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(user.Id), 1))
			},
		},
		"No Rows": {
			"Error updating user password",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(RestPasswordQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(RestPasswordQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.ResetPassword(user.Id, passwordReset)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}

func (t *UsersTestSuite) TestStore_ResetPassword_FailedHash() {
	s := t.Setup(nil)
	s.hashPasswordFunc = func(password string) (string, error) {
		return "", fmt.Errorf("error")
	}
	err := s.ResetPassword(user.Id, passwordReset)
	want := "error"
	t.Equal(want, err.Error())
}

func (t *UsersTestSuite) TestStore_UpdateToken() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateTokenUsedQuery)).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"Internal Error": {
			"Error updating the user last token used column",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateTokenUsedQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.UpdateToken(user.Token)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
