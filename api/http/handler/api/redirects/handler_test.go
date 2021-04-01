// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/redirects"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// RedirectsTestSuite defines the helper used for redirect
// testing.
type RedirectsTestSuite struct {
	test.HandlerSuite
}

// TestRedirects
//
// Assert testing has begun.
func TestRedirects(t *testing.T) {
	suite.Run(t, &RedirectsTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock redirects handler
// for testing.
func (t *RedirectsTestSuite) Setup(mf func(m *mocks.Repository)) *Redirects {
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}
	d := &deps.Deps{
		Store: &store.Repository{
			Redirects: m,
		},
	}
	return New(d)
}

var (
	// The default redirect used for testing.
	redirect = domain.Redirect{
		Id:   123,
		From: "/test",
		To:   "/testing",
		Code: 301,
	}
	// The default redirect with wrong validation used for testing.
	redirectBadValidation = domain.Redirect{
		Id:   123,
		From: "/test",
		Code: 301,
	}
	// The default redirects used for testing.
	redirects = domain.Redirects{
		{
			Id:   123,
			From: "/test",
			To:   "/testing",
		},
		{
			Id:   124,
			From: "/test1",
			To:   "/testing2",
		},
	}
)
