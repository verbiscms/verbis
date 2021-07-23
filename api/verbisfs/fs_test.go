// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbisfs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/common/paths"
	"io"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tt := map[string]struct {
		input bool
		want  interface{}
	}{
		"Production": {
			true,
			&embedFS{},
		},
		"Development": {
			false,
			&osFS{},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := New(test.input, paths.Paths{})
			assert.Equal(t, reflect.TypeOf(test.want).String(), reflect.TypeOf(got.SPA).String())
			assert.Equal(t, reflect.TypeOf(test.want).String(), reflect.TypeOf(got.Web).String())
		})
	}
}

func Open(fs FS, t *testing.T) {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Relative": {
			"file.txt",
			"hello",
		},
		"Absolute": {
			"/file.txt",
			"hello",
		},
		"With Dot": {
			"./file.txt",
			"hello",
		},
		"Error": {
			"wrong",
			ErrFileNotFound.Error(),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f, err := fs.Open(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			buf := &bytes.Buffer{}
			_, err = io.Copy(buf, f)
			assert.NoError(t, err)
			assert.Equal(t, test.want, buf.String())
		})
	}
}

func ReadFile(fs FS, t *testing.T) {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Relative": {
			"file.txt",
			"hello",
		},
		"Absolute": {
			"/file.txt",
			"hello",
		},
		"With Dot": {
			"./file.txt",
			"hello",
		},
		"Error": {
			"wrong",
			ErrFileNotFound.Error(),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := fs.ReadFile(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, string(got))
		})
	}
}

func ReadDir(fs FS, t *testing.T) {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Forward Slash": {
			"/",
			1,
		},
		"Empty": {
			"",
			1,
		},
		"Error": {
			"wrong",
			ErrDirNotFound.Error(),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := fs.ReadDir(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, len(got))
		})
	}
}
