// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
)

var (
	OwnerQuery = SelectStatement + "WHERE `roles`.`id` = '6' LIMIT 1"
)

func (t *UsersTestSuite) TestStore_Owner() {
	tt := map[string]struct {
		want   interface{}
		mock   func(m sqlmock.Sqlmock)
		panics bool
	}{
		"Success": {
			user,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "token", "roles.name"}).
					AddRow(user.ID, user.FirstName, user.LastName, user.Email, user.Token, user.Role.Name)
				m.ExpectQuery(regexp.QuoteMeta(OwnerQuery)).WillReturnRows(rows)
			},
			false,
		},
		"No Rows": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(OwnerQuery)).WillReturnError(sql.ErrNoRows)
			},
			true,
		},
		"Internal Error": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(OwnerQuery)).WillReturnError(fmt.Errorf("error"))
			},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			if test.panics {
				t.Panics(func() { s.Owner() })
				return
			}
			got := s.Owner()
			t.RunT(test.want, got, 12)
		})
	}
}
