// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// FieldsTestSuite defines the helper used for form
// field testing.
type FieldsTestSuite struct {
	test.DBSuite
}

// TestFields
//
// Assert testing has begun.
func TestFields(t *testing.T) {
	suite.Run(t, &FieldsTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock form fields database
// for testing.
func (t *FieldsTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&config.Config{
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
	// The default form field used for testing.
	formField = domain.FormField{
		Id:    1,
		Key:   "key",
		Label: "label",
		Type:  "type",
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
