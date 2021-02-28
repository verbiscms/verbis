// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

const (
	FindQuery       = "SELECT * FROM `categories` WHERE `id` = ? LIMIT 1"
	FindBySlugQuery = "SELECT * FROM `categories` WHERE `slug` = ? LIMIT 1"
	FindByNameQuery = "SELECT * FROM `categories` WHERE `name` = ? LIMIT 1"
)

func (t *CategoriesTestSuite) TestStore_Find() {

	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			category,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "name", "primary"}).
					AddRow(category.Id, category.Slug, category.Name, category.Primary)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WithArgs(category.Id).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No category exists with the ID",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			"Error executing sql query",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(int64(category.Id))
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}

func (t *CategoriesTestSuite) TestStore_FindBySlug() {

	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			category,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "name", "primary"}).
					AddRow(category.Id, category.Slug, category.Name, category.Primary)
				m.ExpectQuery(regexp.QuoteMeta(FindBySlugQuery)).WithArgs(category.Slug).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No category exists with the slug",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindBySlugQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			"Error executing sql query",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindBySlugQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.FindBySlug(category.Slug)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}

func (t *CategoriesTestSuite) TestStore_FindByName() {

	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			category,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "name", "primary"}).
					AddRow(category.Id, category.Slug, category.Name, category.Primary)
				m.ExpectQuery(regexp.QuoteMeta(FindByNameQuery)).WithArgs(category.Slug).WillReturnRows(rows)
			},
		},
		"No Rows": {
			"No category exists with the name",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByNameQuery)).WillReturnError(sql.ErrNoRows)
			},
		},
		"Internal": {
			"Error executing sql query",
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindByNameQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.FindByName(category.Slug)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
