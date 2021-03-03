// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

//func (t *SiteTestSuite) TestSite_Templates() {
//
//	tt := map[string]struct {
//		want    interface{}
//		status  int
//		message string
//		mock    func(m *mocks.Repository)
//	}{
//		"Success": {
//			templates,
//			200,
//			"Successfully obtained templates",
//			func(m *mocks.Repository) {
//				m.On("Templates", t.ThemePath).Return(templates, nil)
//			},
//		},
//		"Not Found": {
//			nil,
//			200,
//			"not found",
//			func(m *mocks.Repository) {
//				m.On("Templates", t.ThemePath).Return(domain.Templates{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
//			},
//		},
//		"Internal Error": {
//			nil,
//			500,
//			"internal",
//			func(m *mocks.Repository) {
//				m.On("Templates", t.ThemePath).Return(domain.Templates{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			t.RequestAndServe(http.MethodGet, "/theme", "/theme", nil, func(ctx *gin.Context) {
//				t.Setup(test.mock).Templates(ctx)
//			})
//			t.RunT(test.want, test.status, test.message)
//		})
//	}
//}
