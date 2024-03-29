// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/store/categories"
	"net/http"
)

func (t *CategoriesTestSuite) TestCategories_Update() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			category,
			http.StatusOK,
			"Successfully updated category with ID: 123",
			category,
			func(m *mocks.Repository) {
				m.On("Update", category).Return(category, nil)
			},
			"/categories/123",
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "slug", Message: "Slug is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			categoryBadValidation,
			func(m *mocks.Repository) {
				m.On("Update", categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
			"/categories/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to update the category",
			category,
			func(m *mocks.Repository) {
				m.On("Update", categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
			"/categories/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			category,
			func(m *mocks.Repository) {
				m.On("Update", category).Return(domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/categories/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			category,
			func(m *mocks.Repository) {
				m.On("Update", category).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
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
