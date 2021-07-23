// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// Config
//
// Returns http.StatusOK if theme config was obtained successfully.
func (t *Themes) Config(ctx *gin.Context) {
	api.Respond(ctx, http.StatusOK, "Successfully obtained theme config", t.Deps.Config)
}
