// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *PostsTestSuite) TestPosts_Create() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.PostsRepository)
	}{
		"Success": {
			postData,
			200,
			"Successfully created post with ID: 123",
			post,
			func(m *mocks.PostsRepository) {
				m.On("Create", &postCreate).Return(postData, nil)
			},
		},
		"Validation Failed": {
			api.ValidationErrJson{Errors: validation.Errors{{Key: "slug", Message: "Post Slug is required.", Type: "required"}}},
			400,
			"Validation failed",
			postBadValidation,
			func(m *mocks.PostsRepository) {
				m.On("Create", &postBadValidation).Return(domain.PostData{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			nil,
			400,
			"invalid",
			post,
			func(m *mocks.PostsRepository) {
				m.On("Create", &postCreate).Return(domain.PostData{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			400,
			"conflict",
			post,
			func(m *mocks.PostsRepository) {
				m.On("Create", &postCreate).Return(domain.PostData{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			post,
			func(m *mocks.PostsRepository) {
				m.On("Create", &postCreate).Return(domain.PostData{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/posts", "/posts", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Create(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}