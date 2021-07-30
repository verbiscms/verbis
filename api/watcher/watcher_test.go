// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package watcher

import (
	"fmt"
	"github.com/radovskyb/watcher"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/errors"
	"sync"
	"testing"
	"time"
)

func TestEvents_IsPath(t *testing.T) {
	tt := map[string]struct {
		input string
		path  string
		want  bool
	}{
		"True": {
			"/users/verbis/cms/theme",
			"/users/verbis/cms/theme/template.cms",
			true,
		},
		"False": {
			"wrong",
			"/users/verbis/cms/theme/template.cms",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			e := Event{
				Event: watcher.Event{Path: test.path},
			}
			got := e.IsPath(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestBatcher_Close(t *testing.T) {
	b := New()
	go func() {
		time.Sleep(200 * time.Millisecond)
		b.Close()
	}()
	err := b.Watch(t.TempDir(), PollingDuration)
	assert.NoError(t, err)
}

func TestBatcher_Watch(t *testing.T) {
	tt := map[string]struct {
		process  func(b *Batcher, wg *sync.WaitGroup)
		callback func(b *Batcher)
	}{
		"Event": {
			func(b *Batcher, wg *sync.WaitGroup) {
				defer wg.Done()
				event := <-b.Events()
				assert.Equal(t, watcher.Create, event.Op)
			},
			func(b *Batcher) {
				b.watcher.TriggerEvent(watcher.Chmod, nil)
				b.watcher.TriggerEvent(watcher.Create, nil)
			},
		},
		"Error": {
			func(b *Batcher, wg *sync.WaitGroup) {
				defer wg.Done()
				err := <-b.Errors()
				assert.Equal(t, "Error watching file system", err.Message)
			},
			func(b *Batcher) {
				b.watcher.Error <- &errors.Error{Message: "message"}
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			b := New()

			wg := sync.WaitGroup{}
			wg.Add(1)

			go test.process(b, &wg)

			go func() {
				err := b.Watch(t.TempDir(), PollingDuration)
				if err != nil {
					fmt.Println(err)
				}
			}()

			test.callback(b)
			wg.Wait()
		})
	}
}

func TestBatcher_Watch_Error(t *testing.T) {
	tt := map[string]struct {
		path string
		poll time.Duration
		want interface{}
	}{
		"Wrong Path": {
			"wrong",
			PollingDuration,
			"Error adding path",
		},
		"Bad Poll": {
			t.TempDir(),
			0,
			"Error starting watcher",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			b := New()
			err := b.Watch(test.path, test.poll)
			if err == nil {
				t.Fatalf("eror")
				return
			}
			assert.Contains(t, errors.Message(err), test.want)
		})
	}
}
