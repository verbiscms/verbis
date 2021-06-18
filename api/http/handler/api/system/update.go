// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Update
//
// TODO, update comments
// Returns http.StatusOK if the user was created.
// Returns http.StatusInternalServerError if there was an error creating the user.
// Returns http.StatusBadRequest if the the validation failed or a user already exists.
func (s *System) Update(ctx *gin.Context) {
	const op = "SystemHandler.Update"

	ver, err := s.System.Update()
	if err != nil && errors.Code(err) == errors.INVALID {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	api.Respond(ctx, http.StatusOK, fmt.Sprintf("Verbis updated successfully to version %s, restarting system....", ver), nil)
}
