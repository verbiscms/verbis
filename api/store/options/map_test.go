// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"regexp"
)

var (
	MapQuery = "SELECT * FROM `options`"
)

func (t *OptionsTestSuite) TestStore_Map() {
	raw := json.RawMessage("\"test\"")

	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			domain.OptionsDBMap{
				"name": raw,
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"option_name", "option_value"}).
					AddRow("name", raw)
				m.ExpectQuery(regexp.QuoteMeta(MapQuery)).
					WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No options available",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(MapQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(MapQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Map()
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
