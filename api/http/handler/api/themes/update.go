// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/services/theme"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// UpdateTheme defines the data to be validated when a
// theme is switched.
type UpdateTheme struct {
	Theme string `json:"theme" binding:"required"`
}

// Update
//
// Returns http.StatusOK if the theme was changed.
// Returns http.StatusInternalServerError if there was an error updating or formatting the post.
// Returns http.StatusBadRequest if the the validation failed, there was a conflict, or the post wasn't found.
func (t *Themes) Update(ctx *gin.Context) {
	const op = "PostHandler.Update"

	var u UpdateTheme
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	ok := t.Theme.Exists(u.Theme)
	if !ok {
		api.Respond(ctx, http.StatusBadRequest, "No theme exists with the name: "+u.Theme, &errors.Error{Code: errors.INVALID, Err: theme.ErrNoTheme, Operation: op})
		return
	}

	err = t.SetTheme(u.Theme)
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, "Error setting theme", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	config.Fetch(t.Paths.Themes + string(os.PathSeparator) + u.Theme)

	api.Respond(ctx, http.StatusOK, "Successfully changed theme with the name: "+u.Theme, config.Get())
}
