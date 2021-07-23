// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// MetaTestSuite defines the helper used for role
// testing.
type MetaTestSuite struct {
	test.DBSuite
}

// TestMeta
//
// Assert testing has begun.
func TestMeta(t *testing.T) {
	suite.Run(t, &MetaTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock post meta database
// for testing.
func (t *MetaTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
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
	postID = "2"
)

var (
	// The default meta used for testing.
	meta = domain.PostOptions{
		Id:     1,
		PostId: 2,
		Meta: &domain.PostMeta{
			Title:       "title",
			Description: "description",
		},
		Seo: &domain.PostSeo{
			Canonical: "canonical",
		},
		EditLock: "",
	}
)
