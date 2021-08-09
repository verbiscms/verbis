// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sizes

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// SizesTestSuite defines the helper used for media
// size testing.
type SizesTestSuite struct {
	test.DBSuite
}

// TestSizes asserts testing has begun.
func TestSizes(t *testing.T) {
	suite.Run(t, &SizesTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup is a a helper to obtain a mock media sizes
// database for testing.
func (t *SizesTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&config.Config{
		Driver: t.Driver,
	})
}

const (
	// The default media item ID used for testing.
	mediaID = "1"
	// SelectStatement the media select statement.
	SelectStatement = "SELECT media_sizes.*, file.id `file.id`, file.url `file.url`, file.name `file.name`, file.bucket_id `file.bucket_id`, file.mime `file.mime`, file.source_type `file.source_type`, file.provider `file.provider`, file.region `file.region`, file.bucket `file.bucket`, file.file_size `file.file_size`, file.private `file.private` FROM `media_sizes` LEFT JOIN `files` AS `file` ON `media_sizes`.`file_id` = `file`.`id` "
)

var (
	// The default media sizes used for testing.
	sizes = domain.MediaSizes{
		"hd": domain.MediaSize{
			ID:       1,
			SizeKey:  "hd",
			SizeName: "gopher-1920x1080.jpg",
		},
	}
)
