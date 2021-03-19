// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	categories "github.com/ainsleyclark/verbis/api/mocks/store/posts/categories"
	fields "github.com/ainsleyclark/verbis/api/mocks/store/posts/fields"
	meta "github.com/ainsleyclark/verbis/api/mocks/store/posts/meta"
	"regexp"
)

var (
	DeleteQuery = "DELETE FROM `posts` WHERE `id` = '" + postID + "'"
)

func (t *PostsTestSuite) TestStore_Delete() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
		repo func(c *categories.Repository, f *fields.Repository, m *meta.Repository)
	}{
		"Success": {
			nil,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
				c.On("Delete", post.Id).Return(nil)
				f.On("Delete", post.Id).Return(nil)
				m.On("Delete", post.Id).Return(nil)
			},
		},
		"No Rows": {
			"No category exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(sql.ErrNoRows)
			},
			nil,
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
			nil,
		},
		"Category Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
				c.On("Delete", post.Id).Return(&errors.Error{Message: database.ErrQueryMessage})
				f.On("Delete", post.Id).Return(nil)
				m.On("Delete", post.Id).Return(nil)
			},
		},
		"Fields Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
				c.On("Delete", post.Id).Return(nil)
				f.On("Delete", post.Id).Return(&errors.Error{Message: database.ErrQueryMessage})
				m.On("Delete", post.Id).Return(nil)
			},
		},
		"Meta Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			func(c *categories.Repository, f *fields.Repository, m *meta.Repository) {
				c.On("Delete", post.Id).Return(nil)
				f.On("Delete", post.Id).Return(nil)
				m.On("Delete", post.Id).Return(&errors.Error{Message: database.ErrQueryMessage})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupMock(test.mock, test.repo)
			err := s.Delete(post.Id)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				if t.Logger.String() != "" {
					t.Contains(t.Logger.String(), test.want)
				}
				return
			}
			t.RunT(nil, err)
		})
	}
}
