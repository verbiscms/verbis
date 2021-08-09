// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/store/categories"
	posts "github.com/verbiscms/verbis/api/mocks/store/posts"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/test"
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
		ID:       123,
		Slug:     "/cat",
		Name:     "Category",
		Resource: "test",
	}
	// The default category with wrong validation used for testing.
	categoryBadValidation = domain.Category{
		ID:       123,
		Name:     "Category",
		Resource: "test",
	}
	// The default categories used for testing.
	categories = domain.Categories{
		{
			ID:   123,
			Slug: "/cat",
			Name: "Category",
		},
		{
			ID:   124,
			Slug: "/cat1",
			Name: "Category1",
		},
	}
)
