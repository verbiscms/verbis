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
	FindQuery       = SelectStatement + " WHERE `id` = '" + mediaID + "' LIMIT 1"
	FindByNameQuery = SelectStatement + " WHERE `name` = 'name' LIMIT 1"
)

func (t *MediaTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			mediaItem,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "url", "title"}).
					AddRow(mediaItem.Id, mediaItem.Url, mediaItem.Title)
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
				rows := sqlmock.NewRows([]string{"id", "url", "title"}).
					AddRow(mediaItem.Id, mediaItem.Url, mediaItem.Title)
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
			got, err := s.FindByName("name")
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
