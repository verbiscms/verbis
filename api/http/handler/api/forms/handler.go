// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/events"
)

// Handler defines methods for forms to interact with the server.
type Handler interface {
	List(ctx *gin.Context)
	Find(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Send(ctx *gin.Context)
}

// Forms defines the handler for all form routes.
type Forms struct {
	*deps.Deps
	// Form send email event.
	formSend events.Dispatcher
}

// New
//
// Creates a new forms handler.
func New(d *deps.Deps) *Forms {
	return &Forms{
		Deps:     d,
		formSend: events.NewFormSend(d),
	}
}
