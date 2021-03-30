// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Upload
//
// If there were no files attached to the body, more than
// 1 attached to the body or the validation failed.
//
// Returns http.StatusUnauthorized if the user wasn't authenticated.
// Returns http.StatusOK if the media item was successfully uploaded.
// Returns http.StatusUnsupportedMediaType if the media item failed to validate.
// Returns http.StatusInternalServerError if there as an error uploading the media item.
// Returns http.StatusBadRequest if the file length was incorrect or there were no files.
func (m *Media) Upload(ctx *gin.Context) {
	const op = "MediaHandler.Upload"

	form, err := ctx.MultipartForm()
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "No files attached to the upload", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	files := form.File["file"]

	if len(files) > 1 {
		api.Respond(ctx, http.StatusBadRequest, "Files are only permitted to be uploaded one at a time", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("too many files uploaded at once"), Operation: op})
		return
	}

	if len(files) == 0 {
		api.Respond(ctx, http.StatusBadRequest, "Attach a file to the request to be uploaded", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("no files attached to upload"), Operation: op})
		return
	}

	err = m.Service.Validate(files[0])
	if err != nil {
		api.Respond(ctx, http.StatusUnsupportedMediaType, errors.Message(err), err)
		return
	}

	item, err := m.Service.Upload(files[0])
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	media, err := m.Store.Media.Create(item)
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully uploaded media item", media)
}
