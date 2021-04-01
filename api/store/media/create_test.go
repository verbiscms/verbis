// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/test"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO `media` (`uuid`, `url`, `title`, `alt`, `description`, `file_path`, `file_size`, `file_name`, `sizes`, `mime`, `user_id`, `updated_at`, `created_at`) VALUES ('00000000-0000-0000-0000-000000000000', '', '', '', '', '', 0, 'gopher.png', ?, '', 0, NOW(), NOW())"
)

func (t *MediaTestSuite) TestStore_Create() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			mediaItem,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAny{}).
					WillReturnResult(sqlmock.NewResult(int64(mediaItem.Id), 1))
			},
		},
		"No Rows": {
			"Error creating category with the name",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAny{}).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAny{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Last Insert ID Error": {
			"Error getting the newly created category ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAny{}).
					WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("err")))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			cat, err := s.Create(mediaItem)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(cat, test.want)
		})
	}
}
