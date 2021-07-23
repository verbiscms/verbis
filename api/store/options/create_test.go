// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/test"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO `options` (`option_name`, `option_value`) VALUES ('" + optionName + "', ?)"
)

func (t *OptionsTestSuite) TestStore_Create() {
	tt := map[string]struct {
		input interface{}
		want  interface{}
		mock  func(m sqlmock.Sqlmock)
	}{
		"Success": {
			optionValue,
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyJSONMessage{}).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"No Rows": {
			optionValue,
			"Error creating option with the name",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyJSONMessage{}).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			optionValue,
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyJSONMessage{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Failed Marshal": {
			make(chan int, 1),
			"Error marshalling the option",
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Create(optionName, test.input)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, err)
		})
	}
}
