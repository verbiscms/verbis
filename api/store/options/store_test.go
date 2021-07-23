// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
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
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&config.Config{
		Driver: t.Driver,
	})
}

const (
	// The default option name used for testing.
	optionName = "name"
	// The default option value used for testing.
	optionValue = "test"
)

var (
	// The default option struct used for testing.
	optionsStruct = domain.Options{
		SiteTitle:        "test",
		SiteDescription:  "test",
		SiteLogo:         "test",
		SiteUrl:          "http://verbiscms.com",
		ActiveTheme:      "theme",
		GeneralLocale:    "test",
		MediaCompression: 10,
	}
	// The default options used for testing.
	options = domain.OptionsDBMap{
		"name": "test",
	}
)
