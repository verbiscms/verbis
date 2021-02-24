// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/suite"
	"testing"
)

// OptionsTestSuite defines the helper used for option
// testing.
type OptionsTestSuite struct {
	api.HandlerSuite
}

// TestOptions
//
// Assert testing has begun.
func TestOptions(t *testing.T) {
	suite.Run(t, &OptionsTestSuite{
		HandlerSuite: api.TestSuite(),
	})
}

// Setup
//
// A helper to obtain a mock options handler
// for testing.
func (t *OptionsTestSuite) Setup(mf func(m *mocks.OptionsRepository)) *Options {
	m := &mocks.OptionsRepository{}
	if mf != nil {
		mf(m)
	}
	return &Options{
		Deps: &deps.Deps{
			Store: &models.Store{
				Options: m,
			},
		},
	}
}

var (
	// The default option struct used for testing.
	optionsStruct = domain.Options{
		SiteTitle:        "test",
		SiteDescription:  "test",
		SiteLogo:         "test",
		SiteUrl:          "http://verbiscms.com",
		GeneralLocale:    "test",
		MediaCompression: 10,
	}
	// The default options with wrong validation used for testing.
	optionsBadValidation = domain.Options{
		SiteTitle:        "test",
		SiteDescription:  "test",
		SiteLogo:         "test",
		GeneralLocale:    "test",
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
