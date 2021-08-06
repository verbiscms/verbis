// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mocks "github.com/verbiscms/verbis/api/mocks/sys"
	"github.com/verbiscms/verbis/api/sys"
	"net/http"
)

func (t *SystemTestSuite) TestInstall_Preflight() {
	tt := map[string]struct {
		want      interface{}
		status    int
		message   string
		url string
		input     interface{}
		installed bool
		mock      func(m *mocks.System)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully validated install data",
			"/install/1",
			&install,
			false,
			func(m *mocks.System) {
				m.On("ValidateInstall", sys.InstallDatabaseStep, install).Return(nil)
			},
		},
		"Already Installed": {
			nil,
			http.StatusBadRequest,
			"Verbis is already installed",
			"/install/1",
			nil,
			true,
			nil,
		},
		"No Input": {
			nil,
			http.StatusBadRequest,
			"Pass a valid number to validate a step",
			"/install/wrong",
			nil,
			false,
			nil,
		},
		"Decode Error": {
			nil,
			http.StatusBadRequest,
			"Error unmarshalling install",
			"/install/1",
			"wrong",
			false,
			nil,
		},
		"Validation Failed": {
			nil,
			http.StatusBadRequest,
			"Validation failed",
			"/install/1",
			&install,
			false,
			func(m *mocks.System) {
				m.On("ValidateInstall", sys.InstallDatabaseStep, install).Return(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, test.url, "/install/:step", test.input, func(ctx *gin.Context) {
				s := t.Setup(test.mock)
				s.Installed = test.installed
				s.ValidateInstall(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
