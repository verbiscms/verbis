// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/storage"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// StorageTestSuite defines the helper used for sys
// testing.
type StorageTestSuite struct {
	test.HandlerSuite
}

// TestStorage asserts testing has begun.
func TestSystem(t *testing.T) {
	suite.Run(t, &StorageTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup is a helper to obtain a mock storage provider
// for testing.
func (t *StorageTestSuite) Setup(mf func(m *mocks.Provider)) *Storage {
	m := &mocks.Provider{}
	if mf != nil {
		mf(m)
	}
	d := &deps.Deps{
		Storage: m,
	}
	return New(d)
}

var (
	// The default buckets used for testing.
	buckets = domain.Buckets{
		domain.Bucket{Id: "1", Name: "verbis"},
		domain.Bucket{Id: "2", Name: "verbis"},
	}
	// The default storageConfiguration used for
	// testing.
	storageConfig = domain.StorageConfiguration{
		ActiveProvider: "test",
		ActiveBucket:   "test",
	}
)
