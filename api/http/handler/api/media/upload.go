// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// Upload
//
// If there were no files attached to the body, more than
// 1 attached to the body or the validation failed.
//
// Returns 401 if the user wasn't authenticated.
// Returns 415 if the media item failed to validate.
// Returns 200 if the media item was successfully uploaded.
// Returns 500 if there as an error uploading the media item.
// Returns 400 if the file length was incorrect or there were no files.
func (m *Media) Upload(ctx *gin.Context) {
	const op = "MediaHandler.Upload"

	form, err := ctx.MultipartForm()
	if err != nil {
		api.Respond(ctx, 400, "No files attached to the upload", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	files := form.File["file"]

	if len(files) > 1 {
		api.Respond(ctx, 400, "Files are only permitted to be uploaded one at a time", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("too many files uploaded at once"), Operation: op})
		return
	}

	if len(files) == 0 {
		api.Respond(ctx, 400, "Attach a file to the request to be uploaded", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("no files attached to upload"), Operation: op})
		return
	}

	err = m.Store.Media.Validate(files[0])
	if err != nil {
		api.Respond(ctx, 415, errors.Message(err), err)
		return
	}

	media, err := m.Store.Media.Upload(files[0], ctx.Request.Header.Get("token"))
	if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully uploaded media item", media)
}
