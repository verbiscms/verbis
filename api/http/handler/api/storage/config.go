// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Config
//
// Returns http.StatusOK if the configuration was successfully retrieved.
// Returns http.StatusInternalServerError if there was an error obtaining the config.
func (s *Storage) Config(ctx *gin.Context) {
	const op = "StorageHandler.Config"

	info, err := s.Storage.Info()
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained configuration", info)
}
