// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	location "github.com/ainsleyclark/verbis/api/mocks/services/fields/location"
	categories "github.com/ainsleyclark/verbis/api/mocks/store/categories"
	users "github.com/ainsleyclark/verbis/api/mocks/store/users"
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
func (t *FieldTestSuite) Setup(mf func(l *location.Finder, u *users.Repository, c *categories.Repository)) *Fields {
	var (
		l = &location.Finder{}
		u = &users.Repository{}
		c = &categories.Repository{}
	)

	if mf != nil {
		mf(l, u, c)
	}

	f := New(&deps.Deps{
		Options: &domain.Options{},
		Store: &store.Repository{
			User:       u,
			Categories: c,
		},
	})
	f.finder = l

	return f
}

var (
	// The default fields used for testing.
	fields = domain.Fields{
		domain.Field{
			Label: "label",
			Name:  "name",
			Type:  "type",
		},
	}
	// The default field groups used for testing.
	fieldGroups = domain.FieldGroups{
		domain.FieldGroup{
			Title:  "group",
			Fields: fields,
		},
	}
	// The default category used for testing.
	category = domain.Category{
		Id:       123,
		Slug:     "/cat",
		Name:     "Category",
		Resource: "test",
	}
	// The default user used for testing.
	user = domain.User{
		UserPart: domain.UserPart{
			Id:        123,
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
			Role: domain.Role{
				Id: 1,
			},
		},
	}
)
