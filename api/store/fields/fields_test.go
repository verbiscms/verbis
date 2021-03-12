// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
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
func (t *FieldsTestSuite) Setup() *Store {
	t.Reset()
	return New(&store.Config{
		Driver: t.Driver,
		Options: &domain.Options{
			CacheServerFields: false,
		},
	})
}

var (
	// The default field groups used for testing.
	groups = domain.FieldGroups{
		domain.FieldGroup{
			Title: "groups",
			Fields: domain.Fields{
				domain.Field{
					Label: "label",
					Name:  "name",
					Type:  "text",
				},
			},
			Locations: [][]domain.FieldLocation{
				{
					domain.FieldLocation{
						Param:    "resource",
						Operator: "=",
						Value:    "news",
					},
				},
			},
		},
	}
)
