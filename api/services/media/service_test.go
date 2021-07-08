// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	webp "github.com/ainsleyclark/verbis/api/mocks/services/webp"
	storage "github.com/ainsleyclark/verbis/api/mocks/storage"
	repo "github.com/ainsleyclark/verbis/api/mocks/store/media"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// MediaServiceTestSuite defines the helper used for media
// library testing.
type MediaServiceTestSuite struct {
	test.MediaSuite
}

// TestMediaService asserts testing has begun.
func TestMediaService(t *testing.T) {
	suite.Run(t, &MediaServiceTestSuite{
		MediaSuite: test.NewMediaSuite(),
	})
}

const (
	// MediaId is the default ID use for testing.
	MediaId = 1
)

// Setup is a helper to obtain a mock media Service
// for testing.
func (t *MediaServiceTestSuite) Setup(cfg *domain.ThemeConfig, opts *domain.Options, mock func(r *repo.Repository, s *storage.Bucket)) *Service {
	m := &webp.Execer{}
	r := &repo.Repository{}
	s := &storage.Bucket{}

	if mock != nil {
		mock(r, s)
	}

	if cfg == nil {
		cfg = &domain.ThemeConfig{}
	}

	if opts == nil {
		opts = &domain.Options{}
	}

	//m.On("Convert", mock.Anything, mock.Anything).Once()
	//m.On("Convert", mock.Anything, mock.Anything).Once()

	return &Service{
		options: opts,
		config:  cfg,
		paths: paths.Paths{
			API:     t.ApiPath,
			Uploads: t.ApiPath + test.MediaTestPath,
		},
		webp:    m,
		repo:    r,
		storage: s,
	}
}
