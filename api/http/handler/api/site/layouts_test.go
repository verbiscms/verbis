// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/site"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *SiteTestSuite) TestSite_Layouts() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
	}{
		"Success": {
			layouts,
			200,
			"Successfully obtained layouts",
			func(m *mocks.Repository) {
				m.On("Layouts", t.ThemePath).Return(layouts, nil)
			},
		},
		"Not Found": {
			nil,
			200,
			"not found",
			func(m *mocks.Repository) {
				m.On("Layouts", t.ThemePath).Return(domain.Layouts{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			func(m *mocks.Repository) {
				m.On("Layouts", t.ThemePath).Return(domain.Layouts{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/theme", "/theme", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Layouts(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
