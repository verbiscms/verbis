// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/media/sizes"
	"github.com/ainsleyclark/verbis/api/store/config"
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
func (t *MediaTestSuite) Setup(mf func(m sqlmock.Sqlmock), mfm func(m *mocks.Repository)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}

	m := &mocks.Repository{}
	if mfm != nil {
		mfm(m)
	}

	s := New(&config.Config{
		Driver: t.Driver,
	})

	s.sizes = m

	return s
}

const (
	// The default media item ID used for testing.
	mediaID = "1"
)

//
var (
	// The default media item used for testing.
	mediaItem = domain.Media{
		Id: 1,
		File: domain.File{
			Name: "gopher.png",
		},
		Sizes: mediaItemSizes,
	}
	// The default media item with URI used
	// for testing.
	mediaItemURL = domain.Media{
		Id: 1,
		File: domain.File{
			URL: "/2020/01/gopher.png",
		},
	}
	// The default media sizes used for testing.
	mediaItemSizes = domain.MediaSizes{
		"hd": domain.MediaSize{
			SizeKey:  "hd",
			SizeName: "gopher-1920x1080.jpg",
		},
		"thumbnail": domain.MediaSize{
			SizeKey:  "thumbnail",
			SizeName: "Thumbnail Size",
		},
	}
	// The default media items used for testing.
	mediaItems = domain.MediaItems{
		{
			Id:    1,
			Title: "title",
		},
		{
			Id:    1,
			Title: "title",
		},
	}
)
