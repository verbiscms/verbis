// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// CreateBucket
//
// Returns http.StatusOK if the bucket was created successfully.
// Returns http.StatusBadRequest if the request was invalid or validation failed.
// Returns http.StatusInternalServerError if there was an error creating the bucket.
func (s *Storage) CreateBucket(ctx *gin.Context) {
	const op = "StorageHandler.CreateBucket"

	var info domain.StorageChange
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	bucket, err := s.Storage.CreateBucket(info.Provider, info.Bucket)
	if err != nil && errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully created bucket: "+bucket.Name, bucket)
}
