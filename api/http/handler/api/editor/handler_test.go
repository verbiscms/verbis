// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package editor

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/store/roles"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// RolesTestSuite defines the helper used for roles
// testing.
type RolesTestSuite struct {
	test.HandlerSuite
}

// TestRoles
//
// Assert testing has begun.
func TestRoles(t *testing.T) {
	suite.Run(t, &RolesTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock roles handler
// for testing.
func (t *RolesTestSuite) Setup(mf func(m *mocks.Repository)) *Roles {
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}
	d := &deps.Deps{
		Store: &store.Repository{
			Roles: m,
		},
	}
	return New(d)
}

var (
	// The default roles used for testing.
	roles = domain.Roles{
		{
			ID:          1,
			Name:        "Banned",
			Description: "Banned Role",
		},
		{
			ID:          2,
			Name:        "Administrator",
			Description: "Administrator Role",
		},
	}
)
