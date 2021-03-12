// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

//
//var (
//	CreateQuery = "INSERT INTO `categories` (`uuid`, `slug`, `name`, `description`, `parent_id`, `resource`, `archive_id`, `updated_at`, `created_at`) VALUES (?, '/cat', 'Category', NULL, NULL, '', NULL, NOW(), NOW())"
//)
//
//func (t *OptionsTestSuite) TestStore_Create() {
//	tt := map[string]struct {
//		want interface{}
//		mock func(m sqlmock.Sqlmock)
//	}{
//		"Success": {
//			category,
//			func(m sqlmock.Sqlmock) {
//				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
//					WithArgs(test.DBAnyString{}).
//					WillReturnResult(sqlmock.NewResult(int64(category.Id), 1))
//			},
//		},
//		"Validation Failed": {
//			"Validation failed, the category name already exists",
//			func(m sqlmock.Sqlmock) {
//				rows := sqlmock.NewRows([]string{"id"}).
//					AddRow(true)
//				m.ExpectQuery(regexp.QuoteMeta(ExistsByFromQuery)).WillReturnRows(rows)
//			},
//		},
//		"No Rows": {
//			"Error creating category with the name",
//			func(m sqlmock.Sqlmock) {
//				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
//					WithArgs(test.DBAnyString{}).
//					WillReturnError(sql.ErrNoRows)
//			},
//		},
//		"Internal Error": {
//			database.ErrQueryMessage,
//			func(m sqlmock.Sqlmock) {
//				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
//					WithArgs(test.DBAnyString{}).
//					WillReturnError(fmt.Errorf("error"))
//			},
//		},
//		"Last Insert ID Error": {
//			"Error getting the newly created category ID",
//			func(m sqlmock.Sqlmock) {
//				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
//					WithArgs(test.DBAnyString{}).
//					WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("err")))
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(test.mock)
//			cat, err := s.Create(category)
//			if err != nil {
//				t.Contains(errors.Message(err), test.want)
//				return
//			}
//			t.RunT(cat, test.want, 2)
//		})
//	}
//}
