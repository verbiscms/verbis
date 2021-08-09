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

// FieldsTestSuite defines the helper used for
// field testing.
type FieldsTestSuite struct {
	test.DBSuite
}

// FieldsTestSuite
//
// Assert testing has begun.
func TestFields(t *testing.T) {
	suite.Run(t, &FieldsTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock fields database
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
	// The default post ID used for testing.
	postID = "1"
)

var (
	// The post field used for testing.
	field = domain.PostField{
		PostID:        1,
		Type:          "text",
		Name:          "name",
		Key:           "key",
		OriginalValue: "val",
	}
	// The post fields used for testing.
	fields = domain.PostFields{
		{
			PostID:        1,
			Type:          "text",
			Name:          "name",
			Key:           "key",
			OriginalValue: "val",
		},
		{
			PostID:        1,
			Type:          "text",
			Name:          "name",
			Key:           "key",
			OriginalValue: "val",
		},
	}
	// The post fields used for testing.
	fieldsSingular = domain.PostFields{
		{
			PostID:        1,
			Type:          "text",
			Name:          "name",
			Key:           "key",
			OriginalValue: "val",
		},
	}
)
