// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package watchers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/fsnotify/fsnotify"
	"github.com/radovskyb/watcher"
	"time"
)

const (
	PollingDuration = time.Millisecond * 100
)

// Batch
type Batch struct {
	watcher *watcher.Watcher

	fsWatcher *fsnotify.Watcher
	Events    chan []fsnotify.Event // Events are returned on this channel
	done      chan struct{}

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
	stop chan struct{}
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

	kk, _ := fsnotify.NewWatcher()

	return &Batch{
		watcher:   watcher.New(),
		fsWatcher: kk,
		Event:     make(chan Event),
		Error:     make(chan *errors.Error),
		path:      themePath,
		stop:      make(chan struct{}, 1),
	}
}

// UpdateTheme
//
//
func (b *Batch) SetTheme(themePath string) {
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
	//close(b.stop)
	//fmt.Println("got here")
	//b.stop = make(chan struct{}, 1)
	b.done <- struct{}{}
	b.fsWatcher.Close()
	//b.watcher.Close()
}

// run
//
//

func (b *Batch) run() {
	interval := time.Millisecond * 100
	tick := time.Tick(interval)
	evs := make([]fsnotify.Event, 0)
OuterLoop:
	for {
		select {
		case ev := <-b.fsWatcher.Events:
			evs = append(evs, ev)
			fmt.Println(ev)
		case <-tick:
			if len(evs) == 0 {
				continue
			}
			b.Events <- evs
			evs = make([]fsnotify.Event, 0)
		case <-b.done:
			break OuterLoop
		}
	}
	close(b.done)
}

//func (b *Batch) run() {
//	const op = "Watcher.Start"
//
//	b.watcher.SetMaxEvents(1)
//
//	//select {
//	//case b.done <- struct{}{}:
//	//	// thing will close eventually
//	//default:
//	//	// thing already has a close; Close probably called twice
//	//}
//	fmt.Println("in run")
//
//	go func() {
//	OuterLoop:
//		for {
//			select {
//			case event := <-b.watcher.Event:
//				name := event.Name()
//				if isExcludedDir(name) || isExcludedFile(name) || name == b.path {
//					continue OuterLoop
//				}
//				ext := filepath.Ext(event.Path)
//				b.Event <- Event{
//					Event:     event,
//					Extension: ext,
//					Mime:      mime.TypeByExtension(ext),
//				}
//			case err := <-b.watcher.Error:
//				b.Error <- &errors.Error{Code: op, Message: "Error watching theme", Operation: op, Err: err}
//			case b.stop <- struct{}{}:
//				break OuterLoop
//			case <-b.watcher.Closed:
//				break OuterLoop
//			default:
//				return
//			}
//		}
//	}()
//
//	err := b.watcher.AddRecursive(b.path)
//	if err != nil {
//		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error adding watcher", Operation: op, Err: err}).Error()
//		return
//	}
//
//	err = b.watcher.Start(PollingDuration)
//	if err != nil {
//		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error starting configuration watcher", Operation: op, Err: err}).Error()
//		return
//	}
//}

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
