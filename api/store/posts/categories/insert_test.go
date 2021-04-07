// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

//
//func (t *PostCategoriesTestSuite) TestStore_Insert() {
//	tt := map[string]struct {
//		want interface{}
//		mock func(m sqlmock.Sqlmock)
//	}{
//		"Update": {
//			nil,
//			func(m sqlmock.Sqlmock) {
//				rows := sqlmock.NewRows([]string{"id"}).
//					AddRow(true)
//				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
//					WillReturnRows(rows)
//				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
//					WithArgs(meta.Seo, meta.Meta, meta.PostId).
//					WillReturnResult(sqlmock.NewResult(int64(meta.Id), 1))
//			},
//		},
//		"Update Error": {
//			"error",
//			func(m sqlmock.Sqlmock) {
//				rows := sqlmock.NewRows([]string{"id"}).
//					AddRow(true)
//				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
//					WillReturnRows(rows)
//				m.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
//					WithArgs(meta.Seo, meta.Meta, meta.PostId).
//					WillReturnError(fmt.Errorf("error"))
//			},
//		},
//		"Create": {
//			nil,
//			func(m sqlmock.Sqlmock) {
//				rows := sqlmock.NewRows([]string{"id"}).
//					AddRow(false)
//				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
//					WillReturnRows(rows)
//				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
//					WithArgs(meta.PostId, meta.Seo, meta.Meta).
//					WillReturnResult(sqlmock.NewResult(int64(meta.Id), 1))
//			},
//		},
//		"Create Error": {
//			"error",
//			func(m sqlmock.Sqlmock) {
//				rows := sqlmock.NewRows([]string{"id"}).
//					AddRow(false)
//				m.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
//					WillReturnRows(rows)
//				m.ExpectExec(regexp.QuoteMeta(CreateQuery)).
//					WithArgs(meta.PostId, meta.Seo, meta.Meta).
//					WillReturnError(fmt.Errorf("error"))
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(test.mock)
//			err := s.Insert(meta.PostId, meta)
//			if err != nil {
//				t.Contains(err.Error(), test.want)
//				return
//			}
//			t.RunT(nil, err)
//		})
//	}
//}
