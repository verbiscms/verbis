// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gorilla/websocket"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAdminSocket(t *testing.T) {
	s := httptest.NewServer(http.Handler(Admin(&deps.Deps{})))
	defer s.Close()

	u := "ws" + strings.TrimPrefix(s.URL, "http")

	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	defer ws.Close()
	if err != nil {
		t.Error("error dialing websocket", err)
	}

	err = ws.WriteMessage(websocket.TextMessage, []byte("test"))
	if err != nil {
		t.Error(err)
	}

	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Hanging here
	fmt.Println(p)
}

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

			conn.WriteMessage(websocket.TextMessage, []byte("ffff"))

			time.Sleep(10 * time.Millisecond)

			fmt.Println(t.logger.String())

			within(t.T(), 1000*time.Millisecond, func() {
				tj, p, err := conn.ReadMessage()
				fmt.Println(tj)
				fmt.Println(err)
				fmt.Println(string(p))
			})

			//color.Red.Println(t.logger.String())

			// t.Equal(test.want, string(p))
		})
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case m := <-time.After(d):
		fmt.Println(m)
		t.Error("timed out")
	case <-done:
	}
}
