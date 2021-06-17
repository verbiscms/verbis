// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbisfs

import (
	"github.com/pkg/errors"
	"io/fs"
	"os"
	"path/filepath"
)

// osFS defines the file system for the OS. The path is
// joined on each resulting call.
type osFS struct {
	path string
}

// Open Uses the stdlib Open() function for retrieving a
// fs.File
// Returns ErrFileNotFound if the lookup failed.
func (s *osFS) Open(name string) (fs.File, error) {
	path := filepath.Join(s.path, name)
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(ErrFileNotFound, path)
	}
	return file, nil
}

// ReadFile Uses the stdlib ReadFile() function for
// retrieving byte content of a file.
// Returns ErrFileNotFound if the lookup failed.
func (s *osFS) ReadFile(name string) ([]byte, error) {
	path := filepath.Join(s.path, name)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(ErrFileNotFound, path)
	}
	return bytes, nil
}

// ReadDir Uses the stdlib ReadDir() function for
// retrieving DirEntries.
// Returns ErrFileNotFound if the lookup failed.
func (s *osFS) ReadDir(name string) ([]fs.DirEntry, error) {
	path := filepath.Join(s.path, name)
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(ErrDirNotFound, path)
	}
	return entries, nil
}
