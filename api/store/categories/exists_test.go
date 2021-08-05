// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
)

var (
	ExistsQuery       = "SELECT EXISTS (SELECT `id` FROM `categories` WHERE `id` =  '" + categoryID + "')"
	ExistsByFromQuery = "SELECT EXISTS (SELECT `id` FROM `categories` WHERE `name` =  '" + category.Name + "')"
	ExistsBySlugQuery = "SELECT EXISTS (SELECT `id` FROM `categories` WHERE `slug` =  '" + category.Slug + "')"
)

func (t *CategoriesTestSuite) TestStore_Exists() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Exists": {
			true,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).WillReturnRows(rows)
			},
		},
		"Not Found": {
			false,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).WillReturnRows(rows)
			},
		},
		"Internal": {
			false,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got := s.Exists(category.ID)
			t.RunT(test.want, got)
		})
	}
}

func (t *CategoriesTestSuite) TestStore_ExistsByName() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Exists": {
			true,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsByFromQuery)).WillReturnRows(rows)
			},
		},
		"Not Found": {
			false,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsByFromQuery)).WillReturnRows(rows)
			},
		},
		"Internal": {
			false,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ExistsByFromQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got := s.ExistsByName(category.Name)
			t.RunT(test.want, got)
		})
	}
}

func (t *CategoriesTestSuite) TestStore_ExistsBySlug() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Exists": {
			true,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(true)
				m.ExpectQuery(regexp.QuoteMeta(ExistsBySlugQuery)).WillReturnRows(rows)
			},
		},
		"Not Found": {
			false,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(false)
				m.ExpectQuery(regexp.QuoteMeta(ExistsBySlugQuery)).WillReturnRows(rows)
			},
		},
		"Internal": {
			false,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(ExistsBySlugQuery)).WillReturnError(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got := s.ExistsBySlug(category.Slug)
			t.RunT(test.want, got)
		})
	}
}
