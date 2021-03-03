// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Find
//
// Returns http.StatusOK if there are no options or success.
// Returns http.StatusBadRequest if there was name param was missing.
// Returns http.StatusInternalServerError if there was an error getting the options.
func (o *Options) Find(ctx *gin.Context) {
	const op = "OptionsHandler.Find"

	name := ctx.Param("name")
	option, err := o.Store.Options.GetByName(name)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained option with name: "+name, option)
}
