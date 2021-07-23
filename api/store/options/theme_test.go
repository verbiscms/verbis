// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/errors"
	"regexp"
)

var (
	GetThemeQuery    = "SELECT `option_value` FROM `options` WHERE `option_name` = 'active_theme' LIMIT 1"
	UpdateThemeQuery = "UPDATE `options` SET `option_value` = ? WHERE `option_name` = 'active_theme'"
)

func (t *OptionsTestSuite) TestStore_GetTheme() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			"theme",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"option_value"}).
					AddRow("theme")
				m.ExpectQuery(regexp.QuoteMeta(GetThemeQuery)).
					WillReturnRows(rows)
			},
		},
		"Not Found": {
			"No theme exists in the option table: `active_theme`",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.GetTheme()
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}

func (t *OptionsTestSuite) TestStore_SetTheme() {
	tt := map[string]struct {
		input string
		want  interface{}
		mock  func(m sqlmock.Sqlmock)
	}{
		"Success": {
			"theme",
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateThemeQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"Error": {
			"theme",
			"Error updating option with the name: active_theme",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateThemeQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.SetTheme(test.input)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, err)
		})
	}
}
