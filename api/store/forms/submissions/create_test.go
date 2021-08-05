// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package submissions

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/test"
	"regexp"
)

var (
	CreateQuery = "INSERT INTO `form_submissions` (`uuid`, `form_id`, `fields`, `ip_address`, `user_agent`, `sent_at`) VALUES (?, 1, ?, '127.0.0.1', 'chrome', NOW())"
)

func (t *SubmissionTestSuite) TestStore_Create() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}, formSubmission.Fields).
					WillReturnResult(sqlmock.NewResult(int64(formSubmission.ID), 1))
			},
		},
		"No Rows": {
			"Error creating form submission with the form ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}, formSubmission.Fields).
					WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(test.DBAnyString{}, formSubmission.Fields).
					WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Create(formSubmission)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(nil, err)
		})
	}
}
