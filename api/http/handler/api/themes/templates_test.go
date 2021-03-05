// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/verbis/theme"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *ThemesTestSuite) TestThemes_Templates() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
	}{
		"Success": {
			templates,
			http.StatusOK,
			"Successfully obtained templates",
			func(m *mocks.Repository) {
				m.On("Templates", TestActiveTheme).Return(templates, nil)
			},
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"not found",
			func(m *mocks.Repository) {
				m.On("Templates", TestActiveTheme).Return(domain.Templates{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Repository) {
				m.On("Templates", TestActiveTheme).Return(domain.Templates{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/theme", "/theme", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Templates(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
