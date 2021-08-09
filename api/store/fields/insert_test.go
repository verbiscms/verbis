// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/test"
	"regexp"
)

func (t *FieldsTestSuite) TestStore_Insert() {
	tt := map[string]struct {
		want   interface{}
		fields domain.PostFields
		mock   func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			fieldsSingular,
			func(m sqlmock.Sqlmock) {
				// Delete
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))

				// Create
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"Delete Error": {
			database.ErrQueryMessage,
			fieldsSingular,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Insert Error": {
			database.ErrQueryMessage,
			fieldsSingular,
			func(m sqlmock.Sqlmock) {
				// Delete
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))

				// Create
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Insert(field.PostID, test.fields)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
