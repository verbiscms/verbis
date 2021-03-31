// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/categories"
	posts "github.com/ainsleyclark/verbis/api/mocks/store/posts"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

// CategoriesTestSuite defines the helper used for category
// testing.
type CategoriesTestSuite struct {
	test.HandlerSuite
}

// TestCategories
//
// Assert testing has begun.
func TestCategories(t *testing.T) {
	suite.Run(t, &CategoriesTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock categories handler
// for testing.
func (t *CategoriesTestSuite) Setup(mf func(m *mocks.Repository)) *Categories {
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}

	pm := &posts.Repository{}
	pm.On("List", mock.Anything, mock.Anything, mock.Anything).Return(domain.PostData{}, 0, nil)

	d := &deps.Deps{
		Store: &store.Repository{
			Categories: m,
			Posts:      pm,
		},
	}

	return New(d)
}

var (
	// The default category used for testing.
	category = domain.Category{
		Id:       123,
		Slug:     "/cat",
		Name:     "Category",
		Resource: "test",
	}
	// The default category with wrong validation used for testing.
	categoryBadValidation = domain.Category{
		Id:       123,
		Name:     "Category",
		Resource: "test",
	}
	// The default categories used for testing.
	categories = domain.Categories{
		{
			Id:   123,
			Slug: "/cat",
			Name: "Category",
		},
		{
			Id:   124,
			Slug: "/cat1",
			Name: "Category1",
		},
	}
	// The default params used for testing.
	defaultParams = params.Params{
		Page:           api.DefaultParams.Page,
		Limit:          15,
		OrderBy:        api.DefaultParams.OrderBy,
		OrderDirection: api.DefaultParams.OrderDirection,
		Filters:        nil,
	}
)
