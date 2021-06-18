// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/sys"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *SystemTestSuite) TestRoles_List() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.System)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Verbis updated successfully to version v0.0.1, restarting system....",
			func(m *mocks.System) {
				m.On("Update").Return("v0.0.1", nil)
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			func(m *mocks.System) {
				m.On("Update").Return("", &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.System) {
				m.On("Update").Return("", &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe("GET", "/roles", "/roles", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Update(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
