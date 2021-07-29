// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/services/theme"
	"net/http"
)

func (t *ThemesTestSuite) TestThemes_Templates() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Service)
	}{
		"Success": {
			templates,
			http.StatusOK,
			"Successfully obtained templates",
			func(m *mocks.Service) {
				m.On("Templates").Return(templates, nil)
			},
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"not found",
			func(m *mocks.Service) {
				m.On("Templates").Return(domain.Templates{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Service) {
				m.On("Templates").Return(domain.Templates{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
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
