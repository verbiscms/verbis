// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"regexp"
)

var (
	FindQuery = "SELECT * FROM `form_fields` WHERE `form_id` = '" + formID + "'"
)

func (t *FieldsTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			formFields,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"key", "label", "type"}).
					AddRow(formFields[0].Key, formFields[0].Label, formFields[0].Type)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No form fields exists with the form ID",
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
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(form.ID)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
