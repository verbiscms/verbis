// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/assert"
	"image"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
)

func Setup(path string) (*multipart.FileHeader, error) {
	base, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	m := test.MediaSuite{}
	part, err := m.ToMultiPartE(filepath.Join(filepath.Dir(base), "testdata", path))
	if err != nil {
		return nil, err
	}
	return part, nil
}

type mockFileSeekErr struct{}

func (m *mockFileSeekErr) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (m *mockFileSeekErr) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, nil
}

func (m *mockFileSeekErr) Seek(offset int64, whence int) (int64, error) {
	return 0, fmt.Errorf("error")
}

func (m *mockFileSeekErr) Close() error {
	return nil
}

func UtilTestDecode(fn func(file multipart.File) Imager, path string, t *testing.T) {
	tt := map[string]struct {
		input func() (multipart.File, func() error)
		want  interface{}
	}{
		"Success": {
			func() (multipart.File, func() error) {
				m, err := Setup(path)
				assert.NoError(t, err)
				file, _ := m.Open() // Ignore on purpose
				return file, file.Close
			},
			nil,
		},
		"Seek Error": {
			func() (multipart.File, func() error) {
				return &mockFileSeekErr{}, func() error { return nil }
			},
			ErrSeekFileMessage,
		},
		"Decode Error": {
			func() (multipart.File, func() error) {
				m, err := Setup("test.txt")
				assert.NoError(t, err)
				file, _ := m.Open() // Ignore on purpose
				return file, file.Close
			},
			ErrDecodeMessage,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			file, teardown := test.input()
			defer teardown() // nolint

			imager := fn(file)
			enc, err := imager.Decode()
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}

			assert.NotNil(t, enc)
			assert.Nil(t, err)
		})
	}
}

func UtilTestEncode(img Imager, t *testing.T) {
	suite := test.MediaSuite{}

	tt := map[string]struct {
		input image.Image
		error bool
	}{
		"Success": {
			suite.Image(),
			false,
		},
		"Error": {
			&image.RGBA{
				Pix:    nil,
				Stride: 0,
				Rect: image.Rectangle{
					Min: image.Point{X: 0, Y: 0},
					Max: image.Point{X: 999999999999999, Y: 999999999999999},
				},
			},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			decode, err := img.Encode(test.input, 0)
			if test.error {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.NotNil(t, decode)
		})
	}
}
