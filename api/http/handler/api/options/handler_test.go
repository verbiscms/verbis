// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/store/options"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// OptionsTestSuite defines the helper used for option
// testing.
type OptionsTestSuite struct {
	test.HandlerSuite
}

// TestOptions
//
// Assert testing has begun.
func TestOptions(t *testing.T) {
	suite.Run(t, &OptionsTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock options handler
// for testing.
func (t *OptionsTestSuite) Setup(mf func(m *mocks.Repository)) *Options {
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}
	d := &deps.Deps{
		Store: &store.Repository{
			Options: m,
		},
	}
	return New(d)
}

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
	// The default options with wrong validation used for testing.
	optionsBadValidation = domain.Options{
		SiteTitle:        "test",
		SiteDescription:  "test",
		SiteLogo:         "test",
		GeneralLocale:    "test",
		ActiveTheme:      "test",
		MediaCompression: 10,
	}
	// The default options used for testing.
	options = domain.OptionsDBMap{
		"test1": domain.OptionDB{
			ID:   123,
			Name: "test",
		},
		"test2": domain.OptionDB{
			ID:   124,
			Name: "test1",
		},
	}
)
