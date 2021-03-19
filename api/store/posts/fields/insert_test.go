// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/test"
	"regexp"
)

func (t *PostFieldsTestSuite) TestStore_Insert() {
	tt := map[string]struct {
		want   interface{}
		fields domain.PostFields
		mock   func(m sqlmock.Sqlmock)
	}{
		"Find Error": {
			database.ErrQueryMessage,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Delete Error": {
			database.ErrQueryMessage,
			fields,
			func(m sqlmock.Sqlmock) {
				// Find
				rows := sqlmock.NewRows([]string{"id", "name", "field_key"}).
					AddRow(fields[0].Id, "WRONG", "WRONG")
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnRows(rows)

				// Delete
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Update": {
			nil,
			fields,
			func(m sqlmock.Sqlmock) {
				// Find
				rows := sqlmock.NewRows([]string{"id", "post_id", "type", "name", "field_key", "value"}).
					AddRow(fields[0].Id, fields[0].PostId, fields[0].Type, fields[0].Name, fields[0].Key, fields[0].OriginalValue)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnRows(rows)

				// Exists
				rowsE := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).WillReturnRows(rowsE)

				// Update
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(field.Id), 1))
			},
		},
		"Update Error": {
			database.ErrQueryMessage,
			fields,
			func(m sqlmock.Sqlmock) {
				// Find
				rows := sqlmock.NewRows([]string{"id", "post_id", "type", "name", "field_key", "value"}).
					AddRow(fields[0].Id, fields[0].PostId, fields[0].Type, fields[0].Name, fields[0].Key, fields[0].OriginalValue)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnRows(rows)

				// Exists
				rowsE := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).WillReturnRows(rowsE)

				// Update
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Insert": {
			nil,
			fields,
			func(m sqlmock.Sqlmock) {
				// Find
				rows := sqlmock.NewRows([]string{"id", "post_id", "type", "name", "field_key", "value"}).
					AddRow(fields[0].Id, fields[0].PostId, fields[0].Type, fields[0].Name, fields[0].Key, fields[0].OriginalValue)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnRows(rows)

				// Exists
				rowsE := sqlmock.NewRows([]string{"id"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).WillReturnRows(rowsE)

				// Create
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(field.Id), 1))
			},
		},
		"Insert Error": {
			database.ErrQueryMessage,
			fields,
			func(m sqlmock.Sqlmock) {
				// Find
				rows := sqlmock.NewRows([]string{"id", "post_id", "type", "name", "field_key", "value"}).
					AddRow(fields[0].Id, fields[0].PostId, fields[0].Type, fields[0].Name, fields[0].Key, fields[0].OriginalValue)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnRows(rows)

				// Exists
				rowsE := sqlmock.NewRows([]string{"id"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).WillReturnRows(rowsE)

				// Create
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Insert(field.PostId, test.fields)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
