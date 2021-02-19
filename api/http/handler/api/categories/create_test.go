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

func (t *CategoriesTestSuite) TestCategories_Create() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.CategoryRepository)
	}{
		"Success": {
			&category,
			200,
			"Successfully created category with ID: 123",
			category,
			func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(category, nil)
			},
		},
		"Validation Failed": {
			api.ValidationErrJson{Errors: validation.Errors{{Key: "slug", Message: "Slug is required.", Type: "required"}}},
			400,
			"Validation failed",
			categoryBadValidation,
			func(m *mocks.CategoryRepository) {
				m.On("Create", &categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			nil,
			400,
			"invalid",
			category,
			func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			400,
			"conflict",
			category,
			func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			category,
			func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/categories", "/categories", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Create(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}