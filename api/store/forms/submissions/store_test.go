// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package submissions

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// SubmissionTestSuite defines the helper used for form
// submission testing.
type SubmissionTestSuite struct {
	test.DBSuite
}

// TestSubmissions
//
// Assert testing has begun.
func TestSubmissions(t *testing.T) {
	suite.Run(t, &SubmissionTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock form submissions
// database for testing.
func (t *SubmissionTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
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
	// The default formSubmission used for testing.
	formSubmission = domain.FormSubmission{
		FormId:    1,
		IPAddress: "127.0.0.1",
		UserAgent: "chrome",
	}
	// The default formSubmission used for testing.
	formSubmissions = domain.FormSubmissions{
		{
			FormId:    1,
			IPAddress: "127.0.0.1",
			UserAgent: "chrome",
		},
	}
)
