// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

// List
//
// Returns http.StatusOK if the users were obtained successfully.
// Returns http.StatusInternalServerError if there was an error getting the users.
// Returns http.StatusBadRequest if there was conflict or the request was invalid.
func (u *Users) List(ctx *gin.Context) {
	const op = "UserHandler.List"

	p := api.Params(ctx).Get()

	users, total, err := u.Store.User.Get(p)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained users", users.HideCredentials(), pagination.Get(p, total))
}
