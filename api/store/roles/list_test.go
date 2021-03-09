// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	ListQuery = "SELECT * FROM `roles` ORDER BY \"id\" desc"
)

func (t *RolesTestSuite) TestStore_List() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			roles,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).
					AddRow(roles[0].Id, roles[0].Name, roles[0].Description).
					AddRow(roles[1].Id, roles[1].Name, roles[1].Description)
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No roles available",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.List()
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
