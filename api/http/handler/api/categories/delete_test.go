// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
)

func (t *CategoriesTestSuite) TestCategories_Delete() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.CategoryRepository)
		url     string
	}{
		"Success": {
			want:    nil,
			status:  200,
			message: "Successfully deleted category with ID: 123",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(nil)
			},
			url: "/categories/123",
		},
		"Invalid ID": {
			want:    nil,
			status:  400,
			message: "A valid ID is required to delete a category",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(nil)
			},
			url: "/categories/wrongid",
		},
		"Not Found": {
			want:    nil,
			status:  400,
			message: "not found",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/categories/123",
		},
		"Conflict": {
			want:    nil,
			status:  400,
			message: "conflict",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			url: "/categories/123",
		},
		"Internal": {
			want:    nil,
			status:  500,
			message: "internal",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/categories/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe("DELETE", test.url, "/categories/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Delete(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}