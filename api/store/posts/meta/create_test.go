// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO post_options (post_id, seo, meta) VALUES (?, ?, ?)"
)

func (t *MetaTestSuite) TestStore_Create() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(meta.PostId, meta.Seo, meta.Meta).
					WillReturnResult(sqlmock.NewResult(int64(meta.Id), 1))
			},
		},
		"No Rows": {
			"Error creating meta with the post ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(meta.PostId, meta.Seo, meta.Meta).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(meta.PostId, meta.Seo, meta.Meta).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.create(meta.PostId, meta)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
