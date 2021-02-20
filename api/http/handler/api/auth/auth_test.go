// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/suite"
	"testing"
)

// AuthTestSuite defines the helper used for auth
// testing.
type AuthTestSuite struct {
	api.HandlerSuite
}

// TestAuth
//
// Assert testing has begun.
func TestAuth(t *testing.T) {
	suite.Run(t, &AuthTestSuite{
		HandlerSuite: api.TestSuite(),
	})
}

// Setup
//
// A helper to obtain a mock categories handler
// for testing.
func (t *AuthTestSuite) Setup(mf func(m *mocks.AuthRepository)) *Auth {
	m := &mocks.AuthRepository{}
	if mf != nil {
		mf(m)
	}
	return &Auth{
		Deps: &deps.Deps{
			Store: &models.Store{
				Auth: m,
			},
		},
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
