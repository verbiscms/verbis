// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// Find
//
// Returns 200 if there are no options or success.
// Returns 400 if there was name param was missing.
// Returns 500 if there was an error getting the options.
func (o *Options) Find(ctx *gin.Context) {
	const op = "OptionsHandler.Find"

	name := ctx.Param("name")
	option, err := o.Store.Options.GetByName(name)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 200, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully obtained option with name: "+name, option)
}