// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package update

import (
	"os"
	"path/filepath"
)

func (u *Update) backup() error {
	for _, v := range Folders {
		path := u.Paths.Base + v
		err := os.Rename(path, path+".old")
		if err != nil {
			return err
		}
	}
	return nil
}

// cleanup removes the temporary folder and the
// `verbis.old` file from the current working
// directory.
func (u *Update) cleanup() error {
	for _, v := range Folders {
		err := os.RemoveAll(u.Paths.Base + v + ".old")
		if err != nil {
			return err
		}
	}

	execBackup := u.Paths.Base + string(os.PathSeparator) + filepath.Base(u.Paths.Exec) + ".old"
	err := os.Remove(execBackup)
	if err != nil {
		return err
	}

	return nil
}
