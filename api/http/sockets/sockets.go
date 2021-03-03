// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"github.com/gorilla/websocket"
)

type Connector interface {
	Reader(ws *websocket.Conn)
	Writer(ws *websocket.Conn)
}

const (
// Time allowed to write the file to the client.
//writeWait = 10 * time.Second

// Time allowed to read the next pong message from the client.
//pongWait = 5 * time.Second
)
