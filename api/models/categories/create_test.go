// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/test"
	"regexp"
)

const (
	CreateQuery = "INSERT INTO `categories` (\"id\", \"uuid\", \"slug\", \"name\", \"primary\", \"description\", \"resource\", \"parent_id\", \"archive_id\", \"updated_at\", \"created_at\") VALUES (123, 00000000-0000-0000-0000-000000000000, '/cat', 'Category', TRUE, <nil>, '', <nil>, <nil>, NOW(), NOW()) "
)

func (t *CategoriesTestSuite) TestStore_Create() {

	tt := map[string]struct {
		mock func(m sqlmock.Sqlmock)
		want interface{}
	}{
		"Success": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).WillReturnResult(sqlmock.NewResult(int64(category.Id), 1))
			},
			category,
		},
		"Internal Error": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).WillReturnError(fmt.Errorf("error"))
			},
			"Error creating category with the name",
		},
		"Last Inserted ID Error": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).WillReturnResult(test.DBMockResultErr{})
			},
			"Error getting the newly created category ID",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			cat, err := s.Create(category)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(cat, test.want)
		})
	}
}
