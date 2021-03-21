// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// FormsTestSuite defines the helper used for
// form testing.
type FormsTestSuite struct {
	test.DBSuite
}

// TestForms
//
// Assert testing has begun.
func TestForms(t *testing.T) {
	suite.Run(t, &FormsTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock forms database
// for testing.
func (t *FormsTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&store.Config{
		Driver: t.Driver,
	})
}

const (
	// The default form ID used for testing.
	formID = "1"
)

var (
	// The default form used for testing.
	form = domain.Form{
		Id:   1,
		Name: "Form",
	}
	// The default forms used for testing.
	formsTest = domain.Forms{
		{
			Id:   1,
			Name: "Form",
		},
		{
			Id:   2,
			Name: "Form1",
		},
	}
	// The default forms used for testing.
	formsTestFields = domain.Forms{
		{
			Id:     1,
			Name:   "Form",
			Fields: formFields,
		},
		{
			Id:   2,
			Name: "Form1",
		},
	}
	// The default form fields used for testing.
	formFields = domain.FormFields{
		domain.FormField{
			Key:   "key",
			Label: "label",
			Type:  "type",
		},
	}
)
