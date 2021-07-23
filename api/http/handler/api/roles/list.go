// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// List
//
// Returns http.StatusOK if the roles were obtained successfully.
// Returns http.StatusInternalServerError if there was an error getting the roles.
func (u *Roles) List(ctx *gin.Context) {
	const op = "RoleHandler.List"

	roles, err := u.Store.Roles.List()
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained roles", roles)
}
