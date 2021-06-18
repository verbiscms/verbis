// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSys_Restart(t *testing.T) {
	tt := map[string]struct {
		sys  func(argv0 string, argv []string, envv []string) (err error)
		want interface{}
	}{
		"Success": {
			func(argv0 string, argv []string, envv []string) (err error) {
				return nil
			},
			nil,
		},
		"Error": {
			func(argv0 string, argv []string, envv []string) (err error) {
				return fmt.Errorf("error")
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			if test.sys == nil {
				t.Fatal("sys function cannot be nil")
				return
			}

			origSys := sysex
			defer func() {
				sysex = origSys
			}()
			sysex = test.sys

			s := Sys{ExecutablePath: "exec"}
			err := s.Restart()
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, err)
		})
	}
}

func TestSys_Restart_Error(t *testing.T) {
	s := &Sys{ExecutablePath: "wrong"}
	err := s.Restart()
	assert.Contains(t, err.Error(), "no such file or directory")
}
