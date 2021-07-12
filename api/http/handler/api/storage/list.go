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

// ListBuckets
//
// Returns http.StatusBadRequest if the request was invalid.
// Returns http.StatusOK if there are no buckets items or success.
// Returns http.StatusInternalServerError if there was an error getting the buckets.
func (s *Storage) ListBuckets(ctx *gin.Context) {
	const op = "StorageHandler.ListBuckets"

	buckets, err := s.Deps.Storage.ListBuckets()
	if err != nil && errors.Code(err) == errors.INVALID {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained buckets", buckets)
}
