// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gin-gonic/gin"
)

// Handler defines methods for fields to interact with the server
type Handler interface {
	Clear(ctx *gin.Context)
}

// Cache defines the handler for Cache
type Cache struct {
	*deps.Deps
}
