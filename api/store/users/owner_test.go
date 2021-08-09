// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"bytes"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"github.com/verbiscms/verbis/api/logger"
	"regexp"
)

var (
	OwnerQuery = SelectStatement + "WHERE `roles`.`id` = '6' LIMIT 1"
)

func (t *UsersTestSuite) TestStore_Owner() {
	l := logrus.New()
	buf := bytes.Buffer{}
	l.ExitFunc = func(i int) {
		buf.WriteString("error")
	}
	l.SetOutput(&buf)
	logger.SetLogger(l)

	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
		err  bool
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
		"Internal Error": {
			"error",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(OwnerQuery)).WillReturnError(fmt.Errorf("error"))
			},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			defer buf.Reset()
			s := t.Setup(test.mock)
			got := s.Owner()
			if test.err {
				t.Equal(test.want, buf.String())
				return
			}
			t.RunT(test.want, got, 12)
		})
	}
}
