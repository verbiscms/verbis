// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/options"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *OptionsTestSuite) TestOptions_List() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
	}{
		"Success": {
			options,
			http.StatusOK,
			"Successfully obtained options",
			func(m *mocks.Repository) {
				m.On("Map").Return(options, nil)
			},
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"no options found",
			func(m *mocks.Repository) {
				m.On("Map").Return(nil, &errors.Error{Code: errors.NOTFOUND, Message: "no options found"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			func(m *mocks.Repository) {
				m.On("Map").Return(nil, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			func(m *mocks.Repository) {
				m.On("Map").Return(nil, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"config",
			func(m *mocks.Repository) {
				m.On("Map").Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "config"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/options", "/options", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).List(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
