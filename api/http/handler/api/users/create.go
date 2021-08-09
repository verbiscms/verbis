// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"strconv"
)

// Create
//
// Returns http.StatusOK if the user was created.
// Returns http.StatusInternalServerError if there was an error creating the user.
// Returns http.StatusBadRequest if the the validation failed or a user already exists.
func (u *Users) Create(ctx *gin.Context) {
	const op = "UserHandler.Create"

	var userCreate domain.UserCreate
	err := ctx.ShouldBindJSON(&userCreate)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := u.Store.User.Create(userCreate)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully created user with ID: "+strconv.Itoa(user.ID), user)
}
