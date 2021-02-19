// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *PostsTestSuite) TestPosts_Delete() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.PostsRepository)
		url     string
	}{
		"Success": {
			nil,
			200,
			"Successfully deleted post with ID: 123",
			func(m *mocks.PostsRepository) {
				m.On("Delete", 123).Return(nil)
			},
			"/posts/123",
		},
		"Invalid ID": {
			nil,
			400,
			"A valid ID is required to delete a post",
			func(m *mocks.PostsRepository) {
				m.On("Delete", 0).Return(nil)
			},
			"/posts/wrongid",
		},
		"Not Found": {
			nil,
			400,
			"not found",
			func(m *mocks.PostsRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/posts/123",
		},
		"Conflict": {
			nil,
			400,
			"conflict",
			func(m *mocks.PostsRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			"/posts/123",
		},
		"Internal": {
			nil,
			500,
			"internal",
			func(m *mocks.PostsRepository) {
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