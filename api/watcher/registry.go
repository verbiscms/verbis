// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package watcher

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/fsnotify/fsnotify"
	"os"
)

//
var Registry FuncRegistrar

// Register
type Register struct {
	path string
	fn   NotifyFunc
	dir  bool
}

//
type FuncRegistrar []Register

//
type NotifyFunc func(event *fsnotify.Event)

// AddWatcherFunc
//
//
func AddWatcherFunc(path string, fn NotifyFunc) {
	if Registry == nil {
		Registry = make([]Register, 0)
	}

	isDir, err := checkPath(path)
	if err != nil {
		panic("Path does not exist")
	}

	if fn == nil {
		panic("Empty watcher method mapping for " + path)
	}

	Registry = append(Registry, Register{
		path: path,
		fn:   fn,
		dir:  isDir,
	})
}

// exists returns whether the given file or directory exists
func checkPath(path string) (bool, error) {
	const op = "Watcher.Registry.exists"
	fi, err := os.Stat(path)
	if err != nil {
		return false, &errors.Error{Code: errors.INTERNAL, Message: "No path exists: " + path, Operation: op, Err: err}
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// Is a directory
		return true, nil
	case mode.IsRegular():
		// Is a file
		return false, nil
	}
	return false, &errors.Error{Code: errors.INTERNAL, Message: "No path exists: " + path, Operation: op, Err: err}
}
