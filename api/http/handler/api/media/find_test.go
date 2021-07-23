// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/services/media"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *MediaTestSuite) TestMedia_Find() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Library)
		url     string
	}{
		"Success": {
			mediaItem.Public(),
			http.StatusOK,
			"Successfully obtained media item with ID: 123",
			func(m *mocks.Library) {
				m.On("Find", 123).Return(mediaItem, nil)
			},
			"/media/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"Pass a valid number to obtain the media item by ID",
			func(m *mocks.Library) {
				m.On("Find", 123).Return(domain.Media{}, fmt.Errorf("error"))
			},
			"/media/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"no media items found",
			func(m *mocks.Library) {
				m.On("Find", 123).Return(domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: "no media items found"})
			},
			"/media/123",
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Library) {
				m.On("Find", 123).Return(domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/media/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/media/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Find(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
