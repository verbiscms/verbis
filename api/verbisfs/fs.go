// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbisfs

import (
	"github.com/ainsleyclark/verbis/admin"
	"github.com/ainsleyclark/verbis/api/common/paths"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/www"
	"io/fs"
	"path/filepath"
)

// FS defines the set of methods used for interacting with
// the Verbis file system.
type FS interface {
	// Open opens the named file for reading. If successful,
	// methods on the returned file can be used for
	// reading; the associated file descriptor has
	// mode O_RDONLY.
	// Returns ErrFileNotFound on failed lookup
	Open(name string) (fs.File, error)
	// ReadFile reads the named file and returns the contents.
	// A successful call returns err == nil, not err == EOF.
	// Because ReadFile reads the whole file, it does not
	// treat an EOF from Read as an error to be reported.
	// Returns ErrFileNotFound on failed lookup.
	ReadFile(name string) ([]byte, error)
	// ReadDir reads the named directory, returning all its
	// directory entries sorted by filename. If an error
	// occurs reading the directory, ReadDir returns
	// the entries it was able to read before the
	// error, along with the error.
	// Returns ErrDirNotFound on failed lookup
	ReadDir(name string) ([]fs.DirEntry, error)
}

var (
	// ErrFileNotFound is returned by the Open and ReadFile
	// functions when a file could not be found.
	ErrFileNotFound = errors.New("file does not exist")
	// ErrDirNotFound is returned by the ReadDir function when
	// a directory could not be found.
	ErrDirNotFound = errors.New("directory does not exist")
)

// FileSystem defines the necessary directory structure
// for Verbis. SPA includes all Vue directories and
// Web includes web files served by Verbis such as
// mail templates and error pages.
type FileSystem struct {
	SPA FS
	Web FS
}

const (
	// SpaDistFolder defines the prefix for Vue, so `/dist` is
	// removed from read functions.
	SpaDistFolder = "dist"
)

// New creates a new Verbis FileSystem dependant on the
// development status. If it is in prod, an OS FS
// will be returned, if production the embedFS
// will be returned.
func New(production bool, p paths.Paths) *FileSystem {
	if !production {
		return &FileSystem{
			SPA: &osFS{path: filepath.Join(p.Admin, SpaDistFolder)},
			Web: &osFS{path: p.Web},
		}
	}
	return &FileSystem{
		SPA: &embedFS{fs: admin.SPA, prefix: SpaDistFolder},
		Web: &embedFS{fs: www.Web},
	}
}
