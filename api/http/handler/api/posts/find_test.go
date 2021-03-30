// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *PostsTestSuite) TestPosts_Find() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			postData,
			http.StatusOK,
			"Successfully obtained post with ID: 123",
			func(m *mocks.Repository) {
				m.On("Find", 123, true).Return(postData, nil)
			},
			"/posts/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"Pass a valid number to obtain the post by ID",
			func(m *mocks.Repository) {
			},
			"/posts/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"no posts found",
			func(m *mocks.Repository) {
				m.On("Find", 123, true).Return(domain.PostDatum{}, &errors.Error{Code: errors.NOTFOUND, Message: "no posts found"})
				m.On("Format", post).Return(post, nil)
			},
			"/posts/123",
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"config",
			func(m *mocks.Repository) {
				m.On("Find", 123, true).Return(domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: "config"})
			},
			"/posts/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/posts/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Find(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
