// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/radovskyb/watcher"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func reader(conn *websocket.Conn) {
	const op = "Sockets.Admin.reader"

	defer conn.Close()

	for {
		// Read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			logger.WithError(err)
			return
		}

		// Print out that message for clarity
		logger.Info(string(p))

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			logger.Error(err)
			return
		}
	}
}

func writer(ws *websocket.Conn) {
	const op = "OPCHANGE"

	w := watcher.New()
	w.SetMaxEvents(1)

	go func() {
		for {
			select {
			case _ = <-w.Event:
				logger.Info("Updating theme configuration file, sending message")

				b, err := json.Marshal(config.Fetch(paths.Theme()))
				if err != nil {
					logger.WithError(errors.Error{
						Code:      op,
						Message:   "Error marshalling ",
						Operation: "",
						Err:       nil,
					})
					logger.Error(err)
					return
				}

				err = ws.WriteMessage(websocket.TextMessage, b)
				if err != nil {
					logger.Error(err)
					return
				}

			case err := <-w.Error:
				logger.Error(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add(paths.Theme()); err != nil {
		logger.Error(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		logger.Error(err)
	}
}

func Admin(ctx *gin.Context) {
	const op = "OPCHANGE"

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error upgrading request to websocket", Operation: op, Err: err})
	}

	logger.Info("Admin client webhook connected")

	go writer(ws)
	reader(ws)
}
