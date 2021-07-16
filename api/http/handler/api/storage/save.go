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

type ()

// Save
//
// Returns http.StatusBadRequest if the request was invalid.
// Returns http.StatusOK if there are no buckets items or success.
// Returns http.StatusInternalServerError if there was an error getting the buckets.
func (s *Storage) Save(ctx *gin.Context) {
	const op = "StorageHandler.SetProvider"

	var info domain.StorageChange
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if !info.Provider.IsLocal() && info.Bucket == "" {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("bucket can't be empty"), Operation: op})
		return
	}

	buckets, err := s.Storage.ListBuckets(info.Provider)
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INTERNAL, Err: err, Operation: op})
	}

	if !info.Provider.IsLocal() && !buckets.IsValid(info.Bucket) {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("invalid bucket" + info.Bucket), Operation: op})
		return
	}

	cfg, err := s.Storage.Info()
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if !cfg.Providers[info.Provider].Connected {
		api.Respond(ctx, http.StatusBadRequest, "Storage provider not connected", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("storage provider dial error"), Operation: op})
		return
	}

	err = s.Deps.Store.Options.Update("storage_provider", info.Provider)
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if info.Provider.IsLocal() {
		info.Bucket = ""
	}

	err = s.Deps.Store.Options.Update("storage_bucket", info.Bucket)
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully set storage provider", nil)
}
