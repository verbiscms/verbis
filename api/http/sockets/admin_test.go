// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/verbiscms/verbis/api/domain"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (t *SocketsTestSuite) TestHandler_Error() {
	t.T().Skip() // Data race
	orig := adminUpgrade
	defer func() { adminUpgrade = orig }()
	adminUpgrade.CheckOrigin = func(r *http.Request) bool {
		return false
	}
	s := httptest.NewServer(Handler())
	_, _, err := websocket.DefaultDialer.Dial("ws"+strings.TrimPrefix(s.URL, "http"), nil)
	t.Error(err)
}

func (t *SocketsTestSuite) TestAdminSocket() {
	conn, teardown := t.Setup(Handler())
	defer teardown()

	cfg := SendData{Theme: domain.ThemeConfig{
		Theme: domain.Theme{
			Title: "verbis",
		},
	}}

	AdminHub.Broadcast <- cfg
	_, got, err := conn.ReadMessage()
	t.NoError(err)

	want, err := json.Marshal(cfg)
	if err != nil {
		t.Fail("error marshalling config", err)
	}

	t.Equal(string(want), string(got))
}
