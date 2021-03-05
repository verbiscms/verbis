// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package watchers

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/radovskyb/watcher"
	"path/filepath"
)

type Batcher interface {
	Start()
	Close()
}

// Batch
type Batch struct {
	watcher *watcher.Watcher

	// Events for the file system are returned on this
	// channel.
	Event chan Event

	// Errors that occurred reading a file are returned
	// on this channel.
	Error chan *errors.Error

	// The path watch watch recursively.
	path string

	// Determines if the watcher is running.
	Running bool

	// Done chan
	done chan struct{}
}

// Event
type Event struct {
	watcher.Event

	Extension string
	Mime      string
}

var (
	excludedDir = []string{
		"node_modules",
		"bower_components",
		".git",
	}
	excludedFiles = []string{
		".DS_STORE",
		"robots.txt",
	}
)

func New(themePath string) *Batch {
	return &Batch{
		watcher: watcher.New(),
		Event:   make(chan Event),
		Error:   make(chan *errors.Error),
		path:    themePath,
		done:    make(chan struct{}),
	}
}

// UpdateTheme
//
//
func (b *Batch) UpdateTheme(themePath string) {
	b.Close()
	b.path = themePath
	b.Start()
}

// Start
//
//
func (b *Batch) Start() {
	b.Running = true
	go b.run()
}

// Close
//
//
func (b *Batch) Close() {
	b.Running = false
	b.done <- struct{}{}
	b.Close()
}

// run
//
//
func (b *Batch) run() {
	const op = "Watcher.Start"

	b.watcher.SetMaxEvents(1)

	go func() {
		//OuterLoop:
		for {
			select {
			case event := <-b.watcher.Event:
				//name := event.Name()
				//if event.IsDir() {
				//	continue OuterLoop
				//}
				//if isExcludedDir(name) || isExcludedFile(name) || name == b.path {
				//	continue OuterLoop
				//}
				ext := filepath.Ext(event.Path)
				b.Event <- Event{
					Event:     event,
					Extension: ext,
					Mime:      mime.TypeByExtension(ext),
				}
			case err := <-b.watcher.Error:
				b.Error <- &errors.Error{Code: op, Message: "Error watching theme", Operation: op, Err: err}
			case <-b.watcher.Closed:
				b.Close()
				return
			}
		}
	}()

	err := b.watcher.AddRecursive(b.path)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error adding watcher", Operation: op, Err: err}).Error()
		return
	}

	err = b.watcher.Start(PollingDuration)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error starting configuration watcher", Operation: op, Err: err}).Error()
		return
	}
}

// isExcludedDir
//
//
func isExcludedDir(name string) bool {
	for _, v := range excludedDir {
		if v == name {
			return true
		}
	}
	return false
}

// isExcludedFile
//
//
func isExcludedFile(name string) bool {
	for _, v := range excludedFiles {
		if v == name {
			return true
		}
	}
	return false
}
