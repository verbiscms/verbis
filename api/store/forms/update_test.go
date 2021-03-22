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
	fields "github.com/ainsleyclark/verbis/api/mocks/store/forms/fields"
	submissions "github.com/ainsleyclark/verbis/api/mocks/store/forms/submissions"
	"github.com/ainsleyclark/verbis/api/test"
	"regexp"
)

var (
	UpdateQuery = "UPDATE `forms` SET `uuid` = ?, `name` = 'Form', `email_send` = FALSE, `email_message` = '', `email_subject` = '', `store_db` = FALSE, `updated_at` = NOW() WHERE `id` = '1'"
)

func (t *FormsTestSuite) TestStore_Update() {
	tt := map[string]struct {
		want      interface{}
		mockForms func(f *fields.Repository, s *submissions.Repository)
		mock      func(m sqlmock.Sqlmock)
	}{
		"Success": {
			form,
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Insert", form.Id, form.Fields[0]).Return(nil)
				s.On("Find", form.Id).Return(form.Submissions, nil)
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(form.Id), 1))
			},
		},
		"No Rows": {
			"Error updating form with the name",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Error Fields": {
			"error",
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Insert", form.Id, formFields[0]).Return(&errors.Error{Message: "error"})
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(form.Id), 1))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock, test.mockForms)
			got, err := s.Update(form)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(got, test.want)
		})
	}
}
