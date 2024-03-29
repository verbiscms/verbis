// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/services/theme"
	"net/http"
)

func (t *ThemesTestSuite) TestThemes_Find() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Service)
		url     string
	}{
		"Success": {
			config.DefaultTheme,
			http.StatusOK,
			"Successfully obtained theme config",
			func(m *mocks.Service) {
				m.On("Find", TestActiveTheme).Return(config.DefaultTheme, nil)
			},
			"/themes/verbis",
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"not found",
			func(m *mocks.Service) {
				m.On("Find", "wrongname").Return(domain.ThemeConfig{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/themes/wrongname",
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Service) {
				m.On("Find", TestActiveTheme).Return(domain.ThemeConfig{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/themes/verbis",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/themes/:name", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Find(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
