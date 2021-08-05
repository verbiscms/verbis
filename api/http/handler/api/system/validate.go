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
	"strconv"
)

// Preflight
//
// Returns http.StatusOK if the preflight check was successful.
// Returns http.StatusBadRequest if the preflight failed (database).
func (s *System) Preflight(ctx *gin.Context) {
	const op = "SystemHandler.Preflight"

	if s.Installed {
		api.Respond(ctx, http.StatusBadRequest, "Verbis is already installed", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("already installed"), Operation: op})
		return
	}

	step, err := strconv.Atoi(ctx.Param("step"))
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Pass a valid number to validate a step", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	var data interface{}
	err = ctx.Bind(&data)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	fmt.Println(data, step)


	err = s.System.Validate(step, data)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}



	//switch step {
	//case sys.InstallDatabaseStep:
	//	var install domain.InstallDatabase
	//	err := ctx.Bind(&install)
	//	if err != nil {
	//		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	//		return
	//	}
	//	s.System.ValidateDatabaseStep()
	//}


	//var install domain.InstallDatabase
	//err := ctx.ShouldBindJSON(&install)
	//if err != nil {
	//	api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	//	return
	//}
	//
	//err = s.System.Preflight(install)
	//if err != nil {
	//	api.Respond(ctx, http.StatusBadRequest, err.Error(), &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("error connecting to database"), Operation: op})
	//	return
	//}

	api.Respond(ctx, http.StatusOK, "Successfully connected to database", nil)
}
