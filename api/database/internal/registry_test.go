// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/version"
)

func (t *InternalTestSuite) TestAddMigration() {
	tt := map[string]struct {
		input Migration
		twice bool
		want  interface{}
	}{
		"Success": {
			Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1")},
			false,
			nil,
		},
		"No version": {
			Migration{Version: ""},
			false,
			"no version provided for update",
		},
		"Nil SemVer": {
			Migration{Version: "v0.0.1", Stage: version.Major, SemVer: nil},
			false,
			nil,
		},
		"Bad Version": {
			Migration{Version: "wrong", Stage: version.Major, SemVer: nil},
			false,
			"malformed version",
		},
		"No Stage": {
			Migration{Version: "v0.0.1"},
			false,
			"no stage set",
		},
		"No CallBackUp": {
			Migration{Version: "v0.0.1", Stage: version.Major, CallBackDown: func() error {
				return nil
			}},
			false,
			ErrCallBackMismatch.Error(),
		},
		"No CallBackDown": {
			Migration{Version: "v0.0.1", Stage: version.Major, CallBackUp: func() error {
				return nil
			}},
			false,
			ErrCallBackMismatch.Error(),
		},
		"Duplicate": {
			Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1")},
			true,
			"duplicate version",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			defer func() {
				migrations = make(MigrationRegistry, 0)
			}()
			err := AddMigration(&test.input)
			if test.twice {
				err = AddMigration(&test.input)
			}
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.input, *migrations[0])
		})
	}
}

func (t *InternalTestSuite) TestMigration_HasCallBack() {
	tt := map[string]struct {
		input Migration
		want  bool
	}{
		"Has CallBack": {
			Migration{CallBackUp: func() error {
				return nil
			}, CallBackDown: func() error {
				return nil
			}},
			true,
		},
		"No Callback": {
			Migration{},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := test.input.hasCallBack()
			t.Equal(test.want, got)
		})
	}
}

func (t *InternalTestSuite) TestMigrationRegistry_Sort() {
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
		t.Run(name, func() {
			test.input.Sort()
			for i, v := range test.input {
				t.Equal(test.want[i].Version, v.Version)
			}
		})
	}
}
