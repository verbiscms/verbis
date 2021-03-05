// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Find
//
// Returns http.StatusOK if theme config was obtained successfully.
func (t *Themes) Find(ctx *gin.Context) {
	const op = "ThemeHandler.Find"

	name := ctx.Param("name")
	if name == "" {
		api.Respond(ctx, http.StatusBadRequest, "Pass a valid string to obtain the redirect by name", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("no theme passed"), Operation: op})
	}

	theme, err := t.Theme.Find(name)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained theme config", theme)
}
