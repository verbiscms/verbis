// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	FindQuery       = "SELECT * FROM `posts` WHERE `id` = '" + postID + "' LIMIT 1"
	FindBySlugQuery = "SELECT * FROM `posts` WHERE `slug` = '" + post.Slug + "' LIMIT 1"
)

func (t *PostsTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			postDatum,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.Id, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).WillReturnRows(rows)
			},
		},
		"Not Found": {
			"No post exists with the ID",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"})
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).WillReturnRows(rows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindQuery))).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(post.Id, false)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}

func (t *PostsTestSuite) TestStore_FindBySlug() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			postDatum,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.Id, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindBySlugQuery))).WillReturnRows(rows)
			},
		},
		"Not Found": {
			"No post exists with the slug",
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"})
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindBySlugQuery))).WillReturnRows(rows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindBySlugQuery))).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.FindBySlug(post.Slug)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
