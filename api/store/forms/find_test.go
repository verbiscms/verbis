// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	fields "github.com/verbiscms/verbis/api/mocks/store/forms/fields"
	submissions "github.com/verbiscms/verbis/api/mocks/store/forms/submissions"
	"regexp"
)

var (
	FindQuery       = "SELECT * FROM `forms` WHERE `id` = '" + formID + "' LIMIT 1"
	FindByUUIDQuery = "SELECT * FROM `forms` WHERE `uuid` = '" + form.UUID.String() + "' LIMIT 1"
)

func (t *FormsTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want      interface{}
		mockForms func(f *fields.Repository, s *submissions.Repository)
		mock      func(m sqlmock.Sqlmock)
	}{
		"Success": {
			form,
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Find", form.Id).Return(form.Fields, nil)
				s.On("Find", form.Id).Return(form.Submissions, nil)
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(form.Id, form.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No form exists with the ID",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock, test.mockForms)
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
		want      interface{}
		mockForms func(f *fields.Repository, s *submissions.Repository)
		mock      func(m sqlmock.Sqlmock)
	}{
		"Success": {
			form,
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Find", form.Id).Return(form.Fields, nil)
				s.On("Find", form.Id).Return(form.Submissions, nil)
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(form.Id, form.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindByUUIDQuery)).
					WillReturnRows(rows)
			},
		},
		"Find Error": {
			"error",
			func(f *fields.Repository, s *submissions.Repository) {
				f.On("Find", form.Id).Return(domain.FormFields{}, &errors.Error{Message: "error"})
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(form.Id, form.Name)
				m.ExpectQuery(regexp.QuoteMeta(FindByUUIDQuery)).
					WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No form exists with the UUID",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByUUIDQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByUUIDQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock, test.mockForms)
			got, err := s.FindByUUID(form.UUID)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
