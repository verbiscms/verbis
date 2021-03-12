// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// MediaTestSuite defines the helper used for
// media testing.
type MediaTestSuite struct {
	test.DBSuite
}

// TestMedia
//
// Assert testing has begun.
func TestMedia(t *testing.T) {
	suite.Run(t, &MediaTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock media database
// for testing.
func (t *MediaTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&store.Config{
		Driver: t.Driver,
	})
}

const (
	// The default media item ID used for testing.
	mediaID = "1"
	// The select statement.
	SelectStatement = "SELECT id, uuid, url, file_path, file_size, file_name, sizes, type, user_id, updated_at, created_at, CASE WHEN title IS NULL THEN '' ELSE title END AS 'title', CASE WHEN alt IS NULL THEN '' ELSE alt END AS 'alt', CASE WHEN description IS NULL THEN '' ELSE description END AS 'description' FROM `media`"
)

var (
	// The default media item used for testing.
	mediaItem = domain.Media{
		Id: 1,
	}
	// The default media items used for testing.
	mediaItems = domain.MediaItems{
		{
			Id:    1,
			Url:   "/uploads/1",
			Title: "title",
		},
		{
			Id:    1,
			Url:   "/uploads/1",
			Title: "title",
		},
	}
)
