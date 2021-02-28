// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/deps"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// FieldTestSuite defines the helper used for cache
// testing.
type FieldTestSuite struct {
	test.HandlerSuite
}

// TestFields
//
// Assert testing has begun.
func TestFields(t *testing.T) {
	suite.Run(t, &FieldTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock fields handler
// for testing.
func (t *FieldTestSuite) Setup(mf func(m *mocks.FieldsRepository)) *Fields {
	m := &mocks.FieldsRepository{}
	if mf != nil {
		mf(m)
	}
	return &Fields{
		Deps: &deps.Deps{
			Store: &models.Store{
				Fields: m,
			},
		},
	}
}
