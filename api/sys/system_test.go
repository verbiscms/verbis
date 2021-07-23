// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/logger"
	mocks "github.com/verbiscms/verbis/api/mocks/database"
	"io/ioutil"
	"testing"
)

func TestNew(t *testing.T) {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		exec  func() (string, error)
		bin   string
		panic bool
		want  interface{}
	}{
		"Success": {
			func() (s string, err error) {
				return "exec", nil
			},
			"test",
			false,
			"exec",
		},
		"Error": {
			func() (s string, err error) {
				return "", fmt.Errorf("error")
			},
			"test",
			true,
			"cannot get path to binary",
		},
		"Absolute": {
			func() (s string, err error) {
				return "exec", nil
			},
			"/test",
			false,
			"/test",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			if test.exec == nil {
				t.Fatal("exec function cannot be nil")
				return
			}

			origExec := exec
			origBin := bin

			defer func() {
				bin = origBin
				exec = origExec
			}()

			exec = test.exec
			bin = test.bin

			if test.panic {
				assert.Panics(t, func() {
					New(&mocks.Driver{})
				})
				return
			}

			got := New(&mocks.Driver{})
			assert.Equal(t, test.want, got.ExecutablePath)
		})
	}
}
