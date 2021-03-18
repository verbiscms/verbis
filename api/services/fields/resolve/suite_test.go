// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
)

// ResolverTestSuite defines the helper used for resolver
// field testing.
type ResolverTestSuite struct {
	suite.Suite
}

// TestResolver
//
// Assert testing has begun.
func TestResolver(t *testing.T) {
	suite.Run(t, new(ResolverTestSuite))
}

// SetupSuite
//
// Discard the logger on setup.
func (t *ResolverTestSuite) SetupSuite() {
	logger.Init(&environment.Env{})
	logger.SetOutput(ioutil.Discard)
}

// GetValue
//
// Returns a default value.
func (t *ResolverTestSuite) GetValue() *Value {
	return &Value{
		&deps.Deps{
			Store: &models.Store{},
		},
	}
}
