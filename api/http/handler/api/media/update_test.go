// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/media"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *MediaTestSuite) TestMedia_Update() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			mediaItem,
			http.StatusOK,
			"Successfully updated media item with ID: 123",
			mediaItem,
			func(m *mocks.Repository) {
				m.On("Update", &mediaItem).Return(nil)
			},
			"/media/123",
		},
		"Validation Failed": {
			nil,
			http.StatusBadRequest,
			"Validation failed",
			`{"id": "wrongid"}`,
			func(m *mocks.Repository) {
				m.On("Update", mediaBadValidation).Return(fmt.Errorf("error"))
			},
			"/media/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to update the media item",
			mediaItem,
			func(m *mocks.Repository) {
				m.On("Update", mediaItem).Return(fmt.Errorf("error"))
			},
			"/media/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			&mediaItem,
			func(m *mocks.Repository) {
				m.On("Update", &mediaItem).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/media/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"config",
			mediaItem,
			func(m *mocks.Repository) {
				m.On("Update", &mediaItem).Return(&errors.Error{Code: errors.INTERNAL, Message: "config"})
			},
			"/media/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPut, test.url, "/media/:id", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Update(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
