// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	mocks "github.com/ainsleyclark/verbis/api/mocks/services/webp"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
)

// MediaTestSuite defines the helper used for media
// library testing.
type MediaTestSuite struct {
	test.HandlerSuite
	apiPath   string
	mediaPath string
}

// TestAuth
//
// Assert testing has begun.
func TestAuth(t *testing.T) {
	suite.Run(t, &MediaTestSuite{})
}

// The default test path.
const TestPath = "/test/testdata/media"

var (
	// The exists function used for testing.
	exists = func(fileName string) bool { return false }
)

// Setup
//
// A helper to obtain a mock media Service for
// testing.
func (t *MediaTestSuite) Setup(cfg domain.ThemeConfig, opts domain.Options) *Service {
	m := &mocks.Execer{}
	m.On("Convert", mock.Anything, mock.Anything).Once()
	m.On("Convert", mock.Anything, mock.Anything).Once()
	return &Service{
		options: &opts,
		config:  &cfg,
		paths: paths.Paths{
			API:     t.apiPath,
			Uploads: t.apiPath + TestPath,
		},
		exists: nil,
		webp:   m,
	}
}

// SetupSuite
//
// Reassign API path for testing.
func (t *MediaTestSuite) SetupSuite() {
	logger.SetOutput(ioutil.Discard)
	wd, err := os.Getwd()
	t.NoError(err)
	t.apiPath = filepath.Join(filepath.Dir(wd), "../")
	t.mediaPath = t.apiPath + TestPath
}

// DummyFile
//
// Creates a dummy file for testing with the
// given path.
func (t *MediaTestSuite) DummyFile(path string) func() {
	file, err := os.Create(path)
	if err != nil {
		t.Fail("Error creating file with the path: "+path, err)
	}
	return func() {
		err := file.Close()
		if err != nil {
			t.Fail("Error closing file", err)
		}
	}
}

// File
//
// Converts a path into a *multipart.FileHeader.
func (t *MediaTestSuite) File(path string) *multipart.FileHeader {
	file, err := os.Open(path)
	t.NoError(err)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	t.NoError(err)
	_, err = io.Copy(part, file)
	t.NoError(err)

	err = writer.Close()
	t.NoError(err)

	mr := multipart.NewReader(body, writer.Boundary())
	mt, err := mr.ReadForm(99999)
	t.NoError(err)
	ft := mt.File["file"][0]

	return ft
}
