// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// UsersTestSuite defines the helper used for user
// testing.
type UsersTestSuite struct {
	test.DBSuite
}

// TestCategories
//
// Assert testing has begun.
func TestCategories(t *testing.T) {
	suite.Run(t, &UsersTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock roles database
// for testing.
func (t *UsersTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&store.Config{
		Driver: t.Driver,
	})
}

const (
	// The default category ID used for testing.
	userID = "1"
)

var (
	// The default user used for testing.
	user = domain.User{
		UserPart: domain.UserPart{
			Id:        1,
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
		},
		Token: "token",
	}
	// The default user create used for testing.
	userCreate = domain.UserCreate{
		User: domain.User{
			UserPart: domain.UserPart{
				FirstName: "Verbis",
				LastName:  "CMS",
				Email:     "verbis@verbiscms.com",
			},
		},
		Password:        "password",
		ConfirmPassword: "password",
	}
	// The default users used for testing.
	users = domain.Users{
		{
			UserPart: domain.UserPart{
				Id: 1, FirstName: "Verbis", LastName: "CMS",
			},
		},
		{
			UserPart: domain.UserPart{
				Id: 2, FirstName: "Verbis", LastName: "CMS",
			},
		},
	}
)
