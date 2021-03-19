// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postcategories

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	UpdateQuery = "UPDATE `post_categories` SET `category_id` = " + categoryID + " WHERE `post_id` = '" + postID + ""
)

func (t *PostCategoriesTestSuite) TestStore_Update() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"No Rows": {
			"Error updating post category with the post ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Update(post.Id, post.Category.Id)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, err)
		})
	}
}
