// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package watchers

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/radovskyb/watcher"
	"os"
	"syscall"
	"time"
)

const (
	PollingDuration = time.Millisecond * 100
)

type Config struct {
	path    string
	watcher *watcher.Watcher
	Event   chan *domain.ThemeConfig
	Error   chan *errors.Error
	Closed  chan struct{}
}

func NewConfig() *Config {
	return &Config{
		path:    "/Users/ainsley/Desktop/Reddico/apis/verbis/themes/Verbis",
		Event:   make(chan *domain.ThemeConfig),
		Error:   make(chan *errors.Error),
		Closed:  make(chan struct{}),
		watcher: watcher.New(),
	}
}

func (c *Config) Start() {
	const op = "Watcher.Config.Start"

	c.watcher.SetMaxEvents(1)

	go func() {
		for {
			select {
			case <-c.watcher.Event:
				cfg, err := config.Find(c.path)
				if err != nil {
					c.Error <- errors.ToError(err)
				} else {
					config.Set(*cfg)
					c.Event <- cfg
				}
			case err := <-c.watcher.Error:
				if err != syscall.EPIPE {
					c.Error <- &errors.Error{Code: op, Message: "Error watching theme configuration", Operation: op, Err: err}
				}

			case <-c.watcher.Closed:
				logger.Info("Closing watcher on theme configuration file")
				c.Closed <- struct{}{}
				return
			}
		}
	}()

	// Watch the config file for changes
	err := c.watcher.Add(c.path + string(os.PathSeparator) + config.FileName)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error adding configuration watcher", Operation: op, Err: err}).Error()
		return
	}

	// Start the config watcher
	err = c.watcher.Start(PollingDuration)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error starting configuration watcher", Operation: op, Err: err}).Error()
		return
	}
}

func (c *Config) Close() {
	c.watcher.Close()
}
