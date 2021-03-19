// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	DeleteQuery              = "DELETE FROM `categories` WHERE `id` = '" + categoryID + "'"
	DeletePivotCategoryQuery = "DELETE FROM `post_categories` WHERE `category_id` = '" + categoryID + "'"
)

func (t *CategoriesTestSuite) TestStore_Delete() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				m.ExpectExec(regexp.QuoteMeta(DeletePivotCategoryQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"No Rows": {
			"No category exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"No Rows Pivot": {
			"No category exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				m.ExpectExec(regexp.QuoteMeta(DeletePivotCategoryQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Internal Error Pivot": {
			"Error executing sql query",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 0))
				m.ExpectExec(regexp.QuoteMeta(DeletePivotCategoryQuery)).
					WillReturnError(fmt.Errorf("error"))
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
			t.RunT(nil, err)
		})
	}
}


// deletePivotByCategoryID
//
// Returns nil if the category was successfully deleted.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if the category was not found.
func (s *Store) deletePivotCategory(id int) error {
	const op = "CategoryStore.deletePivotCategory"

	q := s.Builder().
		DeleteFrom(s.Schema()+PivotTableName).
		Where("category_id", "=", id)

	_, err := s.DB().Exec(q.Build(), id)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}

