// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/stretchr/testify/suite"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

// MediaSuite defines the helper used for media
// library testing.
type MediaSuite struct {
	suite.Suite
	ApiPath   string
	MediaPath string
}

// MediaTestPath is the default media test path.
const MediaTestPath = "/test/testdata/media"

func NewMediaSuite() MediaSuite {
	return MediaSuite{}
}

// SetupSuite reassigns API path for testing.
func (t *MediaSuite) SetupSuite() {
	logger.SetOutput(ioutil.Discard)
	wd, err := os.Getwd()
	t.NoError(err)
	t.ApiPath = filepath.Join(filepath.Dir(wd), "../")
	t.MediaPath = t.ApiPath + MediaTestPath
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

// File converts a file path into a *multipart.FileHeader.
func (t *MediaSuite) File(path string) *multipart.FileHeader {
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

// File converts a file path into a *multipart.FileHeader.
func ToMultiPart(path string) (*multipart.FileHeader, error) {
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
	mt, err := mr.ReadForm(99999)
	if err != nil {
		return nil, err
	}

	ft := mt.File["file"][0]

	return ft, nil
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

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
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


// Image returns a new image.Image for testing.
func Image() image.Image {
	var (
		width  = 5
		height = 5
	)

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
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
