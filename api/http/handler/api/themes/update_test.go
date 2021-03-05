// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	mockStore "github.com/ainsleyclark/verbis/api/mocks/models"
	mocks "github.com/ainsleyclark/verbis/api/mocks/verbis/theme"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *ThemesTestSuite) TestThemes_Update() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository, mo *mockStore.OptionsRepository)
		url     string
	}{
		"Success": {
			config.DefaultTheme,
			http.StatusOK,
			"Successfully changed theme with the name: " + TestActiveTheme,
			UpdateTheme{
				Theme: TestActiveTheme,
			},
			func(m *mocks.Repository, mo *mockStore.OptionsRepository) {
				m.On("Exists", TestActiveTheme).Return(true)
				m.On("Templates", TestActiveTheme).Return(templates, nil)
				mo.On("SetTheme", TestActiveTheme).Return(nil)
			},
			"/themes",
		},
		"Validation Failed": {
			nil,
			http.StatusBadRequest,
			"Validation failed",
			nil,
			nil,
			"/themes",
		},
		"Exists": {
			nil,
			http.StatusBadRequest,
			"No theme exists with the name: " + TestActiveTheme,
			UpdateTheme{
				Theme: TestActiveTheme,
			},
			func(m *mocks.Repository, mo *mockStore.OptionsRepository) {
				m.On("Exists", TestActiveTheme).Return(false)
			},
			"/themes",
		},
		"Error Database Set": {
			nil,
			http.StatusInternalServerError,
			"Error setting theme",
			UpdateTheme{
				Theme: TestActiveTheme,
			},
			func(m *mocks.Repository, mo *mockStore.OptionsRepository) {
				m.On("Exists", TestActiveTheme).Return(true)
				m.On("Templates", TestActiveTheme).Return(templates, nil)
				mo.On("SetTheme", TestActiveTheme).Return(fmt.Errorf("error"))
			},
			"/themes",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/theme", "/theme", test.input, func(ctx *gin.Context) {
				t.SetupOptions(test.mock).Update(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
