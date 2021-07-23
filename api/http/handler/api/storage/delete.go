// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteBucket
//
// Returns http.StatusOK if the bucket was deleted successfully.
// Returns http.StatusBadRequest if the request was invalid or validation failed.
// Returns http.StatusInternalServerError if there was an error deleting the bucket.
func (s *Storage) DeleteBucket(ctx *gin.Context) {
	const op = "StorageHandler.DeleteBucket"

	var info domain.StorageChange
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if info.Provider.IsLocal() {
		api.Respond(ctx, http.StatusBadRequest, "Local bucket cannot be deleted", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("error deleting bucket"), Operation: op})
		return
	}

	err = s.Deps.Storage.DeleteBucket(info.Provider, info.Bucket)
	if err != nil && errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully deleted bucket: "+info.Bucket, nil)
}
