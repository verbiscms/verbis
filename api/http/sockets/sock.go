// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"
)

// S defines
var S = melody.New()

// Handler
//
//
func Handler(d *deps.Deps) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("verbis-session")
		if err != nil || cookie == nil {
			api.Respond(ctx, 401, "Unauthorised used of websocket, no session found", nil)
			return
		}

		err = d.Store.User.CheckSession(cookie.Value)
		if err != nil {
			api.Respond(ctx, 401, "Unauthorised used of websocket, session expired", nil)
			return
		}

		err = S.HandleRequest(ctx.Writer, ctx.Request)
		if err != nil {
			log.Error(err)
			return
		}
	}
}

// Broadcast
//
//
func Broadcast(i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	err = S.Broadcast(b)
	if err != nil {
		return err
	}
	return nil
}
