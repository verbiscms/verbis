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
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	suite "github.com/ainsleyclark/verbis/api/test"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestCategories_Create(t *testing.T) {

	category := domain.Category{Id: 123, Slug: "/cat", Name: "Category", Resource: "test"}
	categoryBadValidation := domain.Category{Id: 123, Name: "Category", Resource: "test"}

	tt := map[string]struct {
		want    *domain.Category
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
			want:    `{"errors":[{"key":"slug","message":"Slug is required.","type":"required"}]}`,
			status:  400,
			message: "Validation failed",
			input:   categoryBadValidation,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			want:    `{}`,
			status:  400,
			message: "invalid",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			want:    `{}`,
			status:  400,
			message: "conflict",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			rr := suite.APITestSuite(t)
			mock := &mocks.CategoryRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("POST", "/categories", "/categories", bytes.NewBuffer(body), func(g *gin.Context) {
				Mock(mock).Create(g)
			})

			rr.Run(&domain.Category{}, test.want, test.status, test.message)
		})
	}
}