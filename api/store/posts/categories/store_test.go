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

// PostCategoriesTestSuite defines the helper used for
// category testing.
type PostCategoriesTestSuite struct {
	test.DBSuite
}

// TestPostCategories
//
// Assert testing has begun.
func TestPostCategories(t *testing.T) {
	suite.Run(t, &PostCategoriesTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock categories database
// for testing.
func (t *PostCategoriesTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
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
	// The default category ID used for testing.
	categoryID = "2"
)

var (
	// The default post used for testing.
	post = domain.PostDatum{
		Post: domain.Post{
			ID: 1,
		},
		Category: &domain.Category{
			ID: 2,
		},
	}
)
