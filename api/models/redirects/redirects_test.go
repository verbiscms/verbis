// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
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
	if mf != nil {
		mf(t.Mock)
	}
	return New(t.DB)
}

const (
	// The default redirect ID used for testing.
	redirectID = "1"
)

var (
	// The default redirect used for testing.
	redirect = domain.Redirect{
		Id:   1,
		From: "/from",
		To:   "/to",
		Code: 301,
	}
	// The default redirects used for testing.
	redirects = domain.Redirects{
		{
			Id:   1,
			From: "/from",
			To:   "/to",
			Code: 301,
		},
		{
			Id:   2,
			From: "/from",
			To:   "/to",
			Code: 301,
		},
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
