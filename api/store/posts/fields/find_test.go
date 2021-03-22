// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	FindQuery             = "SELECT * FROM `post_fields` WHERE `post_id` = '1' LIMIT 1"
	FindByPostAndKeyQuery = "SELECT * FROM `post_fields` WHERE `post_id` = '1' AND `key` = '" + field.Key + "' LIMIT 1"
)

func (t *PostFieldsTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			fields,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "post_id", "type", "name", "field_key", "value"}).
					AddRow(fields[0].Id, fields[0].PostId, fields[0].Type, fields[0].Name, fields[0].Key, fields[0].OriginalValue).
					AddRow(fields[1].Id, fields[1].PostId, fields[1].Type, fields[1].Name, fields[1].Key, fields[1].OriginalValue)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No fields exists with the post ID",
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
			got, err := s.Find(field.PostId)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}

func (t *PostFieldsTestSuite) TestStore_FindByPostAndKey() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			fields,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "post_id", "type", "name", "field_key", "value"}).
					AddRow(fields[0].Id, fields[0].PostId, fields[0].Type, fields[0].Name, fields[0].Key, fields[0].OriginalValue).
					AddRow(fields[1].Id, fields[1].PostId, fields[1].Type, fields[1].Name, fields[1].Key, fields[1].OriginalValue)
				m.ExpectQuery(regexp.QuoteMeta(FindByPostAndKeyQuery)).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No fields exists with the post ID and key",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByPostAndKeyQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByPostAndKeyQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.FindByPostAndKey(field.PostId, field.Key)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
