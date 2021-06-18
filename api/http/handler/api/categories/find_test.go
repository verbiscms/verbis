// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/categories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *CategoriesTestSuite) TestCategories_Find() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			category,
			http.StatusOK,
			"Successfully obtained category with ID: 123",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(category, nil)
			},
			"/categories/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"Pass a valid number to obtain the category by ID",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(domain.Category{}, fmt.Errorf("error"))
			},
			"/categories/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"no categories found",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "no categories found"})
			},
			"/categories/123",
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
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
