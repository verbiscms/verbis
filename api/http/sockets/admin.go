// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gorilla/websocket"
	"github.com/radovskyb/watcher"
	"net/http"
	"os"
	"syscall"
	"time"
)

var (
	// upgrader
	adminUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// adminSocket
type adminSocket struct {
	ThemePath  string
	ConfigPath string
	Watcher    *watcher.Watcher
}

// Admin
//
//
func Admin(d *deps.Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "AdminSocket.Handler"

		ws, err := adminUpgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error upgrading request to websocket", Operation: op, Err: err})
		}

		a := &adminSocket{
			ThemePath:  d.ThemePath(),
			ConfigPath: d.ThemePath() + string(os.PathSeparator) + config.FileName,
			Watcher:    watcher.New(),
		}
		defer a.Watcher.Close()

		a.Init(ws)
	}
}

// Init
//
//
func (a *adminSocket) Init(ws *websocket.Conn) {
	const op = "AdminSocket.Init"

	a.Watcher.SetMaxEvents(1)

	go a.Writer(ws)

	// Watch this folder for changes.
	err := a.Watcher.Add(a.ConfigPath)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error adding configuration watcher", Operation: op, Err: err}).Error()
		ws.Close()
		return
	}

	// Start the watching process - it'll check for changes every 100ms.
	err = a.Watcher.Start(time.Millisecond * 100) //nolint
	defer a.Watcher.Close()
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error starting configuration watcher", Operation: op, Err: err}).Error()
		ws.Close()
		return
	}

	a.Reader(ws)
}

// Reader
//
//
func (a *adminSocket) Reader(conn *websocket.Conn) {
	for {
		// Read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			logger.WithError(err)
			return
		}

		fmt.Println(string(p))

		// Print out that message for clarity
		logger.Info(string(p))

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			fmt.Println(err)
			logger.Error(err)
			return
		}
	}
}

// Writer
//
//
func (a *adminSocket) Writer(ws *websocket.Conn) {
	const op = "AdminSocket.Writer"

	// go func() {
	for {
		select {
		case <-a.Watcher.Event:
			logger.Info("Updating theme configuration file, sending socket")

			// Marshal the configuration file.
			b, err := json.Marshal(config.Fetch(a.ThemePath))
			if err != nil {
				logger.WithError(&errors.Error{Code: op, Message: "Error marshalling theme configuration", Operation: op, Err: err}).Error()
				return
			}

			// Write the file back to the socket.
			err = ws.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				logger.WithError(&errors.Error{Code: op, Message: "Error sending socket message", Operation: op, Err: err}).Error()
				return
			}
		case err := <-a.Watcher.Error:
			if err != syscall.EPIPE {
				logger.WithError(&errors.Error{Code: op, Message: "Error watching theme configuration", Operation: op, Err: err}).Error()
			}
			return
		case <-a.Watcher.Closed:
			logger.Info("Closing watcher on theme configuration file")
			return
		}
	}
	// }()
}
