// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postfields

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	DeleteQuery = "DELETE FROM `post_fields` WHERE `uuid` = '00000000-0000-0000-0000-000000000000' AND `post_id` = '1' AND `field_key` = 'key' AND `name` = 'name'"
)

func (t *PostFieldsTestSuite) TestStore_Delete() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"No Rows": {
			"No field exists",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.delete(field.PostId, field)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Nil(err)
		})
	}
}
