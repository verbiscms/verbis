// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *MediaTestSuite) TestMedia_Update() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.MediaRepository)
		url     string
	}{
		"Success": {
			want:    mediaItem,
			status:  200,
			message: "Successfully updated media item with ID: 123",
			input:   mediaItem,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", &mediaItem).Return(nil)
			},
			url: "/media/123",
		},
		"Validation Failed": {
			want:    nil,
			status:  400,
			message: "Validation failed",
			input:   `{"id": "wrongid"}`,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", mediaBadValidation).Return(fmt.Errorf("error"))
			},
			url: "/media/123",
		},
		"Invalid ID": {
			want:    nil,
			status:  400,
			message: "A valid ID is required to update the media item",
			input:   mediaItem,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", mediaItem).Return(fmt.Errorf("error"))
			},
			url: "/media/wrongid",
		},
		"Not Found": {
			want:    nil,
			status:  400,
			message: "not found",
			input:   &mediaItem,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", &mediaItem).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/media/123",
		},
		"Internal": {
			want:    nil,
			status:  500,
			message: "internal",
			input:   mediaItem,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", &mediaItem).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/media/123",
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
