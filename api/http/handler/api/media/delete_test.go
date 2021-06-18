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

func (t *MediaTestSuite) TestMedia_Delete() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(u *mocks.Repository)
		url     string
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully deleted media item with ID: 123",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(mediaItem, nil)
				m.On("Delete", 123).Return(nil)
			},
			"/media/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to delete a media item",
			func(m *mocks.Repository) {
				m.On("Delete", 123).Return(nil)
			},
			"/media/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(mediaItem, nil)
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/media/123",
		},
		"Find Error": {
			nil,
			http.StatusBadRequest,
			"No media item found with the ID: 123",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(mediaItem, fmt.Errorf("error"))
			},
			"/media/123",
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(mediaItem, nil)
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			"/media/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(mediaItem, nil)
				m.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/media/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodDelete, test.url, "/media/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Delete(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
