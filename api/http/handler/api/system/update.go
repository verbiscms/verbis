// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// Update
//
// Returns http.StatusOK if the system updated successfully.
// Returns http.StatusBadRequest if the system is already updated,
// Returns http.StatusInternalServerError if the system could not be updated.
func (s *System) Update(ctx *gin.Context) {
	const op = "SystemHandler.Update"

	ver, err := s.System.Update(true)
	if err != nil && errors.Code(err) == errors.INVALID {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	api.Respond(ctx, http.StatusOK, fmt.Sprintf("Verbis updated successfully to version %s, restarting system....", ver), nil)
}
