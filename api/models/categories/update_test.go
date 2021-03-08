// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	UpdateQuery = "INSERT INTO `categories` (uuid, slug, name, description, parent_id, resource, archive_id, updated_at, created_at) VALUES (?, '/cat', 'Category', NULL, NULL, '', NULL, NOW(), NOW())"
)

func (t *CategoriesTestSuite) TestStore_Update() {

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
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			cat, err := s.Update(category)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(cat, test.want)
		})
	}
}
