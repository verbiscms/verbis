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
	suite "github.com/ainsleyclark/verbis/api/test"
	"github.com/gin-gonic/gin"
)

var (
	category = domain.Category{Id: 123, Slug: "/cat", Name: "Category", Resource: "test"}
	categoryBadValidation = domain.Category{Id: 123, Name: "Category", Resource: "test"}
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
			want:    "{}",
			status:  400,
			message: "invalid",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			want:    "{}",
			status:  400,
			message: "conflict",
			input:   category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			want:    "{}",
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
			rr := suite.APITestSuite(t.T())
			mock := &mocks.CategoryRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Error(err)
			}

			rr.RequestAndServe("POST", "/categories", "/categories", bytes.NewBuffer(body), func(g *gin.Context) {
				t.Mock(mock).Create(g)
			})

			respond, data := rr.TestRun()

			t.Equal(test.message, respond.Message)
			t.Equal(test.status, rr.Status())
			t.Equal(suite.JsonHeader, rr.ContentType())
			t.JSONEq(rr.TestIn(test.want), data)
		})
	}
}