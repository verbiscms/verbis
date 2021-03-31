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
	"regexp"
)

var (
	FindQuery           = "SELECT * FROM `media` WHERE `id` = '" + mediaID + "' LIMIT 1"
	FindByNameQuery     = "SELECT * FROM `media` WHERE `file_name` = '" + mediaItem.FileName + "' LIMIT 1"
	FindByURLQuery      = "SELECT * FROM `media` WHERE `url` = '" + mediaItemURL.Url + "' LIMIT 1"
	FindByURLSizeQuery  = "SELECT * FROM `media` WHERE `url` = '" + mediaItemSizes.Sizes["test"].Url + "' LIMIT 1"
	FindByURLSizesQuery = "SELECT * FROM `media` WHERE sizes LIKE '%" + mediaItemSizes.Sizes["test"].Url + "%' LIMIT 1"
)

func (t *MediaTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			mediaItem,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "url", "title", "file_name"}).
					AddRow(mediaItem.Id, mediaItem.Url, mediaItem.Title, mediaItem.FileName)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No media item exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(mediaItem.Id)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}

func (t *MediaTestSuite) TestStore_FindByName() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			mediaItem,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "url", "title", "file_name"}).
					AddRow(mediaItem.Id, mediaItem.Url, mediaItem.Title, mediaItem.FileName)
				m.ExpectQuery(regexp.QuoteMeta(FindByNameQuery)).WithArgs().WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No media item exists with the name: ",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByNameQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByNameQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.FindByName(mediaItem.FileName)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}

func (t *MediaTestSuite) TestStore_FindByURL() {
	tt := map[string]struct {
		url  string
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Normal": {
			mediaItemURL.Url,
			mediaItemURL,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "url"}).
					AddRow(mediaItemURL.Id, mediaItemURL.Url)
				m.ExpectQuery(regexp.QuoteMeta(FindByURLQuery)).WithArgs().WillReturnRows(rows)
			},
		},
		"Sizes": {
			mediaItemSizes.Sizes["test"].Url,
			mediaItemSizes,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByURLSizeQuery)).
					WillReturnError(fmt.Errorf("error"))

				rows := sqlmock.NewRows([]string{"id", "url", "sizes"}).
					AddRow(mediaItemSizes.Id, mediaItemSizes.Url, mediaItemSizes.Sizes)
				m.ExpectQuery(regexp.QuoteMeta(FindByURLSizesQuery)).WithArgs().WillReturnRows(rows)
			},
		},
		"Internal Error": {
			mediaItemURL.Url,
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByURLQuery)).
					WillReturnError(fmt.Errorf("error"))

				m.ExpectQuery(regexp.QuoteMeta(FindByURLSizesQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Not Found": {
			mediaItemSizes.Sizes["test"].Url,
			"Error getting media item with the url",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByURLSizeQuery)).
					WillReturnError(fmt.Errorf("error"))

				rows := sqlmock.NewRows([]string{"id", "url"}).
					AddRow(mediaItemSizes.Id, mediaItemSizes.Url)
				m.ExpectQuery(regexp.QuoteMeta(FindByURLSizesQuery)).WithArgs().WillReturnRows(rows)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, _, err := s.FindByURL(test.url)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
