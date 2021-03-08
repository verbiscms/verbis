// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

const (
	DeleteQuery      = "DELETE FROM `categories` WHERE `id` = '" + categoryID + "'"
	DeletePivotQuery = "DELETE FROM `post_categories` WHERE `id` = '" + categoryID + "'"
)

func (t *CategoriesTestSuite) TestStore_Delete() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			category,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).WithArgs(category.Id).WillReturnResult(sqlmock.NewResult(0, 1))
				m.ExpectExec(regexp.QuoteMeta(DeletePivotQuery)).WithArgs(category.Id).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"No Rows": {
			"No category exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).WithArgs(category.Id).WillReturnResult(sqlmock.NewResult(0, 0))
				m.ExpectExec(regexp.QuoteMeta(DeletePivotQuery)).WithArgs(category.Id).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"No Rows Pivot": {
			"No category exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).WithArgs(category.Id).WillReturnResult(sqlmock.NewResult(0, 1))
				m.ExpectExec(regexp.QuoteMeta(DeletePivotQuery)).WithArgs(category.Id).WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		"Internal Error": {
			"Error executing sql query",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
		"Internal Error Pivot": {
			"Error executing sql query",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).WithArgs(category.Id).WillReturnResult(sqlmock.NewResult(0, 0))
				m.ExpectExec(regexp.QuoteMeta(DeletePivotQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Delete(category.Id)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Nil(err)
		})
	}
}
