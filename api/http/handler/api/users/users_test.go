// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/gin-gonic/gin/binding"
	pkgValidate "github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
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
	cache.Init()
	suite.Run(t, &UsersTestSuite{
		HandlerSuite: test.TestSuite(),
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
	pm := &mocks.PostsRepository{}
	pm.On("Get", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(domain.PostData{}, 2, nil)

	if v, ok := binding.Validator.Engine().(*pkgValidate.Validate); ok {
		err := v.RegisterValidation("password", func(fl pkgValidate.FieldLevel) bool {
			return true
		})
		t.NoError(err)
	}

	return &Users{
		Deps: &deps.Deps{
			Store: &models.Store{
				Posts: pm,
				User:  m,
			},
		},
	}
}

var (
	// The default user used for testing.
	user = domain.User{
		UserPart: domain.UserPart{
			Id:        123,
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
			Role: domain.UserRole{
				Id: 1,
			},
		},
	}
	// The default user with wrong validation used for testing.
	userBadValidation = domain.User{
		UserPart: domain.UserPart{
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
		},
	}
	// The default user create used for testing.
	userCreate = domain.UserCreate{
		User: domain.User{
			UserPart: domain.UserPart{
				FirstName: "Verbis",
				LastName:  "CMS",
				Email:     "verbis@verbiscms.com",
				Role: domain.UserRole{
					Id: 123,
				},
			},
		},
		Password:        "password",
		ConfirmPassword: "password",
	}
	// The default user create with wrong validation used for testing.
	userCreateBadValidation = domain.UserCreate{
		User: domain.User{
			UserPart: domain.UserPart{
				FirstName: "Verbis",
				LastName:  "CMS",
				Email:     "verbis@verbiscms.com",
			},
		},
		Password:        "password",
		ConfirmPassword: "password",
	}
	// The default users used for testing.
	users = domain.Users{
		{
			UserPart: domain.UserPart{
				Id: 123, FirstName: "Verbis", LastName: "CMS",
			},
		},
		{
			UserPart: domain.UserPart{
				Id: 123, FirstName: "Verbis", LastName: "CMS",
			},
		},
	}
	// The default roles used for testing.
	roles = []domain.UserRole{
		{
			Id:          1,
			Name:        "Banned",
			Description: "Banned Role",
		},
		{
			Id:          2,
			Name:        "Administrator",
			Description: "Administrator Role",
		},
	}
	// The default reset password used for testing.
	reset = domain.UserPasswordReset{
		DBPassword:      "",
		CurrentPassword: "password",
		NewPassword:     "verbiscms",
		ConfirmPassword: "verbiscms",
	}
	// The default reset password with wrong validation used for testing.
	resetBadValidation = domain.UserPasswordReset{
		CurrentPassword: "password",
		NewPassword:     "verbiscms",
		ConfirmPassword: "verbiscmss",
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
