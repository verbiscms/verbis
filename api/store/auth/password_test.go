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
	"regexp"
)

var (
	UpdateQuery = "UPDATE users SET password = ? WHERE email = ?"
	InsertQuery = "INSERT INTO password_resets (email, token, created_at) VALUES (?, ?, NOW())"
	VerifyQuery = "SELECT * FROM `password_resets` WHERE `token` = 'token' LIMIT 1"
	CleanQuery  = "DELETE FROM `password_resets` WHERE created_at < (NOW() - INTERVAL 2 HOUR"
)

func (t *AuthTestSuite) TestStore_ResetPassword() {

}

func (t *AuthTestSuite) TestStore_SendResetPassword() {

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
