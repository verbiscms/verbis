// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	fields "github.com/verbiscms/verbis/api/mocks/store/forms/fields"
	submissions "github.com/verbiscms/verbis/api/mocks/store/forms/submissions"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
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
func (t *FormsTestSuite) Setup(mf func(m sqlmock.Sqlmock), mfm func(f *fields.Repository, s *submissions.Repository)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}

	f := &fields.Repository{}
	su := &submissions.Repository{}
	if mfm != nil {
		mfm(f, su)
	}

	s := New(&config.Config{
		Driver: t.Driver,
	})
	s.fields = f
	s.submissions = su
	return s
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
		Fields: domain.FormFields{
			domain.FormField{
				Key:   "key",
				Label: "label",
				Type:  "type",
			},
		},
		Submissions: domain.FormSubmissions{
			domain.FormSubmission{
				FormId:    1,
				IPAddress: "127.0.0.1",
				UserAgent: "chrome",
			},
		},
	}
	// The default forms used for testing.
	forms = domain.Forms{
		{
			Id:   1,
			Name: "Form",
			Fields: domain.FormFields{
				domain.FormField{
					Key:   "key",
					Label: "label",
					Type:  "type",
				},
			},
			Submissions: domain.FormSubmissions{
				domain.FormSubmission{
					FormId:    1,
					IPAddress: "127.0.0.1",
					UserAgent: "chrome",
				},
			},
		},
		{
			Id:   2,
			Name: "Form1",
			Fields: domain.FormFields{
				domain.FormField{
					Key:   "key",
					Label: "label",
					Type:  "type",
				},
			},
			Submissions: domain.FormSubmissions{
				domain.FormSubmission{
					FormId:    1,
					IPAddress: "127.0.0.1",
					UserAgent: "chrome",
				},
			},
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
