// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/sys"
	"github.com/verbiscms/verbis/api/test"
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
func (t *SystemTestSuite) Setup(mf func(m *mocks.System)) *System {
	m := &mocks.System{}
	if mf != nil {
		mf(m)
	}
	d := &deps.Deps{
		System: m,
	}
	return New(d)
}

var (
	// The default install database used for testing.
	preflight = domain.InstallPreflight{
		DbHost:     "host",
		DbPort:     "port",
		DbDatabase: "database",
		DbUser:     "user",
		DbPassword: "password",
	}
	// The default install database with wrong validation
	// used for testing.
	preflightBadValidation = domain.InstallPreflight{
		DbPort:     "port",
		DbDatabase: "database",
		DbUser:     "user",
		DbPassword: "password",
	}
	// The default install verbis used for testing.
	install = domain.InstallVerbis{
		DbHost:              "host",
		DbPort:              "port",
		DbDatabase:          "database",
		DbUser:              "user",
		DbPassword:          "password",
		SiteTitle:           "title",
		SiteUrl:             "http://127.0.0.1",
		UserFirstName:       "verbis",
		UserLastName:        "cms",
		UserEmail:           "hello@verbiscms.com",
		UserPassword:        "password",
		UserConfirmPassword: "password",
		Robots:              false,
	}
	// The default install verbis with wrong validation
	// used for testing.
	installBadValidation = domain.InstallVerbis{
		DbHost:              "host",
		DbPort:              "port",
		DbDatabase:          "database",
		DbUser:              "user",
		DbPassword:          "password",
		SiteUrl:             "http://127.0.0.1",
		UserFirstName:       "verbis",
		UserLastName:        "cms",
		UserEmail:           "hello@verbiscms.com",
		UserPassword:        "password",
		UserConfirmPassword: "password",
		Robots:              false,
	}
)
