// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sizes

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/store/config"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// SizesTestSuite defines the helper used for media
// size testing.
type SizesTestSuite struct {
	test.DBSuite
}

// TestSizes
//
// Assert testing has begun.
func TestSizes(t *testing.T) {
	suite.Run(t, &SizesTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock media sizes database
// for testing.
func (t *SizesTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&config.Config{
		Driver: t.Driver,
	})
}

const (
	// The default media item ID used for testing.
	mediaID = "1"
)

var (
// The default meta used for testing.

)
