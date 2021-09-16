// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// Disconnect
//
// Returns http.StatusOK if the provider changed successfully.
// Returns http.StatusBadRequest if the request was invalid or validation failed.
// Returns http.StatusInternalServerError if there was an error processing the change.
func (s *Storage) Disconnect(ctx *gin.Context) {
	const op = "StorageHandler.Disconnect"

	err := s.Storage.Disconnect()
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully updated storage options", nil)
}
