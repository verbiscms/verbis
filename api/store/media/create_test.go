// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/store/media/sizes"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO `media` (`title`, `alt`, `description`, `user_id`, `file_id`, `updated_at`, `created_at`) VALUES ('', '', '', 0, 0, NOW(), NOW())"
)

func (t *MediaTestSuite) TestStore_Create() {
	tt := map[string]struct {
		want      interface{}
		mockSizes func(m *mocks.Repository)
		mock      func(m sqlmock.Sqlmock)
	}{
		"Success": {
			mediaItem,
			func(m *mocks.Repository) {
				m.On("Create", mediaItem.Id, mediaItem.Sizes).Return(mediaItem.Sizes, nil)
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(mediaItem.Id), 1))
			},
		},
		"No Rows": {
			"Error creating media item with the name",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Last Insert ID Error": {
			"Error getting the newly created media item ID",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("err")))
			},
		},
		"Sizes Error": {
			"error",
			func(m *mocks.Repository) {
				m.On("Create", mediaItem.Id, mediaItem.Sizes).Return(nil, fmt.Errorf("error"))
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(mediaItem.Id), 1))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock, test.mockSizes)
			cat, err := s.Create(mediaItem)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(cat, test.want)
		})
	}
}
