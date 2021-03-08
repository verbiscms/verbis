// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// CategoriesTestSuite defines the helper used for role
// testing.
type CategoriesTestSuite struct {
	test.DBSuite
}

// TestCategories
//
// Assert testing has begun.
func TestCategories(t *testing.T) {
	suite.Run(t, &CategoriesTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock roles database
// for testing.
func (t *CategoriesTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	if mf != nil {
		mf(t.Mock)
	}
	return New(t.DB)
}

const (
	// The default category ID used for testing.
	categoryID = "1"
)

var (
	// The default category used for testing.
	category = domain.Category{
		Id:   1,
		Slug: "/cat",
		Name: "Category",
	}
	// The default categories used for testing.
	categories = domain.Categories{
		{
			Id:   1,
			Slug: "/cat",
			Name: "Category",
		},
		{
			Id:   2,
			Slug: "/cat1",
			Name: "Category1",
		},
	}
	// The default params used for testing.
	defaultParams = params.Params{
		Page:           api.DefaultParams.Page,
		Limit:          15,
		OrderBy:        api.DefaultParams.OrderBy,
		OrderDirection: api.DefaultParams.OrderDirection,
		Filters:        nil,
	}
)
