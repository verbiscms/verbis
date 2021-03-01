// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func (t *SocketsTestSuite) TestCategories_Create() {

	ws, teardown := t.Setup()
	defer teardown()

	//// Send message to server, read response and check to see if it's what we expect.
	//for i := 0; i < 10; i++ {
	//

	//}

	if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		t.Error(err)
	}

	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Error(err)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("hh")
	if string(p) != "hello" {
		t.Error(err)
	}

	fmt.Println("hh")
	ws.Close()
}
