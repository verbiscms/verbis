// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

//
//var (
//	UpdateQuery = "UPDATE `categories` SET `uuid` = ?, `slug` = '/cat', `name` = 'Category', `description` = NULL, `parent_id` = NULL, `resource` = '', `archive_id` = NULL, `updated_at` = NOW() WHERE `id` = '1'"
//)
//
//func (t *CategoriesTestSuite) TestStore_Update() {
//	tt := map[string]struct {
//		want interface{}
//		mock func(m sqlmock.Sqlmock)
//	}{
//		"Success": {
//			category,
//			func(m sqlmock.Sqlmock) {
//				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
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
//			"Error updating category with the name",
//			func(m sqlmock.Sqlmock) {
//				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
//					WithArgs(test.DBAnyString{}).
//					WillReturnError(sql.ErrNoRows)
//			},
//		},
//		"Internal Error": {
//			database.ErrQueryMessage,
//			func(m sqlmock.Sqlmock) {
//				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
//					WithArgs(test.DBAnyString{}).
//					WillReturnError(fmt.Errorf("error"))
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(test.mock)
//			cat, err := s.Update(category)
//			if err != nil {
//				t.Contains(errors.Message(err), test.want)
//				return
//			}
//			t.RunT(cat, test.want, 2)
//		})
//	}
//}
