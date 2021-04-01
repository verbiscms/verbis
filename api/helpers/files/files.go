// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"os"
	"path/filepath"
)

// Exists
//
// Checks if a file exists using os.Stat.
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirectoryExists
//
// Checks if directory exists using os.Stat.
func DirectoryExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// RemoveFileExtension
//
// Remove the file extension from a file.
func RemoveFileExtension(file string) string {
	return file[0 : len(file)-len(filepath.Ext(file))]
}
