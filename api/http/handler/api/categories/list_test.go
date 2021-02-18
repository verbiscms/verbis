// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	suite "github.com/ainsleyclark/verbis/api/test"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRedirects_Find(t *testing.T) {

	category := domain.Category{Id: 123, Slug: "/cat", Name: "Category"}

	tt := map[string]struct {
		want    *domain.Category
		status  int
		message string
		mock    func(m *mocks.CategoryRepository)
		url     string
	}{
		"Success": {
			want:   &category,
			status:  200,
			message: "Successfully obtained category with ID: 123",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 123).Return(category, nil)
			},
			url: "/categories/123",
		},
		"Invalid ID": {
			want:    nil,
			status:  400,
			message: "Pass a valid number to obtain the category by ID",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 123).Return(domain.Category{}, fmt.Errorf("error"))
			},
			url: "/categories/wrongid",
		},
		"Not Found": {
			want:    nil,
			status:  200,
			message: "no categories found",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 123).Return(domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "no categories found"})
			},
			url: "/categories/123",
		},
		"Internal Error": {
			want:    nil,
			status:  500,
			message: "internal",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 123).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/categories/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			rr := suite.APITestSuite(t)
			mock := &mocks.CategoryRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", test.url, "/categories/:id", nil, func(g *gin.Context) {
				//Mock(mock).Find(g)
			})

			rr.Run(&domain.Category{}, test.want, test.status, test.message)
		})
	}
}