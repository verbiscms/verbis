// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CategoriesTestSuite struct {
	test.HandlerSuite
}

func TestCategories(t *testing.T) {
	suite.Run(t, &CategoriesTestSuite{
		HandlerSuite: test.APITestSuite(),
	})
}

// Setup
//
// TODO: Post mock?!
// A helper to obtain a mock categories handler
// for testing.
func (t *CategoriesTestSuite) Setup(mf func(m *mocks.CategoryRepository)) *Categories {
	mock := &mocks.CategoryRepository{}
	if mf != nil {
		mf(mock)
	}
	return &Categories{
		Deps: &deps.Deps{
			Store: &models.Store{
				Categories: mock,
			},
		},
	}
}

var (
	//
	category = domain.Category{
		Id: 123,
		Slug: "/cat",
		Name: "Category",
		Resource: "test",
	}
	//
	categoryBadValidation = domain.Category{
		Id: 123,
		Name: "Category",
		Resource: "test",
	}
	//
	categories = []domain.Category{
		{
			Id: 123,
			Slug: "/cat",
			Name: "Category",
		},
		{
			Id: 124,
			Slug: "/cat1",
			Name: "Category1",
		},
	}
	//
	pagination = params.Params{
		Page: api.DefaultParams.Page,
		Limit: 15,
		OrderBy: api.DefaultParams.OrderBy,
		OrderDirection: api.DefaultParams.OrderDirection,
		Filters: nil,
	}
)