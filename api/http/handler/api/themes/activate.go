// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// Activate
//
// Returns http.StatusOK if the theme was changed.
// Returns http.StatusInternalServerError if there was an error updating or formatting the post.
// Returns http.StatusBadRequest if the the validation failed, there was a conflict, or the post wasn't found.
func (t *Themes) Activate(ctx *gin.Context) {
	const op = "ThemeHandler.Update"

	theme := ctx.Param("name")

	config, err := t.Theme.Activate(theme)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully changed theme with the name: "+theme, config)
}
