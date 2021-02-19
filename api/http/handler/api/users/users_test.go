// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

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

// UsersTestSuite defines the helper used for category
// testing.
type UsersTestSuite struct {
	test.HandlerSuite
}

// TestUsers
//
// Assert testing has begun.
func TestUsers(t *testing.T) {
	suite.Run(t, &UsersTestSuite{
		HandlerSuite: test.APITestSuite(),
	})
}

// Setup
//
// A helper to obtain a mock categories handler
// for testing.
func (t *UsersTestSuite) Setup(mf func(m *mocks.UserRepository)) *Users {
	m := &mocks.UserRepository{}
	if mf != nil {
		mf(m)
	}
	return &Users{
		Deps: &deps.Deps{
			Store: &models.Store{
				User: m,
			},
		},
	}
}

var (
	// The default category used for testing.
	category = domain.Category{
		Id: 123,
		Slug: "/cat",
		Name: "Category",
		Resource: "test",
	}
	// The default category with wrong validation used for testing.
	categoryBadValidation = domain.Category{
		Id: 123,
		Name: "Category",
		Resource: "test",
	}
	// The default categories used for testing.
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
	// The default pagination used for testing.
	pagination = params.Params{
		Page: api.DefaultParams.Page,
		Limit: 15,
		OrderBy: api.DefaultParams.OrderBy,
		OrderDirection: api.DefaultParams.OrderDirection,
		Filters: nil,
	}
)