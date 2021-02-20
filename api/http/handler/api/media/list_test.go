// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *MediaTestSuite) TestMedia_List() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(u *mocks.MediaRepository)
	}{
		"Success": {
			mediaItems,
			200,
			"Successfully obtained media",
			func(m *mocks.MediaRepository) {
				m.On("Get", pagination).Return(mediaItems, 1, nil)
			},
		},
		"Not Found": {
			nil,
			200,
			"no media found",
			func(m *mocks.MediaRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no media found"})
			},
		},
		"Conflict": {
			nil,
			400,
			"conflict",
			func(m *mocks.MediaRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			nil,
			400,
			"invalid",
			func(m *mocks.MediaRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			func(m *mocks.MediaRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
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
