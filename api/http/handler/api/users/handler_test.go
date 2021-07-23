// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/gin-gonic/gin/binding"
	pkgValidate "github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	posts "github.com/verbiscms/verbis/api/mocks/store/posts"
	mocks "github.com/verbiscms/verbis/api/mocks/store/users"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// UsersTestSuite defines the helper used for user
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
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock user handler
// for testing.
func (t *UsersTestSuite) Setup(mf func(m *mocks.Repository)) *Users {
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}

	pm := &posts.Repository{}
	pm.On("List", mock.Anything, mock.Anything, mock.Anything).Return(domain.PostData{}, 2, nil)

	if v, ok := binding.Validator.Engine().(*pkgValidate.Validate); ok {
		err := v.RegisterValidation("password", func(fl pkgValidate.FieldLevel) bool {
			return true
		})
		t.NoError(err)
	}

	d := &deps.Deps{
		Store: &store.Repository{
			Posts: pm,
			User:  m,
		},
	}

	return New(d)
}

var (
	// The default user used for testing.
	user = domain.User{
		UserPart: domain.UserPart{
			Id:        123,
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
			Role: domain.Role{
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
				Role: domain.Role{
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
)
