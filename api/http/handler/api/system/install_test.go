// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"github.com/gin-gonic/gin"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/sys"
	"net/http"
)

func (t *SystemTestSuite) TestInstall_Install() {
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
			"Successfully installed Verbis",
			&install,
			false,
			func(m *mocks.System) {
				m.On("Install", install).Return(nil)
			},
		},
		"Already Installed": {
			nil,
			http.StatusBadRequest,
			"Verbis is already installed",
			&install,
			true,
			func(m *mocks.System) {
				m.On("Install", install).Return(nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "site_title", Message: "Site Title is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			&installBadValidation,
			false,
			nil,
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			&install,
			false,
			func(m *mocks.System) {
				m.On("Install", install).Return(&errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			&install,
			false,
			func(m *mocks.System) {
				m.On("Install", install).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/install", "/install", test.input, func(ctx *gin.Context) {
				s := t.Setup(test.mock)
				s.Installed = test.installed
				s.Install(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
