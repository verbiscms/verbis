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
	"regexp"
)

var (
	FindQuery        = SelectStatement + "WHERE `users`.`id` = '" + userID + "' LIMIT 1"
	FindByTokenQuery = SelectStatement + "WHERE `users`.`token` = '" + user.Token + "' LIMIT 1"
	FindByEmailQuery = SelectStatement + "WHERE `users`.`email` = '" + user.Email + "' LIMIT 1"
)

func (t *UsersTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			user,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "token", "roles.name"}).
					AddRow(user.ID, user.FirstName, user.LastName, user.Email, user.Token, user.Role.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No category exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(user.ID)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got, 12)
		})
	}
}

func (t *UsersTestSuite) TestStore_FindByToken() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			user,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "token", "roles.name"}).
					AddRow(user.ID, user.FirstName, user.LastName, user.Email, user.Token, user.Role.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindByTokenQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No user exists with the token",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByTokenQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByTokenQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.FindByToken(user.Token)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got, 12)
		})
	}
}

func (t *UsersTestSuite) TestStore_FindByEmail() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			user,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "token", "roles.name"}).
					AddRow(user.ID, user.FirstName, user.LastName, user.Email, user.Token, user.Role.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindByEmailQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No user exists with the email",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByEmailQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByEmailQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.FindByEmail(user.Email)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got, 12)
		})
	}
}
