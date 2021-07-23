// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sizes

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO `media_sizes` (`file_id`, `media_id`, `size_name`, `size_key`, `width`, `height`, `crop`) VALUES (0, 1, 'gopher-1920x1080.jpg', 'hd', 0, 0, FALSE)"
)

func (t *SizesTestSuite) TestStore_Create() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			sizes,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"No Rows": {
			"Error creating media size with the key",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Last Insert ID Error": {
			"Error getting the newly created media size ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("error")))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			media, err := s.Create(1, sizes)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(media, test.want)
		})
	}
}
