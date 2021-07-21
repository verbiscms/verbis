// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/services/fields/location"
	"github.com/gin-gonic/gin"
)

// Handler defines methods for fields to interact with the server.
type Handler interface {
	List(ctx *gin.Context)
}

// Fields defines the handler for all field routes.
type Fields struct {
	*deps.Deps
	finder location.Finder
}

// New
//
// Creates a new fields handler.
func New(d *deps.Deps) *Fields {
	return &Fields{
		Deps:   d,
		finder: &location.Location{},
	}
}
