// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
)

var (
	ExistsQuery = "SELECT EXISTS (SELECT `id` FROM `post_options` WHERE `post_id` =  '" + postID + "')"
)

func (t *MetaTestSuite) TestStore_Exists() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Exists": {
			true,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)
			},
		},
		"Not Found": {
			false,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnRows(rows)
			},
		},
		"Internal": {
			false,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got := s.Exists(meta.PostId)
			t.RunT(test.want, got)
		})
	}
}
