// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/ainsleyclark/verbis/api/test/testdata/updates"
	"github.com/stretchr/testify/suite"
	"testing"
)

// InternalTestSuite defines the helper used for
// category testing.
type InternalTestSuite struct {
	test.DBSuite
}

// TestInternal - Assert testing has begun.
func TestInternal(t *testing.T) {
	suite.Run(t, &InternalTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup A helper to obtain a mock migration
// database for testing.
func (t *InternalTestSuite) Setup(mf func(m sqlmock.Sqlmock)) Migrator {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return &migrate{
		Down:   nil,
		Driver: MySQLDriver,
		DB:     t.DB,
		Embed:  updates.Static,
	}
}
