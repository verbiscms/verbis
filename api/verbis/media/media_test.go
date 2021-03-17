// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
)

// MediaTestSuite defines the helper used for media
// library testing.
type MediaTestSuite struct {
	test.HandlerSuite
}

// TestAuth
//
// Assert testing has begun.
func TestAuth(t *testing.T) {
	suite.Run(t, &MediaTestSuite{})
}

// Setup
//
// A helper to obtain a mock media client for
// testing.
func (t *MediaTestSuite) Setup(cfg domain.ThemeConfig, opts domain.Options) *client {
	return &client{
		Options: &opts,
		Config:  &cfg,
		paths:   paths.Get(),
		Exists:  nil,
	}
}

// File
//
// Converts a path into a *multipart.FileHeader
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
