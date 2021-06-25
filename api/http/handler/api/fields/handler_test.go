// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/fields"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// FieldTestSuite defines the helper used for cache
// testing.
type FieldTestSuite struct {
	test.HandlerSuite
}

// TestFields asserts testing has begun.
func TestFields(t *testing.T) {
	suite.Run(t, &FieldTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock fields handler
// for testing.
func (t *FieldTestSuite) Setup(mf func(m *mocks.Repository)) *Fields {
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}
	return &Fields{
		Deps: &deps.Deps{
			Store: &store.Repository{
				Fields: m,
			},
		},
	}
}

var (
	// The default fields used for testing.
	fields = domain.Fields{
		domain.Field{
			Label:        "label",
			Name:         "name",
			Type:         "type",
		},
	}
	// The default field groups used for testing.
	fieldGroups = domain.FieldGroups{
		domain.FieldGroup{
			Title:     "group",
			Fields:    fields,
		},
	}

)
