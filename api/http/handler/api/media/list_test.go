// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/services/media"
	"github.com/ainsleyclark/verbis/api/test/dummy"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *MediaTestSuite) TestMedia_List() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Library)
	}{
		"Success": {
			mediaItems.Public(),
			http.StatusOK,
			"Successfully obtained media",
			func(m *mocks.Library) {
				m.On("List", dummy.DefaultParams).Return(mediaItems, 1, nil)
			},
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"no media found",
			func(m *mocks.Library) {
				m.On("List", dummy.DefaultParams).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no media found"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			func(m *mocks.Library) {
				m.On("List", dummy.DefaultParams).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			func(m *mocks.Library) {
				m.On("List", dummy.DefaultParams).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Library) {
				m.On("List", dummy.DefaultParams).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/media", "/media", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).List(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
