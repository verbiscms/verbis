// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

const (
	DeleteQuery = "SELECT * FROM `categories` WHERE `id` = ? LIMIT 1"
)

func (t *CategoriesTestSuite) TestStore_Delete() {

	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			category,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "name", "primary"}).
					AddRow(category.Id, category.Slug, category.Name, category.Primary)
				m.ExpectQuery(regexp.QuoteMeta(DeleteQuery)).WithArgs(category.Id).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No category exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(DeleteQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			"Error executing sql query",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(DeleteQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(int64(category.Id))
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
