// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	tt := map[string]struct {
		version string
		pre     string
		want    string
	}{
		"Empty Prerelease": {
			"v0.0.1",
			"",
			"v0.0.1",
		},
		"Prerelease": {
			"v0.0.1",
			"pre",
			"v0.0.1-pre",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			origVersion := Version
			origPrerelease := Prerelease
			defer func() {
				Version = origVersion
				Prerelease = origPrerelease
			}()
			Version = test.version
			Prerelease = test.pre
			got := String()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMust(t *testing.T) {
	tt := map[string]struct {
		version string
		panics  bool
		want    interface{}
	}{
		"Success": {
			"v0.0.1",
			false,
			"v0.0.1",
		},
		"Panics": {
			"wrong",
			true,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			if test.panics {
				assert.Panics(t, func() {
					Must(test.version)
				})
				return
			}
			got := Must(test.version)
			assert.Equal(t, test.want, got.Original())
		})
	}
}
