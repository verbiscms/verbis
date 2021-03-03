// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

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

func (t *CategoriesTestSuite) TestCategories_Update() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.CategoryRepository)
		url     string
	}{
		"Success": {
			category,
			http.StatusOK,
			"Successfully updated category with ID: 123",
			category,
			func(m *mocks.CategoryRepository) {
				m.On("Update", &category).Return(category, nil)
			},
			"/categories/123",
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "slug", Message: "Slug is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			categoryBadValidation,
			func(m *mocks.CategoryRepository) {
				m.On("Update", categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
			"/categories/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to update the category",
			category,
			func(m *mocks.CategoryRepository) {
				m.On("Update", categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
			"/categories/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			category,
			func(m *mocks.CategoryRepository) {
				m.On("Update", &category).Return(domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/categories/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			category,
			func(m *mocks.CategoryRepository) {
				m.On("Update", &category).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/categories/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPut, test.url, "/categories/:id", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Update(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
