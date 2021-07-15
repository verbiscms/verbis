// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteBucket
//
// Returns http.StatusBadRequest if the request was invalid.
// Returns http.StatusOK if there are no buckets items or success.
// Returns http.StatusInternalServerError if there was an error getting the buckets.
func (s *Storage) DeleteBucket(ctx *gin.Context) {
	const op = "StorageHandler.DeleteBucket"

	var info domain.StorageInfo
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = s.Deps.Storage.DeleteBucket(info.Provider, info.Bucket)
	if err != nil && errors.Code(err) == errors.INVALID {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully deleted bucket: ", info.Bucket)
}
