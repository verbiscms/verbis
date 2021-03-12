// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

//var (
//	FindQuery       = "SELECT * FROM `categories` WHERE `id` = '" + categoryID + "' LIMIT 1"
//	FindByPostQuery = "SELECT `c`.* FROM `post_categories` LEFT JOIN `categories` AS `c` ON `post_categories`.`post_id` = `c`.`id` WHERE `post_categories`.`post_id` = '1'"
//	FindBySlugQuery = "SELECT * FROM `categories` WHERE `slug` = '" + category.Slug + "' LIMIT 1"
//	FindByNameQuery = "SELECT * FROM `categories` WHERE `name` = '" + category.Name + "' LIMIT 1"
//	FindParentQuery = "SELECT * FROM `categories` WHERE `parent_id` = '" + categoryID + "' LIMIT 1"
//)
//
//func (t *CategoriesTestSuite) TestStore_Find() {
//	tt := map[string]struct {
//		want interface{}
//		mock func(m sqlmock.Sqlmock)
//	}{
//		"Success": {
//			category,
//			func(m sqlmock.Sqlmock) {
//				rows := sqlmock.NewRows([]string{"id", "slug", "name"}).
//					AddRow(category.Id, category.Slug, category.Name)
//				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnRows(rows)
//			},
//		},
//		"No Rows": {
//			"No category exists with the ID",
//			func(m sqlmock.Sqlmock) {
//				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnError(sql.ErrNoRows)
//			},
//		},
//		"Internal Error": {
//			database.ErrQueryMessage,
//			func(m sqlmock.Sqlmock) {
//				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).WillReturnError(fmt.Errorf("error"))
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(test.mock)
//			got, err := s.Find(category.Id)
//			if err != nil {
//				t.Contains(errors.Message(err), test.want)
//				return
//			}
//			t.RunT(test.want, got)
//		})
//	}
//}
//
//func (t *CategoriesTestSuite) TestStore_FindByPost() {
//	tt := map[string]struct {
//		want interface{}
//		mock func(m sqlmock.Sqlmock)
//	}{
//		"Success": {
//			category,
//			func(m sqlmock.Sqlmock) {
//				rows := sqlmock.NewRows([]string{"id", "slug", "name"}).
//					AddRow(category.Id, category.Slug, category.Name)
//				m.ExpectQuery(regexp.QuoteMeta(FindByPostQuery)).WithArgs().WillReturnRows(rows)
//			},
//		},
//		"No Rows": {
//			"No category exists with the post ID",
//			func(m sqlmock.Sqlmock) {
//				m.ExpectQuery(regexp.QuoteMeta(FindByPostQuery)).WillReturnError(sql.ErrNoRows)
//			},
//		},
//		"Internal Error": {
//			database.ErrQueryMessage,
//			func(m sqlmock.Sqlmock) {
//				m.ExpectQuery(regexp.QuoteMeta(FindByPostQuery)).WillReturnError(fmt.Errorf("error"))
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(test.mock)
//			got, err := s.FindByPost(1)
//			if err != nil {
//				t.Contains(errors.Message(err), test.want)
//				return
//			}
//			t.RunT(test.want, got)
//		})
//	}
//}
//
//func (t *CategoriesTestSuite) TestStore_FindBySlug() {
//	tt := map[string]struct {
//		want interface{}
//		mock func(m sqlmock.Sqlmock)
//	}{
//		"Success": {
//			category,
//			func(m sqlmock.Sqlmock) {
//				rows := sqlmock.NewRows([]string{"id", "slug", "name"}).
//					AddRow(category.Id, category.Slug, category.Name)
//				m.ExpectQuery(regexp.QuoteMeta(FindBySlugQuery)).WillReturnRows(rows)
//			},
//		},
//		"No Rows": {
//			"No category exists with the slug",
//			func(m sqlmock.Sqlmock) {
//				m.ExpectQuery(regexp.QuoteMeta(FindBySlugQuery)).WillReturnError(sql.ErrNoRows)
//			},
//		},
//		"Internal Error": {
//			database.ErrQueryMessage,
//			func(m sqlmock.Sqlmock) {
//				m.ExpectQuery(regexp.QuoteMeta(FindBySlugQuery)).WillReturnError(fmt.Errorf("error"))
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(test.mock)
//			got, err := s.FindBySlug(category.Slug)
//			if err != nil {
//				t.Contains(errors.Message(err), test.want)
//				return
//			}
//			t.RunT(test.want, got)
//		})
//	}
//}
//
//func (t *CategoriesTestSuite) TestStore_FindByName() {
//	tt := map[string]struct {
//		want interface{}
//		mock func(m sqlmock.Sqlmock)
//	}{
//		"Success": {
//			category,
//			func(m sqlmock.Sqlmock) {
//				rows := sqlmock.NewRows([]string{"id", "slug", "name"}).
//					AddRow(category.Id, category.Slug, category.Name)
//				m.ExpectQuery(regexp.QuoteMeta(FindByNameQuery)).WillReturnRows(rows)
//			},
//		},
//		"No Rows": {
//			"No category exists with the name",
//			func(m sqlmock.Sqlmock) {
//				m.ExpectQuery(regexp.QuoteMeta(FindByNameQuery)).WillReturnError(sql.ErrNoRows)
//			},
//		},
//		"Internal Error": {
//			database.ErrQueryMessage,
//			func(m sqlmock.Sqlmock) {
//				m.ExpectQuery(regexp.QuoteMeta(FindByNameQuery)).WillReturnError(fmt.Errorf("error"))
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(test.mock)
//			got, err := s.FindByName(category.Name)
//			if err != nil {
//				t.Contains(errors.Message(err), test.want)
//				return
//			}
//			t.RunT(test.want, got)
//		})
//	}
//}
//
//func (t *CategoriesTestSuite) TestStore_FindParent() {
//	tt := map[string]struct {
//		want interface{}
//		mock func(m sqlmock.Sqlmock)
//	}{
//		"Success": {
//			category,
//			func(m sqlmock.Sqlmock) {
//				rows := sqlmock.NewRows([]string{"id", "slug", "name"}).
//					AddRow(category.Id, category.Slug, category.Name)
//				m.ExpectQuery(regexp.QuoteMeta(FindParentQuery)).WillReturnRows(rows)
//			},
//		},
//		"No Rows": {
//			"No category exists with the parent ID",
//			func(m sqlmock.Sqlmock) {
//				m.ExpectQuery(regexp.QuoteMeta(FindParentQuery)).WillReturnError(sql.ErrNoRows)
//			},
//		},
//		"Internal Error": {
//			database.ErrQueryMessage,
//			func(m sqlmock.Sqlmock) {
//				m.ExpectQuery(regexp.QuoteMeta(FindParentQuery)).WillReturnError(fmt.Errorf("error"))
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(test.mock)
//			got, err := s.FindParent(category.Id)
//			if err != nil {
//				t.Contains(errors.Message(err), test.want)
//				return
//			}
//			t.RunT(test.want, got)
//		})
//	}
//}
