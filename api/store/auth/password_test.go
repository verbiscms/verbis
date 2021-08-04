// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/errors"
	"regexp"
)

var (
	UpdateQuery = "UPDATE `users` SET `password` = ? WHERE `email` = 'verbis@verbiscms.com'"
)

func (t *AuthTestSuite) TestStore_ResetPassword() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
		hash func(password string) (string, error)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
			func(password string) (string, error) {
				return "password", nil
			},
		},
		"Hash": {
			"error",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
			func(password string) (string, error) {
				return "", fmt.Errorf("error")
			},
		},
		"DB Error": {
			"Error updating users table with the new password",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
			func(password string) (string, error) {
				return "password", nil
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			s.hashPasswordFunc = test.hash
			err := s.ResetPassword("verbis@verbiscms.com", "password")
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
