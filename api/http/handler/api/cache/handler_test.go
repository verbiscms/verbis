// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	mocks "github.com/verbiscms/verbis/api/mocks/cache"
	"github.com/verbiscms/verbis/api/test"
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
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock cache handler
// for testing.
func (t *CacheTestSuite) Setup(mock func(m *mocks.Store)) *Cache {
	m := &mocks.Store{}
	if mock != nil {
		mock(m)
	}
	return New(&deps.Deps{
		Cache: m,
	})
}
