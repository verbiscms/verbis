// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/suite"
	"testing"
)

// FormsTestSuite defines the helper used for form
// testing.
type FormsTestSuite struct {
	api.HandlerSuite
}

// TestForms
//
// Assert testing has begun.
func TestForms(t *testing.T) {
	suite.Run(t, &FormsTestSuite{
		HandlerSuite: api.TestSuite(),
	})
}

// Setup
//
// A helper to obtain a mock form handler
// for testing.
func (t *FormsTestSuite) Setup(mf func(m *mocks.FormRepository)) *Forms {
	m := &mocks.FormRepository{}
	if mf != nil {
		mf(m)
	}
	return &Forms{
		Deps: &deps.Deps{
			Store: &models.Store{
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
	forms = []domain.Form{
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
	// The default pagination used for testing.
	pagination = params.Params{
		Page:           api.DefaultParams.Page,
		Limit:          15,
		OrderBy:        api.DefaultParams.OrderBy,
		OrderDirection: api.DefaultParams.OrderDirection,
		Filters:        nil,
	}
)
