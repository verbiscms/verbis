// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// migration represents the data send from the frontend
// to start a migration.
type migration struct {
	From   domain.StorageChange `json:"from"`
	To     domain.StorageChange `json:"to"`
	Delete bool                 `json:"delete"`
}

// Migrate
//
// Returns http.StatusOK if the migration started successfully.
// Returns http.StatusBadRequest if the request was invalid or validation failed.
// Returns http.StatusInternalServerError if there was an error obtaining files to migrate.
func (s *Storage) Migrate(ctx *gin.Context) {
	const op = "StorageHandler.Migrate"

	var migrate migration
	err := ctx.ShouldBindJSON(&migrate)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	total, err := s.Storage.Migrate(ctx, migrate.From, migrate.To, migrate.Delete)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, fmt.Sprintf("Successfully started migration, processing %d files", total), nil)
}
