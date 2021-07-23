// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/store/posts"
	"net/http"
)

func (t *PostsTestSuite) TestPosts_Delete() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully deleted post with ID: 123",
			func(m *mocks.Repository) {
				m.On("Delete", 123).Return(nil)
			},
			"/posts/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to delete a post",
			func(m *mocks.Repository) {
				m.On("Delete", 0).Return(nil)
			},
			"/posts/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			func(m *mocks.Repository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/posts/123",
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			func(m *mocks.Repository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			"/posts/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Repository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/posts/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodDelete, test.url, "/posts/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Delete(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
