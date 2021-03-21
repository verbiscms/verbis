// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	FindQuery       = "SELECT * FROM `forms` WHERE `id` = '" + formID + "' LIMIT 1"
	FindByUUIDQuery = "SELECT * FROM `forms` WHERE `uuid` = '" + form.UUID.String() + "' LIMIT 1"
)

func (t *FormsTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			form,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(form.Id, form.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No form exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Fields": {
			formsTestFields[0],
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(form.Id, form.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnRows(rows)

				fieldRows := sqlmock.NewRows([]string{"key", "label", "type"}).
					AddRow(formFields[0].Key, formFields[0].Label, formFields[0].Type)
				m.ExpectQuery(regexp.QuoteMeta(FieldsQuery)).
					WillReturnRows(fieldRows)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(form.Id)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}

func (t *FormsTestSuite) TestStore_FindByUUID() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			form,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(form.Id, form.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindByUUIDQuery)).
					WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No form exists with the UUID",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByUUIDQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByUUIDQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Fields": {
			formsTestFields[0],
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(form.Id, form.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindByUUIDQuery)).
					WillReturnRows(rows)

				fieldRows := sqlmock.NewRows([]string{"key", "label", "type"}).
					AddRow(formFields[0].Key, formFields[0].Label, formFields[0].Type)
				m.ExpectQuery(regexp.QuoteMeta(FieldsQuery)).
					WillReturnRows(fieldRows)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.FindByUUID(form.UUID)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
