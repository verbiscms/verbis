// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// SiteTestSuite defines the helper used for site
// testing.
type SiteTestSuite struct {
	suite.Suite
}

// TestSite
//
// Assert testing has begun.
func TestSite(t *testing.T) {
	suite.Run(t, &SiteTestSuite{})
}
