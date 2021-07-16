// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	UpdateQuery = "UPDATE `files` SET `provider` = 'provider', `bucket` = 'bucket', `region` = '' WHERE `id` = '1'"
)

func (t *FilesTestSuite) TestStore_Update() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(file.Id), 1))
			},
		},
		"No Rows": {
			"Error updating file with the ID: 1",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Update(1, storageChange)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(err, test.want)
		})
	}
}
