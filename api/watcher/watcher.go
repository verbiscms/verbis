// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package watcher

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
	"strings"
)

type Watcher interface {
	Close()
	Watch()
}

type watch struct {
	watcher *fsnotify.Watcher
	dir     string
}

// New
//
// Creates a new fsnotify Watcher
func New(dir string) (Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &watch{
		watcher: watcher,
		dir:     dir,
	}, nil
}

// Close
//
//
func (w *watch) Close() {
	err := w.watcher.Close()
	if err != nil {
		logger.WithError(err).Error()
	}
}

// Watch
//
//
func (w *watch) Watch() {
	const op = "Watcher.Watch"

	logger.Debug("Starting watcher")

	// Starting at the theme path, walk each file/directory
	// searching for directories
	err := filepath.Walk(w.dir, w.watchDir)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error watching theme directory:" + paths.Theme(), Operation: op, Err: err}).Error()
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			// Watch for events
			case event := <-w.watcher.Events:

				fmt.Println("innn")
				fn := w.canExecute(event.Name)
				if fn != nil {
					fn(&event)
				}
			// Watch for errors
			case _ = <-w.watcher.Errors:
				logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error watching file", Operation: op, Err: err}).Error()
			}
		}
	}()

	<-done
}

// canExecute
//
//
func (w *watch) canExecute(path string) NotifyFunc {
	for i, v := range Registry {
		if path == v.path {
			return Registry[i].fn
		}
		if v.dir {
			if strings.Contains(path, v.path) {
				return Registry[i].fn
			}
		}
	}
	return nil
}

// watchDir
//
// watchDir gets run as a walk func, searching for directories
// to add watchers to/
func (w *watch) watchDir(path string, fi os.FileInfo, err error) error {
	const op = "Watcher.watchDir"

	// Since fsnotify can watch all the files in a directory,
	// watchers only need to be added to each nested
	// directory.
	if fi.Mode().IsDir() {
		err := w.watcher.Add(path)
		if err != nil {
			return &errors.Error{Code: errors.INVALID, Message: "Error walking path: " + path, Operation: op, Err: err}
		}
	}
	return nil
}
