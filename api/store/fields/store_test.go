// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store/config"
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
	// The post field used for testing.
	field = domain.PostField{
		Id:            1,
		PostId:        1,
		Type:          "text",
		Name:          "name",
		Key:           "key",
		OriginalValue: "val",
	}
	// The post fields used for testing.
	fields = domain.PostFields{
		{
			Id:            1,
			PostId:        1,
			Type:          "text",
			Name:          "name",
			Key:           "key",
			OriginalValue: "val",
		},
		{
			Id:            2,
			PostId:        1,
			Type:          "text",
			Name:          "name",
			Key:           "key",
			OriginalValue: "val",
		},
	}
	// The post fields copy used for testing.
	fieldsCopy = domain.PostFields{
		{
			Id:            1,
			PostId:        1,
			Type:          "text",
			Name:          "test1",
			Key:           "key1",
			OriginalValue: "val",
		},
		{
			Id:            2,
			PostId:        1,
			Type:          "text",
			Name:          "test2",
			Key:           "key2",
			OriginalValue: "val",
		},
	}
)
