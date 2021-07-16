// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/gin-gonic/gin"
)

// CreateBucket
//
// Returns http.StatusBadRequest if the request was invalid.
// Returns http.StatusOK if there are no buckets items or success.
// Returns http.StatusInternalServerError if there was an error getting the buckets.
func (s *Storage) CreateBucket(ctx *gin.Context) {
	const op = "StorageHandler.CreateBucket"

	//var info domain.StorageInfo
	//err := ctx.ShouldBindJSON(&info)
	//if err != nil {
	//	api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	//	return
	//}
	//
	//err = s.Deps.Storage.CreateBucket(info.Provider, info.Bucket)
	//if err != nil && errors.Code(err) == errors.INVALID {
	//	api.Respond(ctx, http.StatusBadRequest, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	//	return
	//} else if err != nil {
	//	api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	//	return
	//}
	//
	//api.Respond(ctx, http.StatusOK, "Successfully created bucket: ", info.Bucket)
}
