// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

// List
//
// Returns http.StatusOK if there are no media items or success.
// Returns http.StatusInternalServerError if there was an error getting the media items.
// Returns http.StatusBadRequest if there was conflict or the request was invalid.
func (m *Media) List(ctx *gin.Context) {
	const op = "MediaHandler.List"

	p := api.Params(ctx).Get()

	media, total, err := m.service.List(p)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained media", media.Public(), pagination.Get(p, total))
}
