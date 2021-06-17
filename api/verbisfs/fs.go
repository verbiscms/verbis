// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbisfs

import (
	"github.com/ainsleyclark/verbis/admin"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/www"
	"io/fs"
	"path/filepath"
)

// FS defines the set of methods used for interacting with
// the Verbis file system.
type FS interface {
	Open(name string) (fs.File, error)
	ReadFile(name string) ([]byte, error)
	ReadDir(name string) ([]fs.DirEntry, error)
}

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
func New(production bool) *FileSystem {
	p := paths.Get()
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
