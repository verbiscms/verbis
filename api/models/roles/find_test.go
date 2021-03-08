// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

func (t *RolesTestSuite) TestStore_Find() {
	query := "SELECT * FROM `roles` WHERE `id` = ? LIMIT 1"

	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			role,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).
					AddRow(role.Id, role.Name, role.Description)
				m.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No redirect exists with the ID: 1",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			"Error executing sql query",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(role.Id)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
