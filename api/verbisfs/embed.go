// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbisfs

import (
	"embed"
	"github.com/pkg/errors"
	"io/fs"
	"path/filepath"
)

// embedFS defines the file system for the embedded files
// in development. The prefix is joined on each
// resulting call.
type embedFS struct {
	fs     embed.FS
	prefix string
}

// Open Uses the embed.FS Open() function for
// retrieving a fs.File
// Returns ErrFileNotFound if the lookup failed.
func (s *embedFS) Open(name string) (fs.File, error) {
	path := filepath.Join(s.prefix, name)
	file, err := s.fs.Open(path)
	if err != nil {
		return nil, errors.Wrap(ErrFileNotFound, path)
	}
	return file, nil
}

// ReadFile Uses the embed.FS ReadFile() function
// for retrieving byte content of a file.
// Returns ErrFileNotFound if the lookup failed.
func (s *embedFS) ReadFile(name string) ([]byte, error) {
	path := filepath.Join(s.prefix, name)
	bytes, err := s.fs.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(ErrFileNotFound, name)
	}
	return bytes, nil
}

// ReadDir Uses the embed.FS ReadDir() function
// for retrieving DirEntries.
// Returns ErrFileNotFound if the lookup failed.
func (s *embedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	path := filepath.Join(s.prefix, name)
	entries, err := s.fs.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(ErrDirNotFound, path)
	}
	return entries, nil
}
