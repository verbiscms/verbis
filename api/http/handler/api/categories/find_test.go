// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *CategoriesTestSuite) TestCategories_Find() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.CategoryRepository)
		url     string
	}{
		"Success": {
			category,
			200,
			"Successfully obtained category with ID: 123",
			func(m *mocks.CategoryRepository) {
				m.On("GetByID", 123).Return(category, nil)
			},
			"/categories/123",
		},
		"Invalid ID": {
			nil,
			400,
			"Pass a valid number to obtain the category by ID",
			func(m *mocks.CategoryRepository) {
				m.On("GetByID", 123).Return(domain.Category{}, fmt.Errorf("error"))
			},
			"/categories/wrongid",
		},
		"Not Found": {
			nil,
			200,
			"no categories found",
			func(m *mocks.CategoryRepository) {
				m.On("GetByID", 123).Return(domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "no categories found"})
			},
			"/categories/123",
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			func(m *mocks.CategoryRepository) {
				m.On("GetByID", 123).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/categories/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/categories/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Find(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
