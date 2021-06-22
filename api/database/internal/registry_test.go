// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddMigration(t *testing.T) {
	tt := map[string]struct {
		input Migration
		want  interface{}
	}{
		"Success": {
			Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1")},
			nil,
		},
		"No version": {
			Migration{Version: ""},
			"no version provided for update",
		},
		"Nil SemVer": {
			Migration{Version: "v0.0.1", Stage: version.Major, SemVer: nil},
			"missing sem ver",
		},
		"No Stage": {
			Migration{Version: "v0.0.1"},
			"no stage set",
		},
		"No CallBackUp": {
			Migration{Version: "v0.0.1", Stage: version.Major, CallBackDown: func() error {
				return nil
			}},
			ErrCallBackMismatch.Error(),
		},
		"No CallBackDown": {
			Migration{Version: "v0.0.1", Stage: version.Major, CallBackUp: func() error {
				return nil
			}},
			ErrCallBackMismatch.Error(),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			defer func() {
				migrations = make(MigrationRegistry, 0)
			}()
			err := AddMigration(&test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.input, *migrations[0])
		})
	}
}

func TestMigrationRegistry_Sort(t *testing.T) {
	var (
		v1 = version.Must("v1.0.0")
		v2 = version.Must("v2.0.0")
		v3 = version.Must("v3.0.0")
	)

	tt := map[string]struct {
		input MigrationRegistry
		want  MigrationRegistry
	}{
		"Success": {
			MigrationRegistry{
				&Migration{Version: "v3.0.0", SemVer: v3},
				&Migration{Version: "v1.0.0", SemVer: v1},
				&Migration{Version: "v2.0.0", SemVer: v2},
			},
			MigrationRegistry{
				&Migration{Version: "v1.0.0", SemVer: v1},
				&Migration{Version: "v2.0.0", SemVer: v2},
				&Migration{Version: "v3.0.0", SemVer: v3},
			},
		},
		"Already Sorted": {
			MigrationRegistry{
				&Migration{Version: "v1.0.0", SemVer: v1},
				&Migration{Version: "v2.0.0", SemVer: v2},
				&Migration{Version: "v3.0.0", SemVer: v3},
			},
			MigrationRegistry{
				&Migration{Version: "v1.0.0", SemVer: v1},
				&Migration{Version: "v2.0.0", SemVer: v2},
				&Migration{Version: "v3.0.0", SemVer: v3},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			test.input.Sort()
			for i, v := range test.input {
				assert.Equal(t, test.want[i].Version, v.Version)
			}
		})
	}
}
