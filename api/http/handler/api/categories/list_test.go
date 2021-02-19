// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
)

func (t *CategoriesTestSuite) TestCategories_List() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.CategoryRepository)
	}{
		"Success": {
			want:    categories,
			status:  200,
			message: "Successfully obtained categories",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", pagination).Return(categories, 1, nil)
			},
		},
		"Not Found": {
			want:    nil,
			status:  200,
			message: "no categories found",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no categories found"})
			},
		},
		"Conflict": {
			want:    nil,
			status:  400,
			message: "conflict",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			want:    nil,
			status:  400,
			message: "invalid",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			want:    nil,
			status:  500,
			message: "internal",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			mock := &mocks.CategoryRepository{}
			test.mock(mock)

			t.RequestAndServe("GET", "/categories", "/categories", nil, func(g *gin.Context) {
				t.Setup(mock).List(g)
			})

			t.RunT(test.want, test.status, test.message)
		})
	}
}