// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
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
func TestCategories(t *testing.T) {
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
	return New(&store.Config{
		Driver: t.Driver,
	})
}

const (
	// The default category ID used for testing.
	categoryID = "1"
)

var (
// The default category used for testing.

// The default categories used for testing.

)
