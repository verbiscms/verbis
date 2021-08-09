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
	DeleteQuery = "DELETE FROM `media` WHERE `id` = '" + mediaID + "'"
)

func (t *MediaTestSuite) TestStore_Delete() {
	tt := map[string]struct {
		want      interface{}
		mockSizes func(m *mocks.Repository)
		mock      func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(m *mocks.Repository) {
				m.On("Delete", mediaItem.ID).Return(nil)
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"No Rows": {
			"No media item exists with the ID",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Sizes Error": {
			"error",
			func(m *mocks.Repository) {
				m.On("Delete", mediaItem.ID).Return(fmt.Errorf("error"))
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock, test.mockSizes)
			err := s.Delete(mediaItem.ID)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
