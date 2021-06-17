// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbisfs

import (
	"github.com/gookit/color"
	"io/fs"
	"os"
	"path/filepath"
)

type osFS struct {
	path string
}

// Open opens the named file for reading. If successful,
// methods on the returned file can be used for
// reading; the associated file descriptor has
// mode O_RDONLY.
// Returns *PathError if there was one.
func (s *osFS) Open(name string) (fs.File, error) {
	color.Red.Println(name)
	return os.Open(filepath.Join(s.path, name))
}

// ReadFile reads the named file and returns the contents.
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not
// treat an EOF from Read as an error to be reported.
func (s *osFS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.path, name))
}

// ReadDir reads the named directory, returning all its
// directory entries sorted by filename. If an error
// occurs reading the directory, ReadDir returns
// the entries it was able to read before the
// error, along with the error.
func (s *osFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(filepath.Join(s.path, name))
}
