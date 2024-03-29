// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// RedirectsTestSuite defines the helper used for role
// testing.
type RedirectsTestSuite struct {
	test.DBSuite
}

// TestRedirects
//
// Assert testing has begun.
func TestRedirects(t *testing.T) {
	suite.Run(t, &RedirectsTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock redirects database
// for testing.
func (t *RedirectsTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&config.Config{
		Driver: t.Driver,
	})
}

const (
	// The default redirect ID used for testing.
	redirectID = "1"
)

var (
	// The default redirect used for testing.
	redirect = domain.Redirect{
		ID:   1,
		From: "/from",
		To:   "/to",
		Code: 301,
	}
	// The default redirects used for testing.
	redirects = domain.Redirects{
		{
			ID:   1,
			From: "/from",
			To:   "/to",
			Code: 301,
		},
		{
			ID:   2,
			From: "/from",
			To:   "/to",
			Code: 301,
		},
	}
)
