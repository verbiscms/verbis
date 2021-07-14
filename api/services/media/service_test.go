// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	resizer "github.com/ainsleyclark/verbis/api/mocks/services/media/resizer"
	storage "github.com/ainsleyclark/verbis/api/mocks/storage"
	repo "github.com/ainsleyclark/verbis/api/mocks/store/media"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// MediaServiceTestSuite defines the helper used for testMedia
// library testing.
type MediaServiceTestSuite struct {
	test.MediaSuite
	TestDataPath string
	LogWriter    *bytes.Buffer
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

// BeforeTest setups the LogWriter
func (t *MediaServiceTestSuite) BeforeTest(suiteName, testName string) {
	b := &bytes.Buffer{}
	t.LogWriter = b
	logger.SetOutput(b)
}

// SetupSuite reassigns API path for testing.
func (t *MediaServiceTestSuite) SetupSuite() {
	wd, err := os.Getwd()
	t.NoError(err)
	t.TestDataPath = filepath.Join(wd, "testdata")
}

// Reset the log writer.
func (t *MediaServiceTestSuite) Reset() {
	t.LogWriter.Reset()
}

// File returns a byte value of a path in the testdata
// directory for testing.
func (t *MediaServiceTestSuite) File(path string) []byte {
	b, err := ioutil.ReadFile(filepath.Join(t.TestDataPath, path))
	t.NoError(err)
	return b
}

const (
	TestFileURL     = "/uploads/2020/01/file.jpg"
	TestFileURLWebP = TestFileURL + domain.WebPExtension
)

var (
	SVGFile = domain.File{
		Id:       1,
		Url:      "/uploads/gopher.svg",
		Name:     "gopher.svg",
		BucketId: "/uploads/gopher.svg",
		Mime:     "image/svg+xml",
		Private:  false,
	}
	PNGFile = domain.File{
		Id:       1,
		Url:      "/uploads/gopher.png",
		Name:     "gopher.png",
		BucketId: "/uploads/gopher.png",
		Mime:     "image/png",
		Private:  false,
	}
	JPGFile = domain.File{
		Id:       1,
		Url:      "/uploads/gopher.jpg",
		Name:     "gopher.jpg",
		BucketId: "/uploads/gopher.jpg",
		Mime:     "image/jpeg",
		Private:  false,
	}
	opts = &domain.Options{
		MediaSizes: domain.MediaSizes{"thumbnail": domain.MediaSize{SizeKey: "thumb", SizeName: "thumb", Width: 300, Height: 300, Crop: false}},
	}
	testMedia      = domain.Media{Id: MediaId, File: domain.File{Id: 1, Url: TestFileURL}}
	testMediaSizes = domain.Media{
		Id: MediaId,
		Sizes: domain.MediaSizes{
			"thumnbnail": domain.MediaSize{
				SizeKey:  "key",
				SizeName: "name",
				File:     domain.File{Id: 1, Url: TestFileURL},
			},
		},
		File: domain.File{Id: 1, Url: TestFileURL},
	}
)

// Setup is a helper to obtain a mock testMedia Service
// for testing.
func (t *MediaServiceTestSuite) Setup(cfg *domain.ThemeConfig, opts *domain.Options, mock func(r *repo.Repository, s *storage.Bucket)) *Service {
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

	c := New(opts, s, r)
	c.resizer = &resizer.Resizer{}
	c.config = cfg

	return c
}
