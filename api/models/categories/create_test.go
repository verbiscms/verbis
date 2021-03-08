// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/test"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO `categories` (uuid, slug, name, description, parent_id, resource, archive_id, updated_at, created_at) VALUES (?, '/cat', 'Category', NULL, NULL, '', NULL, NOW(), NOW())"
)

type AnyUUID struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyUUID) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

func (t *CategoriesTestSuite) TestStore_Create() {

	tt := map[string]struct {
		mock func(m sqlmock.Sqlmock)
		want interface{}
	}{
		"Success": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(AnyUUID{}).
					WillReturnResult(sqlmock.NewResult(int64(category.Id), 1))
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
