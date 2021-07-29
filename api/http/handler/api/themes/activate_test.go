// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/services/theme"
	"net/http"
)

func (t *ThemesTestSuite) TestThemes_Activate() {
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
			"Successfully changed theme with the name: " + TestActiveTheme,
			func(m *mocks.Service) {
				m.On("Activate", TestActiveTheme).Return(config.DefaultTheme, nil)
			},
			"/themes/" + TestActiveTheme,
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			func(m *mocks.Service) {
				m.On("Activate", TestActiveTheme).Return(config.DefaultTheme, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			"/themes/" + TestActiveTheme,
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			func(m *mocks.Service) {
				m.On("Activate", TestActiveTheme).Return(config.DefaultTheme, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
			"/themes/" + TestActiveTheme,
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Service) {
				m.On("Activate", TestActiveTheme).Return(config.DefaultTheme, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/themes/" + TestActiveTheme,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, test.url, "/themes/:name", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Activate(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
