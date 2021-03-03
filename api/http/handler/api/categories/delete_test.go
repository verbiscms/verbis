// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
			nil,
			http.StatusOK,
			"Successfully deleted category with ID: 123",
			func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(nil)
			},
			"/categories/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to delete a category",
			func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(nil)
			},
			"/categories/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/categories/123",
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			"/categories/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/categories/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodDelete, test.url, "/categories/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Delete(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
