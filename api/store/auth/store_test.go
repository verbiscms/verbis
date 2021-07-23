// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// AuthTestSuite defines the helper used for role
// testing.
type AuthTestSuite struct {
	test.DBSuite
}

// TestAuth
//
// Assert testing has begun.
func TestAuth(t *testing.T) {
	suite.Run(t, &AuthTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock auth database
// for testing.
func (t *AuthTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&config.Config{
		Driver: t.Driver,
	})
}

var (
	// The default user used for testing.
	user = domain.User{
		UserPart: domain.UserPart{
			Id:        1,
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
			Role: domain.Role{
				Name: "Role",
			},
		},
		Token: "token",
	}
	// The default password reset used for
	// testing.
	passwordReset = domain.PasswordReset{
		Id:    1,
		Email: user.Email,
		Token: user.Token,
	}
)
