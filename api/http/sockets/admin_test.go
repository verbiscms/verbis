// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func (t *SocketsTestSuite) Test_AdminSocket() {
	tt := map[string]struct {
		message string
		want    interface{}
	}{
		"Success": {
			"test",
			"test",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			conn, _ := t.Setup()
			//defer teardown()

			err := conn.WriteMessage(websocket.TextMessage, []byte(test.message))
			if err != nil {
				t.Error(err)
			}

			err = conn.WriteMessage(websocket.TextMessage, []byte("ffff"))
			if err != nil {
				t.Error(err)
			}

			time.Sleep(10 * time.Millisecond)

			fmt.Println(t.logger.String())

			//color.Red.Println(t.logger.String())

			// t.Equal(test.want, string(p))
		})
	}
}
