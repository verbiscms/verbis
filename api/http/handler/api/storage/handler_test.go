// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/services/storage"
	"github.com/verbiscms/verbis/api/services/storage"
	"github.com/verbiscms/verbis/api/test"
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
		domain.Bucket{ID: "1", Name: "verbis"},
		domain.Bucket{ID: "2", Name: "verbis"},
	}
	// The default storageConfiguration used for
	// testing.
	storageConfig = storage.Configuration{
		Info: domain.StorageConfig{
			Provider: "test",
			Bucket:   "test",
		},
	}
	// The default storageChange used for testing.
	storageChange = domain.StorageConfig{
		Provider: domain.StorageAWS,
		Bucket:   "verbis-bucket",
		Region:   "",
	}
	// The default storage change with wrong
	// validation used for testing.
	storageChangeBadValidation = domain.StorageConfig{}
)
