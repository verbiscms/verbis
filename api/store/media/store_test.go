// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/store/media/sizes"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
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
	// SelectStatement the media select statement.
	SelectStatement = "SELECT media.*, file.id `file.id`, file.url `file.url`, file.name `file.name`, file.bucket_id `file.bucket_id`, file.mime `file.mime`, file.source_type `file.source_type`, file.provider `file.provider`, file.region `file.region`, file.bucket `file.bucket`, file.file_size `file.file_size`, file.private `file.private` FROM `media` LEFT JOIN `files` AS `file` ON `media`.`file_id` = `file`.`id` "
)

//
var (
	// The default media item used for testing.
	mediaItem = domain.Media{
		ID: 1,
		File: domain.File{
			Name: "gopher.png",
		},
		Sizes: mediaItemSizes,
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
			ID: 1,
			File: domain.File{
				Name: "gopher.png",
			},
			Sizes: mediaItemSizes,
		},
		{
			ID: 1,
			File: domain.File{
				Name: "gopher-2.png",
			},
			Sizes: mediaItemSizes,
		},
	}
)
