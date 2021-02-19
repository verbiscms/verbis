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
	"github.com/gin-gonic/gin"
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
			want:    category,
			status:  200,
			message: "Successfully updated category with ID: 123",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Update", &category).Return(category, nil)
			},
			url: "/categories/123",
		},
		"Validation Failed": {
			want:    `{"errors":[{"key":"slug","message":"Slug is required.","type":"required"}]}`,
			status:  400,
			message: "Validation failed",
			input:   categoryBadValidation,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Update", categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
			url: "/categories/123",
		},
		"Invalid ID": {
			want:    nil,
			status:  400,
			message: "A valid ID is required to update the category",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Update", categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
			url: "/categories/wrongid",
		},
		"Not Found": {
			want:    nil,
			status:  400,
			message: "not found",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Update", &category).Return(domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/categories/123",
		},
		"Internal": {
			want:    nil,
			status:  500,
			message: "internal",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Update", &category).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/categories/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			mock := &mocks.CategoryRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Error(err)
			}

			t.RequestAndServe("PUT", test.url, "/categories/:id", bytes.NewBuffer(body), func(g *gin.Context) {
				t.Setup(mock).Update(g)
			})

			t.RunT(test.want, test.status, test.message)
		})
	}
}