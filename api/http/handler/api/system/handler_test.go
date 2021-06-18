// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/roles"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// SystemTestSuite defines the helper used for sys
// testing.
type SystemTestSuite struct {
	test.HandlerSuite
}

// TestSystem
//
// Assert testing has begun.
func TestSystem(t *testing.T) {
	suite.Run(t, &SystemTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock system and
// updater handler for testing.
func (t *SystemTestSuite) Setup(mf func(m *mocks.Repository)) *System {
	//m := &mocks.Repository{}
	//if mf != nil {
	//	mf(m)
	//}
	//d := &deps.Deps{
	//	Store: &store.Repository{
	//		Roles: m,
	//	},
	//}
	//return New(d)
}
