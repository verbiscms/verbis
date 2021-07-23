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

// ListBuckets
//
// Returns http.StatusForbidden if the provider is local.
// Returns http.StatusBadRequest if the request was invalid.
// Returns http.StatusOK if there are no buckets items or success.
// Returns http.StatusInternalServerError if there was an error getting the buckets.
func (s *Storage) ListBuckets(ctx *gin.Context) {
	const op = "StorageHandler.ListBuckets"

	provider := domain.StorageProvider(ctx.Param("name"))
	if provider.IsLocal() {
		api.Respond(ctx, http.StatusForbidden, "Obtaining local buckets are forbidden", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("error bad provider"), Operation: op})
		return
	}

	buckets, err := s.Deps.Storage.ListBuckets(provider)
	if err != nil && errors.Code(err) == errors.INVALID {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained buckets", buckets)
}
