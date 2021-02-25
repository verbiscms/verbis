// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	"github.com/gin-gonic/gin"
)

// List
//
// Returns 200 if there are no media items or success.
// Returns 500 if there was an error getting the media items.
// Returns 400 if there was conflict or the request was invalid.
func (m *Media) List(ctx *gin.Context) {
	const op = "MediaHandler.List"

	p := api.Params(ctx).Get()

	media, total, err := m.Store.Media.Get(p)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully obtained media", media, pagination.Get(p, total))
}
