// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	mocks "github.com/verbiscms/verbis/api/mocks/store/posts"
	"net/http"
)

var mockCacheSuccess = func(c *cache.Store) {
	c.On("Invalidate", mock.Anything, mock.Anything).Return(nil)
}

func (t *PostsTestSuite) TestPosts_Update() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
		cache   func(c *cache.Store)
		url     string
	}{
		"Success": {
			postData,
			http.StatusOK,
			"Successfully updated post with ID: 123",
			post,
			func(m *mocks.Repository) {
				m.On("Update", postCreate).Return(postData, nil)
			},
			mockCacheSuccess,
			"/posts/123",
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "title", Message: "Title is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			postBadValidation,
			func(m *mocks.Repository) {
				m.On("Update", postBadValidation).Return(domain.PostDatum{}, fmt.Errorf("error"))
			},
			mockCacheSuccess,
			"/posts/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to update the post",
			post,
			func(m *mocks.Repository) {
				m.On("Update", postCreate).Return(domain.PostDatum{}, fmt.Errorf("error"))
			},
			mockCacheSuccess,
			"/posts/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			post,
			func(m *mocks.Repository) {
				m.On("Update", postCreate).Return(domain.PostDatum{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			mockCacheSuccess,
			"/posts/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			post,
			func(m *mocks.Repository) {
				m.On("Update", postCreate).Return(domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			mockCacheSuccess,
			"/posts/123",
		},
		"Cache Invalid Error": {
			nil,
			http.StatusInternalServerError,
			"cache",
			post,
			func(m *mocks.Repository) {
				m.On("Update", postCreate).Return(postData, nil)
			},
			func(c *cache.Store) {
				c.On("Invalidate", mock.Anything, mock.Anything).Return(&errors.Error{Code: errors.INTERNAL, Message: "cache"})
			},
			"/posts/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPut, test.url, "/posts/:id", test.input, func(ctx *gin.Context) {
				s := t.Setup(test.mock)
				c := &cache.Store{}
				test.cache(c)
				s.Deps.Cache = c
				s.Update(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
