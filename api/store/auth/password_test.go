// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

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
	UpdateQuery = "UPDATE `users` SET `password` = ? WHERE `email` = 'verbis@verbiscms.com'"
	VerifyQuery = "SELECT * FROM `password_resets` WHERE `token` = 'token' LIMIT 1"
	DeleteQuery = "DELETE FROM `password_resets` WHERE `token` = 'token'"
	CleanQuery  = "DELETE FROM `password_resets` WHERE created_at < (NOW() - INTERVAL 2 HOUR)"
)

func (t *AuthTestSuite) TestStore_ResetPassword() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "email", "token"}).
					AddRow(passwordReset.Id, passwordReset.Email, passwordReset.Token)
				m.ExpectQuery(regexp.QuoteMeta(VerifyQuery)).WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))

				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"Verify Error": {
			"No user exists with the token: token",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(VerifyQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Update Error": {
			"Error updating the users table with the new password",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "email", "token"}).
					AddRow(passwordReset.Id, passwordReset.Email, passwordReset.Token)
				m.ExpectQuery(regexp.QuoteMeta(VerifyQuery)).WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Delete Error": {
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "email", "token"}).
					AddRow(passwordReset.Id, passwordReset.Email, passwordReset.Token)
				m.ExpectQuery(regexp.QuoteMeta(VerifyQuery)).WillReturnRows(rows)

				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))

				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.ResetPassword("token", "password")
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}

func (t *AuthTestSuite) TestStore_ResetPassword_FailedHash() {
	fn := func(m sqlmock.Sqlmock) {
		rows := sqlmock.NewRows([]string{"id", "email", "token"}).
			AddRow(passwordReset.Id, passwordReset.Email, passwordReset.Token)
		m.ExpectQuery(regexp.QuoteMeta(VerifyQuery)).WillReturnRows(rows)
	}

	s := t.Setup(fn)
	s.hashPasswordFunc = func(password string) (string, error) {
		return "", fmt.Errorf("error")
	}
	err := s.ResetPassword("token", "password")
	want := "error"
	t.Equal(want, err.Error())
}

func (t *AuthTestSuite) TestStore_SendResetPassword() {
	t.T().Skip("Come back to when mailer is done, mock driver.")
}

func (t *AuthTestSuite) TestStore_VerifyPassword() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			passwordReset,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "email", "token"}).
					AddRow(passwordReset.Id, passwordReset.Email, passwordReset.Token)
				m.ExpectQuery(regexp.QuoteMeta(VerifyQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No user exists with the token",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(VerifyQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(VerifyQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			pr, err := s.VerifyPasswordToken(user.Token)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, pr)
		})
	}
}

func (t *AuthTestSuite) TestStore_CleanResets() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CleanQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"Internal Error": {
			"Error deleting from the reset passwords table",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CleanQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.CleanPasswordResets()
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
