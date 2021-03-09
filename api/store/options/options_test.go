// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// OptionsTestSuite defines the helper used for role
// testing.
type OptionsTestSuite struct {
	test.DBSuite
}

// TestOptions
//
// Assert testing has begun.
func TestOptions(t *testing.T) {
	suite.Run(t, &OptionsTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock options database
// for testing.
func (t *OptionsTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	if mf != nil {
		mf(t.Mock)
	}
	return New(t.DB)
}
