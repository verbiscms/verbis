// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/store/media/sizes"
	"github.com/verbiscms/verbis/api/test/dummy"
	"regexp"
)

var (
	ListQuery  = SelectStatement + "ORDER BY created_at desc LIMIT 15 OFFSET 0"
	CountQuery = "SELECT COUNT(*) AS rowcount FROM (" + SelectStatement + " ORDER BY created_at desc) AS rowdata"
)

func (t *MediaTestSuite) TestStore_List() {
	tt := map[string]struct {
		meta      params.Params
		mockSizes func(m *mocks.Repository)
		mock      func(m sqlmock.Sqlmock)
		total     int
		want      interface{}
	}{
		"Success": {
			dummy.DefaultParams,
			func(m *mocks.Repository) {
				m.On("Find", mediaItem.ID).Return(mediaItem.Sizes, nil)
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "file.name", "title"}).
					AddRow(mediaItems[0].ID, mediaItems[0].File.Name, mediaItems[0].Title).
					AddRow(mediaItems[1].ID, mediaItems[1].File.Name, mediaItems[1].Title)
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnRows(rows)
				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).WillReturnRows(countRows)
			},
			2,
			mediaItems,
		},
		"Filter Error": {
			params.Params{
				Page:           api.DefaultParams.Page,
				Limit:          15,
				OrderBy:        api.DefaultParams.OrderBy,
				OrderDirection: api.DefaultParams.OrderDirection,
				Filters:        params.Filters{"wrong_column": {{Operator: "=", Value: "verbis"}}}},
			nil,
			nil,
			-1,
			"The wrong_column search query does not exist",
		},
		"No Rows": {
			dummy.DefaultParams,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnError(sql.ErrNoRows)
			},
			-1,
			"No media items available",
		},
		"Internal": {
			dummy.DefaultParams,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnError(fmt.Errorf("error"))
			},
			-1,
			database.ErrQueryMessage,
		},
		"Count Error": {
			dummy.DefaultParams,
			nil,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "file.name", "title"}).
					AddRow(mediaItems[0].ID, mediaItems[0].File.Name, mediaItems[0].Title).
					AddRow(mediaItems[1].ID, mediaItems[1].File.Name, mediaItems[1].Title)
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnRows(rows)
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).WillReturnError(fmt.Errorf("error"))
			},
			-1,
			"Error getting the total number of media items",
		},
		"Sizes Error": {
			dummy.DefaultParams,
			func(m *mocks.Repository) {
				m.On("Find", mediaItem.ID).Return(nil, fmt.Errorf("error"))
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "file.name", "title"}).
					AddRow(mediaItems[0].ID, mediaItems[0].File.Name, mediaItems[0].Title).
					AddRow(mediaItems[1].ID, mediaItems[1].File.Name, mediaItems[1].Title)
				m.ExpectQuery(regexp.QuoteMeta(ListQuery)).WillReturnRows(rows)
				countRows := sqlmock.NewRows([]string{"rowdata"}).AddRow("2")
				m.ExpectQuery(regexp.QuoteMeta(CountQuery)).WillReturnRows(countRows)
			},
			2,
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock, test.mockSizes)
			got, total, err := s.List(test.meta)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.total, total)
			t.RunT(test.want, got)
		})
	}
}
