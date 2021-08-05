// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// CategoriesTestSuite defines the helper used for
// category testing.
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
// A helper to obtain a mock categories database
// for testing.
func (t *CategoriesTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&config.Config{
		Driver: t.Driver,
	})
}

const (
	// The default category ID used for testing.
	categoryID = "1"
)

var (
	// The default category used for testing.
	category = domain.Category{
		ID:   1,
		Slug: "/cat",
		Name: "Category",
	}
	// The default categories used for testing.
	categories = domain.Categories{
		{
			ID:   1,
			Slug: "/cat",
			Name: "Category",
		},
		{
			ID:   2,
			Slug: "/cat1",
			Name: "Category1",
		},
	}
)
