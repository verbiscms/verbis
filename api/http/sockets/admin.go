// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/watchers"
	"github.com/gorilla/websocket"
	"github.com/radovskyb/watcher"
	"net/http"
	"syscall"
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

		conn, err := adminUpgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error upgrading request to websocket", Operation: op, Err: err})
		}

		go func() {
			for {
				writer(conn, d.Watcher, d.ThemePath())
			}
		}()
	}
}

// writer
//
//
func writer(conn *websocket.Conn, w *watchers.Batch, path string) {
	const op = "AdminSocket.Handler.Writer"

	select {
	case event := <-w.Event:
		if event.Name() != config.FileName && event.Op != watcher.Write {
			return
		}

		logger.Info("Updating theme configuration file, sending socket")
		cfg := config.Fetch(path)

		// Marshal the configuration file.
		b, err := json.Marshal(cfg)
		if err != nil {
			logger.WithError(&errors.Error{Code: op, Message: "Error marshalling theme configuration", Operation: op, Err: err}).Error()
			return
		}

		// Set the file back to the socket.
		err = conn.WriteMessage(websocket.TextMessage, b)
		if err != nil {
			logger.WithError(&errors.Error{Code: op, Message: "Error sending socket message", Operation: op, Err: err}).Error()
		}

	case err := <-w.Error:
		if err.Err != syscall.EPIPE {
			logger.WithError(&errors.Error{Code: op, Message: "Error watching theme configuration", Operation: op, Err: err}).Error()
		}
	}
}
