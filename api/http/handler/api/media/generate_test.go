// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/services/media"
	"net/http"
)

func (t *MediaTestSuite) TestMedia_ReGenerateWebP() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Library)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully started regeneration of WebP images: 5 items are being processed",
			func(m *mocks.Library) {
				m.On("ReGenerateWebP").Return(5, nil)
			},
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"not found",
			func(m *mocks.Library) {
				m.On("ReGenerateWebP").Return(0, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			func(m *mocks.Library) {
				m.On("ReGenerateWebP").Return(0, &errors.Error{Code: errors.CONFLICT, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			func(m *mocks.Library) {
				m.On("ReGenerateWebP").Return(0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Library) {
				m.On("ReGenerateWebP").Return(0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/media/generate/webp", "/media/generate/webp", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).ReGenerateWebP(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
