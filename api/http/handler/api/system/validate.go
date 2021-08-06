// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"strconv"
)

// ValidateInstall
//
// Returns http.StatusOK if the preflight check was successful.
// Returns http.StatusBadRequest if the preflight failed (database).
func (s *System) ValidateInstall(ctx *gin.Context) {
	const op = "SystemHandler.Validate"

	if s.Installed {
		api.Respond(ctx, http.StatusBadRequest, "Verbis is already installed", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("already installed"), Operation: op})
		return
	}

	step, err := strconv.Atoi(ctx.Param("step"))
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Pass a valid number to validate a step", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	var install domain.InstallVerbis
	err = json.NewDecoder(ctx.Request.Body).Decode(&install) // Skip validation for step
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Error unmarshalling install", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = s.System.ValidateInstall(step, install)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully validated install data", nil)
}
