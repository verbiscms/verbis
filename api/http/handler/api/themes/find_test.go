// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/verbis/theme"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *ThemesTestSuite) TestThemes_Find() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			config.DefaultTheme,
			http.StatusOK,
			"Successfully obtained theme config",
			func(m *mocks.Repository) {
				m.On("Find", TestActiveTheme).Return(&config.DefaultTheme, nil)
			},
			"/themes/verbis",
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"not found",
			func(m *mocks.Repository) {
				m.On("Find", "wrongname").Return(nil, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/themes/wrongname",
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Repository) {
				m.On("Find", TestActiveTheme).Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
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
