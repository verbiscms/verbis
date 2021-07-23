// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"github.com/verbiscms/verbis/api/errors"
	"os"
	"path/filepath"
	"syscall"
)

var (
	// exec represents stdlib os.Executable func.
	exec = os.Executable
	// bin defines the original os arguments.
	bin = os.Args[0]
	// sysex represents stdlib syscall.Exec func.
	sysex = syscall.Exec
)

// Restart stops the currently running process and
// restarts the executable with the original os
// arguments.
func (s *Sys) Restart() error {
	const op = "System.Restart"
	err := sysex(s.ExecutablePath, append([]string{s.ExecutablePath}, os.Args[1:]...), os.Environ())
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error restating system", Operation: op, Err: err}
	}
	return nil
}

// execPath returns the location to the current
// executable.
func execPath() (string, error) {
	const op = "System.ExecPath"

	if !filepath.IsAbs(bin) {
		var err error
		bin, err = exec()
		if err != nil {
			err = fmt.Errorf("cannot get path to binary %q (launch with absolute path): %w", os.Args[0], err)
			return "", &errors.Error{Code: errors.INTERNAL, Message: "Error getting executable name", Operation: op, Err: err}
		}
	}

	return bin, nil
}
