// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	suite "github.com/ainsleyclark/verbis/api/http/handler/api/test"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"testing"
)

// getCategoriesMock is a helper to obtain a mock categories controller
// for testing.
func getCategoriesMock(m models.CategoryRepository) *api.Categories {
	return &api.Categories{
		Deps: &deps.Deps{
			Store: &models.Store{
				Categories: m,
			},
		},
	}
}

func TestRedirects_Find(t *testing.T) {

	category := domain.Category{Id: 123, Slug: "/cat", Name: "Category"}

	tt := map[string]struct {
		want    domain.Category
		status  int
		message string
		mock    func(m *mocks.CategoryRepository)
		url     string
	}{
		"Success": {
			want:   category,
			status:  200,
			message: "Successfully obtained category with ID: 123",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 123).Return(category, nil)
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
				getCategoriesMock(mock).GetById(g)
			})

			rr.Run(&test.want, &domain.Category{}, test.status, test.message)
		})
	}
}