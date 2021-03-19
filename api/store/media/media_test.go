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
)

var (
	// The default media item used for testing.
	mediaItem = domain.Media{
		Id:       1,
		FileName: "gopher.png",
	}
	// The default media item with URL used
	// for testing.
	mediaItemURL = domain.Media{
		Id:  1,
		Url: "/2020/01/gopher.png",
	}
	// The default media item with sizes used
	// for testing.
	mediaItemSizes = domain.Media{
		Id:  1,
		Url: "/2020/01/gopher.png",
		Sizes: domain.MediaSizes{
			"test": domain.MediaSize{
				Url: "/2020/01/gopher-100x100.png",
			},
		},
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
