// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"os"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

func TestNamespace_Env(t *testing.T) {
	tt := map[string]struct {
		env   func() error
		input string
		want  string
	}{
		"Valid": {
			func() error { return os.Setenv("verbis", "cms") },
			"verbis",
			"cms",
		},
		"Valid 2": {
			func() error { return os.Setenv("foo", "bar") },
			"foo",
			"bar",
		},
		"Not found": {
			func() error { return nil },
			"",
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			err := test.env()
			got := ns.Env(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_ExpandEnv(t *testing.T) {
	tt := map[string]struct {
		env   func() error
		input string
		want  string
	}{
		"Valid": {
			func() error { return os.Setenv("path", "verbis") },
			"$path is my name",
			"verbis is my name",
		},
		"Valid 2": {
			func() error { return os.Setenv("foo", "bar") },
			"hello $foo",
			"hello bar",
		},
		"Not found": {
			func() error { return nil },
			"hello $test verbis",
			"hello  verbis",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			err := test.env()
			got := ns.ExpandEnv(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
