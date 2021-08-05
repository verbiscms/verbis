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
	"regexp"
)

var (
	DeleteQuery = "DELETE FROM `forms` WHERE `id` = '" + formID + "'"
)

func (t *FormsTestSuite) TestStore_Delete() {
	tt := map[string]struct {
		want      interface{}
		mockForms func(f *fields.Repository, s *submissions.Repository)
		mock      func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Delete", form.ID).Return(nil)
				s.On("Delete", form.ID).Return(nil)
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"No Rows": {
			"No form exists with the ID",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Fields Error": {
			"error",
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Delete", form.ID).Return(&errors.Error{Message: "error"})
				s.On("Delete", form.ID).Return(nil)
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"Submission Error": {
			"error",
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Delete", form.ID).Return(nil)
				s.On("Delete", form.ID).Return(&errors.Error{Message: "error"})
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock, test.mockForms)
			err := s.Delete(form.ID)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
