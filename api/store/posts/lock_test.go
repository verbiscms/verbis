// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"regexp"
)

var (
	LockQuery   = "UPDATE `post_options` SET `edit_lock` = 'token' WHERE `post_id` = '" + postID + "'"
	UnlockQuery = "UPDATE `post_options` SET `edit_lock` = '' WHERE `post_id` = '" + postID + "'"
)

func (t *PostsTestSuite) TestStore_Lock() {
	tt := map[string]struct {
		mock func(m sqlmock.Sqlmock)
		want interface{}
	}{
		"Success": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(LockQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
			nil,
		},
		"No Rows": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(LockQuery)).
					WillReturnError(sql.ErrNoRows)
			},
			"No post exists with the ID",
		},
		"Internal Error": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(LockQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
			database.ErrQueryMessage,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Lock(1, "token")
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}

func (t *PostsTestSuite) TestStore_Unlock() {
	tt := map[string]struct {
		mock func(m sqlmock.Sqlmock)
		want interface{}
	}{
		"Success": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UnlockQuery)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
			nil,
		},
		"No Rows": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UnlockQuery)).
					WillReturnError(sql.ErrNoRows)
			},
			"No post exists with the ID",
		},
		"Internal Error": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(UnlockQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
			database.ErrQueryMessage,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Unlock(1)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
