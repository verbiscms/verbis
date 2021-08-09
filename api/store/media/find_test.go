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
	FindQuery = SelectStatement + "WHERE `media`.`id` = '" + mediaID + "' LIMIT 1"
)

func (t *MediaTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want      interface{}
		mockSizes func(*mocks.Repository)
		mock      func(m sqlmock.Sqlmock)
	}{
		"Success": {
			mediaItem,
			func(m *mocks.Repository) {
				m.On("Find", mediaItem.ID).Return(mediaItem.Sizes, nil)
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "file.name", "title"}).
					AddRow(mediaItem.ID, mediaItem.File.Name, mediaItem.Title)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No media item exists with the ID",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Sizes Error": {
			"error",
			func(m *mocks.Repository) {
				m.On("Find", mediaItem.ID).Return(nil, fmt.Errorf("error"))
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "file.name", "title"}).
					AddRow(mediaItem.ID, mediaItem.File.Name, mediaItem.Title)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnRows(rows)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock, test.mockSizes)
			got, err := s.Find(mediaItem.ID)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
