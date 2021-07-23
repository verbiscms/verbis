// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

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
	CreateQuery = "INSERT INTO `post_fields` (`uuid`, `post_id`, `type`, `name`, `value`, `field_key`) VALUES (?, 1, 'text', 'name', 'val', 'key')"
)

func (t *FieldsTestSuite) TestStore_Create() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			field,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"No Rows": {
			"Error creating field with the name",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			cat, err := s.create(field)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(cat, test.want)
		})
	}
}
