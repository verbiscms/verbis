// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resizer

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/mocks/services/media/image"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	stdimage "image"
	"testing"
)

func TestResize_Resize(t *testing.T) {
	tt := map[string]struct {
		mock func(m *mocks.Imager)
		size domain.MediaSize
		want interface{}
	}{
		"Success": {
			func(m *mocks.Imager) {
				m.On("Decode").Return(&stdimage.NRGBA{}, nil)
				m.On("Encode", mock.Anything, 0).Return(bytes.NewBuffer([]byte("image")), nil)
			},
			domain.MediaSize{Crop: false},
			"image",
		},
		"Success Crop": {
			func(m *mocks.Imager) {
				m.On("Decode").Return(&stdimage.NRGBA{}, nil)
				m.On("Encode", mock.Anything, 0).Return(bytes.NewBuffer([]byte("image")), nil)
			},
			domain.MediaSize{Crop: true},
			"image",
		},
		"Decode Error": {
			func(m *mocks.Imager) {
				m.On("Decode").Return(nil, fmt.Errorf("error"))
			},
			domain.MediaSize{Crop: false},
			"error",
		},
		"Encode Error": {
			func(m *mocks.Imager) {
				m.On("Decode").Return(&stdimage.NRGBA{}, nil)
				m.On("Encode", mock.Anything, 0).Return(nil, fmt.Errorf("error"))
			},
			domain.MediaSize{Crop: false},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			r := Resize{}

			m := &mocks.Imager{}
			if test.mock != nil {
				test.mock(m)
			}

			got, err := r.Resize(m, test.size)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			buf := &bytes.Buffer{}
			_, err = got.WriteTo(buf)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, buf.String(), test.want)
		})
	}
}

func TestResize_Resize_NilImager(t *testing.T) {
	r := Resize{}
	_, err := r.Resize(nil, domain.MediaSize{})
	if err == nil {
		t.Fatal("expecting ErrNilImager, got nil")
	}
	assert.Contains(t, err.Error(), ErrNilImager.Error())
}
