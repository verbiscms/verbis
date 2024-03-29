// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// UsersTestSuite defines the helper used for user
// testing.
type UsersTestSuite struct {
	test.DBSuite
}

// TestUsers
//
// Assert testing has begun.
func TestUsers(t *testing.T) {
	suite.Run(t, &UsersTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock users database
// for testing.
func (t *UsersTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&config.Config{
		Driver: t.Driver,
	})
}

const (
	// The default user ID used for testing.
	userID = "1"
	// The select statement.
	SelectStatement = "SELECT users.*, roles.id `roles.id`, roles.name `roles.name`, roles.description `roles.description` FROM `users` LEFT JOIN `user_roles` AS `user_roles` ON `users`.`id` = `user_roles`.`user_id` LEFT JOIN `roles` AS `roles` ON `user_roles`.`role_id` = `roles`.`id` "
)

var (
	// The default user used for testing.
	user = domain.User{
		UserPart: domain.UserPart{
			ID:        1,
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
			Role: domain.Role{
				Name: "Role",
			},
		},
		Token: "token",
	}
	// The default user create used for testing.
	userCreate = domain.UserCreate{
		User: domain.User{
			UserPart: domain.UserPart{
				ID:        1,
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
				ID: 1, FirstName: "Verbis", LastName: "CMS", Role: domain.Role{
					Name: "Role",
				},
			},
		},
		{
			UserPart: domain.UserPart{
				ID: 2, FirstName: "Verbis", LastName: "CMS", Role: domain.Role{
					Name: "Role",
				},
			},
		},
	}
	// The default reset password used for testing.
	passwordReset = domain.UserPasswordReset{
		NewPassword:     "password1",
		ConfirmPassword: "password1",
	}
)
