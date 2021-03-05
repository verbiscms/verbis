// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Find
//
// Returns http.StatusOK if the theme config was obtained.
// Returns http.StatusBadRequest if the name wasn't passed.
// Returns http.StatusInternalServerError if there as an error obtaining the config.
func (t *Themes) Find(ctx *gin.Context) {
	const op = "ThemeHandler.Find"

	theme, err := t.Theme.Find(ctx.Param("name"))
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained theme config", theme)
}
