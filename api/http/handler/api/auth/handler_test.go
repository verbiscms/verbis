// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	events "github.com/ainsleyclark/verbis/api/mocks/events"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/auth"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// AuthTestSuite defines the helper used for auth
// testing.
type AuthTestSuite struct {
	test.HandlerSuite
}

// TestAuth
//
// Assert testing has begun.
func TestAuth(t *testing.T) {
	suite.Run(t, &AuthTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock categories handler
// for testing.
func (t *AuthTestSuite) Setup(mf func(m *mocks.Repository)) *Auth {
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}

	mr := events.Dispatcher{}
	mr.On("")

	return &Auth{
		Deps: &deps.Deps{
			Store: &store.Repository{
				Auth: m,
			},
		},
		resetPassword: nil,
	}
}

var (
	// The default user used for testing
	user = domain.User{
		UserPart: domain.UserPart{
			Id:        1,
			FirstName: "verbis",
			LastName:  "cms",
			Email:     "hello@verbiscms.com",
		},
		Password:      "",
		Token:         "",
		TokenLastUsed: nil,
	}
)
