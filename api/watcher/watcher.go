// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package watcher

import (
	"github.com/radovskyb/watcher"
	"github.com/verbiscms/verbis/api/common/mime"
	"github.com/verbiscms/verbis/api/errors"
	"path/filepath"
	"strings"
	"time"
)

// FileWatcher is an interface for implementing file notification
// watchers.
type FileWatcher interface {
	Watch(path string, poll time.Duration) error
	Events() <-chan Event
	Errors() <-chan errors.Error
	Close()
}

// Event is what is sent back through the Event channel on the
// FileWatcher. It includes the original event, mime type
// and extension.
type Event struct {
	watcher.Event
	Mime      string
	Extension string
}

const Watch = watcher.Chmod

// Batcher describes the implementation of a FileWatcher.
type Batcher struct {
	watcher *watcher.Watcher
	events  chan Event
	errors  chan errors.Error
}

// IsPath determines if the event contains a given path.
func (e *Event) IsPath(path string) bool {
	return strings.Contains(e.Path, path)
}

// New creates a new watcher and setups up relevant
// channels. This function will not call start.
func New() *Batcher {
	return &Batcher{
		watcher: watcher.New(),
		events:  make(chan Event),
		errors:  make(chan errors.Error),
	}
}

// Events returns a file watcher event
func (b *Batcher) Events() <-chan Event {
	return b.events
}

// Errors returns the error channel for file watching.
func (b *Batcher) Errors() <-chan errors.Error {
	return b.errors
}

// Close closes the file watcher.
func (b *Batcher) Close() {
	b.watcher.Close()
}

// PollingDuration is the default duration that file
// watchers should poll for.
const PollingDuration = time.Millisecond * 1

// Watch is a blocking operation. It should be called with
// go func(). It adds the path to the watcher and polls
// for file changes with the given time.Duration.
// Events or errors.Error's are returned on
// the channels.
func (b *Batcher) Watch(path string, poll time.Duration) error {
	const op = "Watcher.Watch"

	// Iterate over the event, error and closed channels.
	go func() {
		for {
			select {
			case event := <-b.watcher.Event:
				if event.Op == watcher.Chmod {
					continue
				}
				b.events <- Event{
					Event:     event,
					Mime:      mime.TypeByExtension(filepath.Ext(event.Path)),
					Extension: filepath.Ext(event.Path),
				}
			case err := <-b.watcher.Error:
				b.errors <- errors.Error{
					Code:      errors.INTERNAL,
					Message:   "Error watching file system",
					Operation: op,
					Err:       err,
				}
			case <-b.watcher.Closed:
				return
			}
		}
	}()

	// Add the recursive path to watch for.
	err := b.watcher.AddRecursive(path)
	if err != nil {
		return err
	}

	// Start the watching process - it'll check for changes
	// with the given polling duration.
	err = b.watcher.Start(poll)
	if err != nil {
		return err
	}

	return nil
}
