// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
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
			want:    &category,
			status:  200,
			message: "Successfully created category with ID: 123",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(category, nil)
			},
		},
		"Validation Failed": {
			want:    api.ValidationErrJson{Errors: validation.Errors{{Key: "slug", Message: "Slug is required.", Type: "required"}}},
			status:  400,
			message: "Validation failed",
			input:   categoryBadValidation,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			want:    nil,
			status:  400,
			message: "invalid",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			want:    nil,
			status:  400,
			message: "conflict",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			want:    nil,
			status:  500,
			message: "internal",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			body, err := json.Marshal(test.input)
			if err != nil {
				t.Error(err)
			}

			t.RequestAndServe("POST", "/categories", "/categories", bytes.NewBuffer(body), func(g *gin.Context) {
				t.Setup(test.mock).Create(g)
			})

			t.RunT(test.want, test.status, test.message)
		})
	}
}