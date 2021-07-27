// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"net/http"
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

type AdminSocketData struct {
	Theme domain.ThemeConfig `json:"theme"`
}

// Admin
type Admin struct {
	ws *websocket.Conn
}

type adminhub struct {
	Broadcast chan AdminSocketData
	Receive   chan []byte
	Close     chan bool
}

var hub = adminhub{
	Broadcast: make(chan AdminSocketData),
	Receive:   nil,
	Close:     make(chan bool),
}

func Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "Admin.Handler"
		conn, err := adminUpgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error upgrading request to websocket", Operation: op, Err: err})
		}
		run(conn)
	}
}

func run(conn *websocket.Conn) {
	for {
		select {
		case as := <-hub.Broadcast:
			b, err := json.Marshal(as)
			if err != nil {
				logger.WithError(err).Error()
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				logger.WithError(err).Error()
				return
			}
		case <-hub.Close:
			conn.Close()
		}
	}
	conn.Close()
}

func BroadCast(a AdminSocketData) {
	hub.Broadcast <- a
}

func Close() {
	hub.Close <- true
}
