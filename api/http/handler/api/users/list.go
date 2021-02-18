// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// List
//
// Returns 200 if the users were obtained successfully.
// Returns 500 if there was an error getting the users.
// Returns 400 if there was conflict or the request was invalid.
func (u *Users) List(ctx *gin.Context) {
	const op = "UserHandler.List"

	params := params.ApiParams(ctx, api.DefaultParams).Get()

	users, total, err := u.Store.User.Get(params)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	pagination := http.NewPagination().Get(params, total)

	api.Respond(ctx, 200, "Successfully obtained users", users.HideCredentials(), pagination)
}