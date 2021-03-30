// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/forms"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// FormsTestSuite defines the helper used for form
// testing.
type FormsTestSuite struct {
	test.HandlerSuite
}

// TestForms
//
// Assert testing has begun.
func TestForms(t *testing.T) {
	suite.Run(t, &FormsTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock form handler
// for testing.
func (t *FormsTestSuite) Setup(mf func(m *mocks.Repository)) *Forms {
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}
	return &Forms{
		Deps: &deps.Deps{
			Store: &store.Repository{
				Forms: m,
			},
		},
	}
}

// The dynamic struct to be validated.
type body struct {
	Name string `binding:"required"`
}

var (
	// The default form used for testing.
	form = domain.Form{
		Id:   123,
		Name: "Form",
		Fields: domain.FormFields{
			{
				Key:      "key",
				Label:    "label",
				Type:     "text",
				Required: true,
			},
		},
	}
	// The default form with body used for testing.
	formBody = domain.Form{
		Id:   123,
		Name: "Form",
		Fields: domain.FormFields{
			{
				Key:      "key",
				Label:    "label",
				Type:     "text",
				Required: true,
			},
		},
		Body: body{Name: "test"},
	}
	// The default form with wrong validation used for testing.
	formBadValidation = domain.Form{
		Id: 123,
		Body: struct {
			Name string `binding:"required"`
		}{},
	}
	// The default forms used for testing.
	forms = domain.Forms{
		{
			Id:   123,
			Name: "Form",
			Body: body{Name: "test"},
		},
		{
			Id:   124,
			Name: "Form1",
			Body: body{Name: "test"},
		},
	}
	// The default params used for testing.
	defaultParams = params.Params{
		Page:           api.DefaultParams.Page,
		Limit:          15,
		OrderBy:        api.DefaultParams.OrderBy,
		OrderDirection: api.DefaultParams.OrderDirection,
		Filters:        nil,
	}
)
