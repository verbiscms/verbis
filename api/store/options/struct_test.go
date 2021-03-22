// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"math"
	"regexp"
)

func (t *OptionsTestSuite) TestStore_Struct() {
	raw := json.RawMessage("\"verbis\"")

	tt := map[string]struct {
		panics bool
		want   interface{}
		mock   func(m sqlmock.Sqlmock)
	}{
		"Success": {
			false,
			domain.Options{
				ActiveTheme: "verbis",
			},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"option_name", "option_value"}).
					AddRow("active_theme", raw)
				m.ExpectQuery(regexp.QuoteMeta(MapQuery)).
					WillReturnRows(rows)
			},
		},
		"Internal": {
			true,
			domain.Options{},
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(MapQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Marshal Error": {
			true,
			domain.Options{},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"option_name", "option_value"}).
					AddRow("active_theme", math.Inf(1))
				m.ExpectQuery(regexp.QuoteMeta(MapQuery)).
					WillReturnRows(rows)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)

			if test.panics {
				t.Panics(func() {
					got := s.Struct()
					t.Equal(test.want, got)
				})
				return
			}

			got := s.Struct()
			t.RunT(test.want, got)
		})
	}
}
