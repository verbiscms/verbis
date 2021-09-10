// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"github.com/verbiscms/verbis/api/services/storage"
	"net/http"
)

// Download
//
// Returns downloaded zip with "application/octet-stream" header
// if successful.
// Returns http.StatusBadRequest if the request was invalid or there was a conflict.
// Returns http.StatusInternalServerError if there was an error obtaining the files.
func (s *Storage) Download(ctx *gin.Context) {
	const op = "StorageHandler.Download"

	ctx.Header("Content-type", "application/octet-stream")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename='%s'", storage.DownloadFileName))
	ctx.Header("X-Filename", storage.DownloadFileName)

	err := s.Storage.Download(ctx.Writer)
	if err != nil && errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
}
