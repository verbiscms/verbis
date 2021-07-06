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
	FindQuery      = "SELECT * FROM `files` WHERE `id` = '" + fileID + "' LIMIT 1"
	FindByURLQuery = "SELECT * FROM `files` WHERE `url` = '" + file.URL + "' LIMIT 1"
)

func (t *FilesTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			file,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "url", "name", "path", "provider"}).
					AddRow(file.Id, file.URL, file.Name, file.Path, file.Provider)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No file exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(file.Id)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}

func (t *FilesTestSuite) TestStore_FindByURL() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			file,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "url", "name", "path", "provider"}).
					AddRow(file.Id, file.URL, file.Name, file.Path, file.Provider)
				m.ExpectQuery(regexp.QuoteMeta(FindByURLQuery)).
					WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No file exists with the URL",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByURLQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByURLQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.FindByURL(file.URL)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
