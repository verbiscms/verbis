// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/sys"
	"net/http"
)

func (t *SystemTestSuite) TestInstall_Preflight() {
	tt := map[string]struct {
		want      interface{}
		status    int
		message   string
		input     interface{}
		installed bool
		mock      func(m *mocks.System)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully connected to database",
			&preflight,
			false,
			func(m *mocks.System) {
				m.On("Preflight", preflight).Return(nil)
			},
		},
		"Already Installed": {
			nil,
			http.StatusBadRequest,
			"Verbis is already installed",
			&preflight,
			true,
			func(m *mocks.System) {
				m.On("Preflight", preflight).Return(nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "db_host", Message: "Db Host is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			&preflightBadValidation,
			false,
			nil,
		},
		"Error": {
			nil,
			http.StatusBadRequest,
			"connection error",
			&preflight,
			false,
			func(m *mocks.System) {
				m.On("Preflight", preflight).Return(fmt.Errorf("connection error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/install", "/install", test.input, func(ctx *gin.Context) {
				s := t.Setup(test.mock)
				s.Installed = test.installed
				s.Preflight(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
