// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/logger"
	sm "github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestUpdateRegistry_Sort(t *testing.T) {
	tt := map[string]struct {
		input UpdateRegistry
		want  UpdateRegistry
	}{
		"Success": {
			UpdateRegistry{&Update{Version: "v3.0.0"}, &Update{Version: "v1.0.0"}, &Update{Version: "v2.0.0"}},
			UpdateRegistry{&Update{Version: "v1.0.0"}, &Update{Version: "v2.0.0"}, &Update{Version: "v3.0.0"}},
		},
		"Already Sorted": {
			UpdateRegistry{&Update{Version: "v1.0.0"}, &Update{Version: "v2.0.0"}, &Update{Version: "v3.0.0"}},
			UpdateRegistry{&Update{Version: "v1.0.0"}, &Update{Version: "v2.0.0"}, &Update{Version: "v3.0.0"}},
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

func TestUpdate_ToSemVer(t *testing.T) {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		input  Update
		panics bool
		want   interface{}
	}{
		"Success": {
			Update{Version: "v0.0.1"},
			false,
			sm.Must(sm.NewVersion("v0.0.1")),
		},
		"Error": {
			Update{Version: "wrong"},
			true,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			if test.panics {
				assert.Panics(t, func() { test.input.ToSemVer() })
				return
			}
			assert.Equal(t, test.want, test.input.ToSemVer())
		})
	}
}

func TestUpdate_HasCallBack(t *testing.T) {
	tt := map[string]struct {
		input Update
		want  bool
	}{
		"Has CallBack": {
			Update{CallBackUp: func() error {
				return nil
			}, CallBackDown: func() error {
				return nil
			}},
			true,
		},
		"No Callback": {
			Update{},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.HasCallBack()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestUpdateRegistry_AddUpdate(t *testing.T) {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		input  Update
		panics bool
	}{
		"Success": {
			Update{Version: "v0.0.1", Stage: Minor, MigrationPath: "v0.0.1.sql"},
			false,
		},
		"No Version": {
			Update{Version: ""},
			true,
		},
		"No Stage": {
			Update{Version: "v0.0.1"},
			true,
		},
		"No Migration Path": {
			Update{Version: "v0.0.1", Stage: Minor},
			true,
		},
		"Bad Version": {
			Update{Version: "v1.3.3.3", MigrationPath: "test", Stage: Minor},
			true,
		},
		"No CallBackUp": {
			Update{Version: "v0.0.1", MigrationPath: "test", Stage: Minor, CallBackDown: func() error {
				return nil
			}},
			true,
		},
		"No CallBackDown": {
			Update{Version: "v0.0.1", MigrationPath: "test", Stage: Minor, CallBackUp: func() error {
				return nil
			}},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			u := make(UpdateRegistry, 0)
			if test.panics {
				assert.Panics(t, func() { u.AddUpdate(&test.input) })
				return
			}
			u.AddUpdate(&test.input)
			assert.Equal(t, test.input, *u[0])
		})
	}
}
