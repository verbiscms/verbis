// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	mocks "github.com/verbiscms/verbis/api/mocks/store/forms"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/test"
	"io/ioutil"
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
	d := &deps.Deps{
		Store: &store.Repository{
			Forms: m,
		},
	}
	logger.SetOutput(ioutil.Discard)
	return New(d)
}

// The dynamic struct to be validated.
type body struct {
	Name string `binding:"required"`
}

var (
	// The default form used for testing.
	form = domain.Form{
		ID:   123,
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
		ID:   123,
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
		ID: 123,
		Body: struct {
			Name string `binding:"required"`
		}{},
	}
	// The default forms used for testing.
	forms = domain.Forms{
		{
			ID:   123,
			Name: "Form",
			Body: body{Name: "test"},
		},
		{
			ID:   124,
			Name: "Form1",
			Body: body{Name: "test"},
		},
	}
)
