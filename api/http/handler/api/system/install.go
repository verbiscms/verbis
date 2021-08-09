// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"github.com/verbiscms/verbis/api/logger"
	"net/http"
	"time"
)

// Install
//
// Returns http.StatusOK if the roles were obtained successfully.
// Returns http.StatusBadRequest if validation failed or the version is installed.
func (s *System) Install(ctx *gin.Context) {
	const op = "SystemHandler.Install"

	if s.Installed {
		api.Respond(ctx, http.StatusBadRequest, "Verbis is already installed", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("already installed"), Operation: op})
		return
	}

	var install domain.InstallVerbis
	err := ctx.ShouldBindJSON(&install)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = s.System.Install(install)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully installed Verbis", nil)

	go func() {
		time.Sleep(time.Second * 1)
		err = s.System.Restart()
		if err != nil {
			logger.WithError(err).Panic()
		}
	}()
}
