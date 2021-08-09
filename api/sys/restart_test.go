// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
)

func (t *SysTestSuite) TestSys_Restart() {
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
		t.Run(name, func() {
			if test.sys == nil {
				t.Fail("sys function cannot be nil")
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
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}

func (t *SysTestSuite) TestSys_Restart_Error() {
	s := &Sys{ExecutablePath: "wrong"}
	err := s.Restart()
	t.Contains(err.Error(), "no such file or directory")
}
