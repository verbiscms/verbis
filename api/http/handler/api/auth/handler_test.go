// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/logger"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	events "github.com/verbiscms/verbis/api/mocks/events"
	mocks "github.com/verbiscms/verbis/api/mocks/store/auth"
	users "github.com/verbiscms/verbis/api/mocks/store/users"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/test"
	"io/ioutil"
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
// A helper to obtain a mock auth handler
// for testing.
func (t *AuthTestSuite) Setup(mf func(m *mocks.Repository)) *Auth {
	logger.SetOutput(ioutil.Discard)
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}
	d := &deps.Deps{
		Store: &store.Repository{
			Auth: m,
		},
		Env:     &environment.Env{},
		Options: &domain.Options{},
	}
	return New(d)
}

// Setup
//
// A helper to obtain a mock auth handler
// for testing with cache.
func (t *AuthTestSuite) SetupCache(mf func(m *mocks.Repository, c *cache.Store)) *Auth {
	logger.SetOutput(ioutil.Discard)
	m := &mocks.Repository{}
	c := &cache.Store{}
	if mf != nil {
		mf(m, c)
	}
	d := &deps.Deps{
		Store: &store.Repository{
			Auth: m,
		},
		Cache:   c,
		Env:     &environment.Env{},
		Options: &domain.Options{},
	}
	return New(d)
}

// Setup
//
// A helper to obtain a mock categories handler
// for testing.
func (t *AuthTestSuite) SetupDispatcher(mf func(m *mocks.Repository, c *cache.Store, u *users.Repository), ms func(m *events.Dispatcher)) *Auth {
	logger.SetOutput(ioutil.Discard)
	m := &mocks.Repository{}
	c := &cache.Store{}
	u := &users.Repository{}
	if mf != nil {
		mf(m, c, u)
	}
	d := &deps.Deps{
		Store: &store.Repository{
			Auth: m,
			User: u,
		},
		Cache:   c,
		Env:     &environment.Env{},
		Options: &domain.Options{},
	}

	mr := &events.Dispatcher{}
	if ms != nil {
		ms(mr)
	}
	a := New(d)
	a.resetPassword = mr
	return a
}

var (
	// The default user used for testing
	user = domain.User{
		UserPart: domain.UserPart{
			ID:        1,
			FirstName: "verbis",
			LastName:  "cms",
			Email:     "hello@verbiscms.com",
		},
		Password:      "",
		Token:         "",
		TokenLastUsed: nil,
	}
)
