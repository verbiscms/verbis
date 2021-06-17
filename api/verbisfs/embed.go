// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbisfs

import (
	"embed"
	"github.com/gookit/color"
	"io/fs"
)

type embedFS struct {
	fs     embed.FS
	prefix string
}

func (s *embedFS) Open(name string) (fs.File, error) {
	color.Red.Println(name)
	return s.fs.Open(s.prefix + name)
}

func (s *embedFS) ReadFile(name string) ([]byte, error) {
	return s.fs.ReadFile(s.prefix + name)
}

func (s *embedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return s.fs.ReadDir(s.prefix + name)
}
