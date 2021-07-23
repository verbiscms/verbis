// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"image"
	"image/color"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// MediaSuite defines the helper used for media
// library testing.
type MediaSuite struct {
	suite.Suite
}

// MediaTestPath is the default media test path.
const MediaTestPath = "/test/testdata/media"

func NewMediaSuite() MediaSuite {
	return MediaSuite{}
}

// DummyFile creates a dummy file for testing with the
// given path.
func (t *MediaSuite) DummyFile(path string) func() {
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

// ToMultiPartE converts a file path into a
// *multipart.FileHeader.
func (t *MediaSuite) ToMultiPartE(path string) (*multipart.FileHeader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	mr := multipart.NewReader(body, writer.Boundary())
	mt, err := mr.ReadForm(99999) //nolint
	if err != nil {
		return nil, err
	}

	ft := mt.File["file"][0]

	return ft, nil
}

// ToMultiPart returns a multipart.FilHeader with
// test checks.
func (t *MediaSuite) ToMultiPart(path string) *multipart.FileHeader {
	header, err := t.ToMultiPartE(path)
	t.NoError(err)
	return header
}

// Image returns a new image.Image for testing.
func (t *MediaSuite) Image() image.Image {
	var (
		width  = 5
		height = 5
	)

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// Colours are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{R: 100, G: 200, B: 200, A: 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
			}
		}
	}

	return img
}
