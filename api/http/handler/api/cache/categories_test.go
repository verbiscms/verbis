// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// CacheTestSuite defines the helper used for cache
// testing.
type CacheTestSuite struct {
	test.HandlerSuite
}

// TestCache
//
// Assert testing has begun.
func TestCache(t *testing.T) {
	suite.Run(t, &CacheTestSuite{
		HandlerSuite: test.TestSuite(),
	})
}

// Setup
//
// A helper to obtain a mock cache handler
// for testing.
func (t *CacheTestSuite) Setup() {
	cache.Init()
}
