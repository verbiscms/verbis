// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	fields "github.com/verbiscms/verbis/api/mocks/store/forms/fields"
	submissions "github.com/verbiscms/verbis/api/mocks/store/forms/submissions"
	"github.com/verbiscms/verbis/api/test"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO `forms` (`uuid`, `name`, `email_send`, `email_message`, `email_subject`, `store_db`, `updated_at`, `created_at`) VALUES (?, 'Form', FALSE, '', '', FALSE, NOW(), NOW())"
)

func (t *FormsTestSuite) TestStore_Create() {
	tt := map[string]struct {
		want      interface{}
		mockForms func(f *fields.Repository, s *submissions.Repository)
		mock      func(m sqlmock.Sqlmock)
	}{
		"Success": {
			form,
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Insert", form.ID, form.Fields[0]).Return(nil)
				s.On("Find", form.ID).Return(form.Submissions, nil)
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(form.ID), 1))
			},
		},
		"No Rows": {
			"Error creating form with the name",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Last Insert ID Error": {
			"Error getting the newly created form ID",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("err")))
			},
		},
		"Error Fields": {
			"error",
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Insert", form.ID, formFields[0]).Return(&errors.Error{Message: "error"})
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(form.ID), 1))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock, test.mockForms)
			got, err := s.Create(form)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(got, test.want)
		})
	}
}
